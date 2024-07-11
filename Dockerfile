# Build the manager binary
FROM registry.access.redhat.com/ubi8/go-toolset:1.21 as builder
ARG TARGETOS
ARG TARGETARCH
ARG SEALIGHTS_TOKEN='eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL1BST0QtUE9DLmF1dGguc2VhbGlnaHRzLmlvLyIsImp3dGlkIjoiUFJPRC1QT0MsbmVlZFRvUmVtb3ZlLEFQSUdXLWVmZmQwY2NiLTlmZjctNGI4ZS1iZWY3LTZiNWM2YmNhNDIwYywxNzE1ODU0NzA5OTc4Iiwic3ViamVjdCI6InJlZGhhdC1wb2NAYWdlbnQiLCJhdWRpZW5jZSI6WyJhZ2VudHMiXSwieC1zbC1yb2xlIjoiYWdlbnQiLCJ4LXNsLXNlcnZlciI6Imh0dHBzOi8vcmVkaGF0LXBvYy5zZWFsaWdodHMuY28vYXBpIiwic2xfaW1wZXJfc3ViamVjdCI6IiIsImlhdCI6MTcxNTg1NDcwOX0.bWXLTcrXtCT1LR02OlTElMomwd3f1uOG5OUoPDPENW8reZ3T6CCAGxXP_z-zDVIN-_3P3HPDlmh8w1hru4yPWvIY4kiDqzh1vWHsB1abTlLhzBFhOcIWrX2ynPGzq65cBHwPmcNSNH65ps68O75Y0FcFPx271ZqTTxNTfLP2xlTeveNi5Ma8eWWe1BR_nooaKx3RuuIV5SYe_ZJKvELaJ4nIzWXeuL9kxMjAJJThXC7d69-3UeUDGwdt5abuNcdaoNBgAshL9MZQSsn22oI477g9wCNM3UkaUUVaFlwFML5JOjkheM2t6Xn4gGY7i1u2dB0iUBKhJARQoMFcRKHeXAb4lGF0D15fPdsA19bMa7dUtGtvXTC23LVSsiFzWIbpXbx-6AWS6flPAthlD0S-UPS7oY-gmxVQqYMj3GtgBBmoz2oRAO6Shbfl8KcHJxENxLAwxMkT4hRiWX2MJoKmDrXB6e4oxasW20gmrS5CRwbZw5gv5vrtsI883yRFcN8byQXED1J_LAH2pVvyY2v9RXJyAoba6VT72KPz4rBqtB7tEoXVct5_6m_OeYlp1O6wpsoIOBP7fB_gFZgglingpXOVGEt7otIysaFSV7gjqyO90c4KYTd92-tjFhoNNBgTl-jOXrrG-mOG5X9nt_YkzEiZ0iFBNyEOCRmMiibUZPQ'
ARG BRANCH=sealights_poc
ARG BUILD_NAME=240610_1723
ARG SERVICE_NAME=dspo
ENV AGENT_URL='https://agents.sealights.co/slgoagent/latest/slgoagent-linux-amd64.tar.gz'
ENV AGENT_URL_SLCI='https://agents.sealights.co/slcli/latest/slcli-linux-amd64.tar.gz'

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

COPY . .
#Downloading the agents 
USER root
RUN wget -U "slgoagent" -q -O slgoagent.tar.gz \
   ${AGENT_URL} \
    && tar -xvf slgoagent.tar.gz \
    && rm slgoagent.tar.gz
RUN wget -U "slcli" -q -O slcli.tar.gz \
    ${AGENT_URL_SLCI} \
    && tar -xvf slcli.tar.gz \
    && rm slcli.tar.gz

#Adding sealights to the code
RUN ./slcli config init --lang go --token ${SEALIGHTS_TOKEN}
#Making the build session id 
RUN ./slcli config create-bsid --app ${SERVICE_NAME} --branch ${BRANCH} --build ${BUILD_NAME}
#Scanning the code
RUN ./slcli scan --bsid buildSessionId.txt --path-to-scanner ./slgoagent --workspacepath ./ --scm git 


# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
USER root
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o manager main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
WORKDIR /
COPY --from=builder /workspace/manager .
COPY config/internal config/internal
ENV SEALIGHTS_LOG_LEVEL=info
ENV SEALIGHTS_LAB_ID=openshift-ai-poc

CMD export SEALIGHTS_TOKEN="${SL_TOKEN}"

ENTRYPOINT ["/manager"]
