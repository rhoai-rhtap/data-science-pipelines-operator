/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package systemtests

import (
	"context"
	"flag"
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"testing"
	"time"
)

const (
	DSPAtimeout = time.Second * 240
	interval    = time.Millisecond * 2
)

var (
	loggr            logr.Logger
	ctx              context.Context
	cfg              *rest.Config
	cancel           context.CancelFunc
	clientmgr        ClientManager
	kubeconfig       string
	k8sApiServerHost string
	DSPAPath         string
	DSPANamespace    string
)

type ClientManager struct {
	k8sClient client.Client
}

// TestAPIs - This is the entry point for Ginkgo -
// the go test runner will run this function when you run go test or ginkgo.
// Under the hood, Ginkgo is simply calling go test.
// You can run go test instead of the ginkgo CLI, But Ginkgo has several capabilities that can only be accessed via ginkgo.
// It is best practice to embrace the ginkgo CLI and treat it as a first-class member of the testing toolchain.
func TestAPIs(t *testing.T) {
	// Single line of glue code connecting Ginkgo to Gomega
	// Inform our matcher library (Gomega) which function to call (Ginkgo's Fail) in the event a failure is detected.
	RegisterFailHandler(Fail)

	// Inform Ginkgo to start the test suite, passing it the *testing.T instance and a description of the suite.
	// Only call RunSpecs once and let Ginkgo worry about calling *testing.T for us.
	RunSpecs(t, "Controller Suite")
}

// Register flags in an init function. This ensures they are registered _before_ `go test` calls flag.Parse()
func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "~/.kube/config", "The path to the kubeconfig.")
	flag.StringVar(&k8sApiServerHost, "k8sApiServerHost", "localhost:6443", "The k8s cluster api server host.")
	flag.StringVar(&DSPAPath, "DSPAPath", "Path to DSPA", "The DSP resource file to deploy for testing.")
	flag.StringVar(&DSPANamespace, "DSPANamespace", "Namespace to deploy DSPA", "The namespace to deploy DSPA.")
}

var _ = BeforeSuite(func() {
	ctx, cancel = context.WithCancel(context.TODO())

	// Initialize logger
	opts := zap.Options{
		Development: true,
		TimeEncoder: zapcore.TimeEncoderOfLayout(time.RFC3339),
	}
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseFlagOptions(&opts)))

	loggr = logf.Log
	var err error

	clientmgr = ClientManager{}

	// Set up client auth configs
	cfg, err = clientcmd.BuildConfigFromFlags(k8sApiServerHost, kubeconfig)
	Expect(err).ToNot(HaveOccurred())

	// Initialize Kubernetes client
	clientmgr.k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).ToNot(HaveOccurred())

})

var _ = BeforeEach(func() {
})

var _ = AfterSuite(func() {
})
