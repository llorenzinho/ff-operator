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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	deploymentv1alpha1 "github.com/llorenzinho/ff-operator/api/v1alpha1"
)

// FeatureFlagReconciler reconciles a FeatureFlag object
type FeatureFlagReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=deployment.github.com,resources=featureflags,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=deployment.github.com,resources=featureflags/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=deployment.github.com,resources=featureflags/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FeatureFlag object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *FeatureFlagReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx, "TraceID", ctx.Value("trace_id"))

	logger.Info("Reconciling FeatureFlag", "name", req.NamespacedName)

	// Get the actual object
	featureFlag := &deploymentv1alpha1.FeatureFlag{}
	if err := r.Get(ctx, req.NamespacedName, featureFlag); err != nil {
		logger.Error(err, "unable to fetch FeatureFlag, ignoring")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Set the status of the FF based on the actual state
	if featureFlag.Spec.Enabled {
		desiredStatus := featureFlag.Spec.Status
		featureFlag.Status.Active = desiredStatus
	} else {
		featureFlag.Status.Active = false
	}

	// Update the status of the FF
	if err := r.Status().Update(ctx, featureFlag); err != nil {
		logger.Error(err, "unable to update FeatureFlag status")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FeatureFlagReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&deploymentv1alpha1.FeatureFlag{}).
		Complete(r)
}
