// /*

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

package controllers

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/kedacore/http-add-on/operator/api/v1alpha1"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

// var cfg *rest.Config
// var k8sClient client.Client
// var testEnv *envtest.Environment

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}

type commonTestInfra struct {
	ns      string
	appName string
	ctx     context.Context
	cl      client.Client
	logger  logr.Logger
	httpso  v1alpha1.HTTPScaledObject
}

func newCommonTestInfra(namespace, appName string) *commonTestInfra {
	ctx := context.Background()
	cl := fake.NewFakeClient()
	logger := logr.Discard()

	httpso := v1alpha1.HTTPScaledObject{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      appName,
		},
		Spec: v1alpha1.HTTPScaledObjectSpec{
			ScaleTargetRef: &v1alpha1.ScaleTargetRef{
				Deployment: appName,
				Service:    appName,
				Port:       8081,
			},
			Replicas: v1alpha1.ReplicaStruct{
				Min: 0,
				Max: 20,
			},
		},
	}

	return &commonTestInfra{
		ns:      namespace,
		appName: appName,
		ctx:     ctx,
		cl:      cl,
		logger:  logger,
		httpso:  httpso,
	}
}

var _ = BeforeSuite(func(done Done) {
	// The commented code in this function connects to a test Kubernetes cluster.
	// We don't currently have tests that exercise functionality that needs a cluster,
	// so we're keeping it commented.
	logf.SetLogger(zap.New(func(opts *zap.Options) {
		opts.DestWriter = GinkgoWriter
	}))

	// By("bootstrapping test environment")
	// useExistingCluster := true
	// testEnv = &envtest.Environment{
	// 	CRDDirectoryPaths:  []string{filepath.Join("..", "config", "crd", "bases")},
	// 	UseExistingCluster: &useExistingCluster,
	// }

	var err error
	// cfg, err = testEnv.Start()
	// Expect(err).ToNot(HaveOccurred())
	// Expect(cfg).ToNot(BeNil())

	err = v1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme

	// k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	// Expect(err).ToNot(HaveOccurred())
	// Expect(k8sClient).ToNot(BeNil())

	close(done)
}, 60)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	// err := testEnv.Stop()
	// Expect(err).ToNot(HaveOccurred())
})
