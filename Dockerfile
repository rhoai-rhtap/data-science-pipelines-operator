# Build arguments
ARG SOURCE_CODE=.
ARG CI_CONTAINER_VERSION="unknown"

FROM registry.redhat.io/ubi8/go-toolset@sha256:4ec05fd5b355106cc0d990021a05b71bbfb9231e4f5bdc0c5316515edf6a1c96 as builder

ARG SOURCE_CODE

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
COPY config/internal config/internal
# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
USER root
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

FROM registry.redhat.io/ubi8/ubi-minimal:latest AS runtime

## Build args to be used at this step
ARG CI_CONTAINER_VERSION
ARG USER=65532


## Install additional packages
# TODO: is this needed?
RUN microdnf install -y shadow-utils &&\
    microdnf clean all

WORKDIR /
COPY --from=builder /workspace/manager .
COPY --from=builder /workspace/config/internal ./config/internal

## Create a non-root user with UID 65532 and switch to it
USER ${USER}:${USER}

ENTRYPOINT ["/manager"]
