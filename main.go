/*


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

package main

//
//import (
//	"flag"
//	"os"
//
//	"k8s.io/apimachinery/pkg/runtime"
//	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
//	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
//	ctrl "sigs.k8s.io/controller-runtime"
//	"sigs.k8s.io/controller-runtime/pkg/log/zap"
//
//	devopskubesphereiov1alpha3 "kubesphere.io/devops/pkg/api/devops/v1alpha3"
//	devopsv1alpha3 "kubesphere.io/devops/pkg/api/devops/v1alpha3"
//	"kubesphere.io/devops/controllers"
//	// +kubebuilder:scaffold:imports
//)
//
//var (
//	scheme   = runtime.NewScheme()
//	setupLog = ctrl.Log.WithName("setup")
//)
//
//func init() {
//	_ = clientgoscheme.AddToScheme(scheme)
//
//	_ = devopskubesphereiov1alpha3.AddToScheme(scheme)
//	_ = devopsv1alpha3.AddToScheme(scheme)
//	// +kubebuilder:scaffold:scheme
//}
//
//func main() {
//	var metricsAddr string
//	var enableLeaderElection bool
//	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
//	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
//		"Enable leader election for controller manager. "+
//			"Enabling this will ensure there is only one active controller manager.")
//	flag.Parse()
//
//	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))
//
//	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
//		Scheme:             scheme,
//		MetricsBindAddress: metricsAddr,
//		Port:               9443,
//		LeaderElection:     enableLeaderElection,
//		LeaderElectionID:   "cf6f1023.kubesphere.io",
//	})
//	if err != nil {
//		setupLog.Error(err, "unable to start manager")
//		os.Exit(1)
//	}
//
//	if err = (&controllers.PipelineReconciler{
//		Client: mgr.GetClient(),
//		Log:    ctrl.Log.WithName("controllers").WithName("Pipeline"),
//		Scheme: mgr.GetScheme(),
//	}).SetupWithManager(mgr); err != nil {
//		setupLog.Error(err, "unable to create controller", "controller", "Pipeline")
//		os.Exit(1)
//	}
//	if err = (&controllers.FakeReconciler{
//		Client: mgr.GetClient(),
//		Log:    ctrl.Log.WithName("controllers").WithName("Fake"),
//		Scheme: mgr.GetScheme(),
//	}).SetupWithManager(mgr); err != nil {
//		setupLog.Error(err, "unable to create controller", "controller", "Fake")
//		os.Exit(1)
//	}
//	// +kubebuilder:scaffold:builder
//
//	setupLog.Info("starting manager")
//	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
//		setupLog.Error(err, "problem running manager")
//		os.Exit(1)
//	}
//}
