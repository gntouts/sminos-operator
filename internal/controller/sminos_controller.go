/*
Copyright 2024.

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

package controller

import (
	"context"

	cachev1alpha1 "github.com/gntouts/sminos-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// SminosReconciler reconciles a Sminos object
type SminosReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Config   *rest.Config // Add this field
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=cache.nubificus.co.uk,resources=sminos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.nubificus.co.uk,resources=sminos/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cache.nubificus.co.uk,resources=sminos/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Sminos object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *SminosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	Log := log.FromContext(ctx)
	// Create a client for the apiextensions API group
	clientset, err := apiextensionsclient.NewForConfig(r.Config)
	if err != nil {
		Log.Error(err, "Failed to create apiextensions client")
		return ctrl.Result{}, err
	}

	// List the CRDs in the cluster
	crds, err := clientset.ApiextensionsV1().CustomResourceDefinitions().List(ctx, metav1.ListOptions{})
	if err != nil {
		Log.Error(err, "Failed to list CRDs")
		return ctrl.Result{}, err
	}

	// Log the names of the CRDs
	Log.Info("Existing CRDs in the cluster:")
	for _, crd := range crds.Items {
		Log.Info("CRD", "Name", crd.Name)
	}

	// TODO(user): Add your reconciliation logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SminosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.Config = mgr.GetConfig() // Assign the cluster configuration
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.Sminos{}).
		Owns(&appsv1.Deployment{}).
		Watches(&appsv1.Deployment{}, &MyEventHandler{}).
		Complete(r)
}

type MyEventHandler struct{}

// Create handles create events.
func (d *MyEventHandler) Create(ctx context.Context, e event.TypedCreateEvent[client.Object], q workqueue.RateLimitingInterface) {
	Log := log.FromContext(ctx)
	Log.Info("Create event received for object: %v\n", e.Object.GetName())
	// Add the object to the workqueue for reconciliation
	q.Add(reconcile.Request{NamespacedName: client.ObjectKeyFromObject(e.Object)})
}

// Update handles update events.
func (d *MyEventHandler) Update(ctx context.Context, e event.TypedUpdateEvent[client.Object], q workqueue.RateLimitingInterface) {
	Log := log.FromContext(ctx)
	Log.Info("Update event received for object: %v\n", e.ObjectNew.GetName())
	// Add the object to the workqueue for reconciliation
	q.Add(reconcile.Request{NamespacedName: client.ObjectKeyFromObject(e.ObjectNew)})
}

// Delete handles delete events.
func (d *MyEventHandler) Delete(ctx context.Context, e event.TypedDeleteEvent[client.Object], q workqueue.RateLimitingInterface) {
	Log := log.FromContext(ctx)
	Log.Info("Delete event received for object: %v\n", e.Object.GetName())
	// Add the object to the workqueue for reconciliation
	q.Add(reconcile.Request{NamespacedName: client.ObjectKeyFromObject(e.Object)})
}

// Generic handles generic events.
func (d *MyEventHandler) Generic(ctx context.Context, e event.TypedGenericEvent[client.Object], q workqueue.RateLimitingInterface) {
	Log := log.FromContext(ctx)
	Log.Info("Generic event received for object: %v\n", e.Object.GetName())
	// Add the object to the workqueue for reconciliation
	q.Add(reconcile.Request{NamespacedName: client.ObjectKeyFromObject(e.Object)})
}
