/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - PaaS 平台 (BlueKing - PaaS System) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package controllers

import (
	"context"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	paasv1alpha1 "bk.tencent.com/paas-app-operator/api/v1alpha1"
	paasv1alpha2 "bk.tencent.com/paas-app-operator/api/v1alpha2"

	autoscaling "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-general-pod-autoscaler/pkg/apis/autoscaling/v1alpha1"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var (
	cfg       *rest.Config
	k8sClient client.Client
	testEnv   *envtest.Environment
	ctx       context.Context
	cancel    context.CancelFunc
	scheme    *runtime.Scheme
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	ctx, cancel = context.WithCancel(context.Background())
	scheme = runtime.NewScheme()

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: true,
	}

	var err error
	// cfg is defined in this file globally.
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	Expect(paasv1alpha1.AddToScheme(scheme)).NotTo(HaveOccurred())
	Expect(paasv1alpha2.AddToScheme(scheme)).NotTo(HaveOccurred())
	Expect(corev1.AddToScheme(scheme)).NotTo(HaveOccurred())
	Expect(appsv1.AddToScheme(scheme)).NotTo(HaveOccurred())
	Expect(networkingv1.AddToScheme(scheme)).NotTo(HaveOccurred())
	Expect(autoscaling.AddToScheme(scheme)).NotTo(HaveOccurred())

	//+kubebuilder:scaffold:scheme

	/*
		A client is created for our test CRUD operations.
	*/
	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	/*
		One thing that this autogenerated file is missing, however, is a way to actually start your controller.
		The code above will set up a client for interacting with your custom Kind,
		but will not be able to test your controller behavior.
		If you want to test your custom controller logic, you’ll need to add some familiar-looking manager logic
		to your BeforeSuite() function, so you can register your custom controller to run on this test cluster.
		You may notice that the code below runs your controller with nearly identical logic to your CronJob project’s main.go!
		The only difference is that the manager is started in a separate goroutine so it does not block the cleanup of envtest
		when you’re done running your tests.
		Note that we set up both a "live" k8s client and a separate client from the manager. This is because when making
		assertions in tests, you generally want to assert against the live state of the API server. If you use the client
		from the manager (`k8sManager.GetClient`), you'd end up asserting against the contents of the cache instead, which is
		slower and can introduce flakiness into your tests. We could use the manager's `APIReader` to accomplish the same
		thing, but that would leave us with two clients in our test assertions and setup (one for reading, one for writing),
		and it'd be easy to make mistakes.
		Note that we keep the reconciler running against the manager's cache client, though -- we want our controller to
		behave as it would in production, and we use features of the cache (like indicies) in our controller which aren't
		available when talking directly to the API server.
	*/
	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	err = NewBkAppReconciler(
		k8sManager.GetClient(), k8sManager.GetScheme(),
	).SetupWithManager(ctx, k8sManager, controller.Options{})
	Expect(err).ToNot(HaveOccurred())

	err = NewDomainGroupMappingReconciler(
		k8sManager.GetClient(), k8sManager.GetScheme(),
	).SetupWithManager(ctx, k8sManager, controller.Options{})
	Expect(err).ToNot(HaveOccurred())

	go func() {
		defer GinkgoRecover()
		err = k8sManager.Start(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to run manager")
	}()
})

var _ = AfterSuite(func() {
	cancel()
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})
