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

package v1

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.

func (r *DorisCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:unnamedwatches:path=/mutate-doris-selectdb-com-v1-doriscluster,mutating=true,failurePolicy=fail,sideEffects=None,groups=doris.selectdb.com,resources=dorisclusters,verbs=create;update;delete,versions=v1,name=mdoriscluster.kb.io,admissionReviewVersions=v1
var _ webhook.Defaulter = &DorisCluster{}

// Default implements webhook.Defaulter so a unnamedwatches will be registered for the type
func (r *DorisCluster) Default() {
	klog.Infof("mutatingwebhook mutate doriscluster name=%s.", r.Name)
	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:unnamedwatches:path=/validate-doris-selectdb-com-v1-doriscluster,mutating=false,failurePolicy=fail,sideEffects=None,groups=doris.selectdb.com,resources=dorisclusters,verbs=create;update,versions=v1,name=vdoriscluster.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &DorisCluster{}

// ValidateCreate implements webhook.Validator so a unnamedwatches will be registered for the type
func (r *DorisCluster) ValidateCreate() error {
	klog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a unnamedwatches will be registered for the type
func (r *DorisCluster) ValidateUpdate(old runtime.Object) error {
	klog.Info("validate update", "name", r.Name)
	var errors []error
	oldCluster := old.(*DorisCluster)

	//operableErr := checkClusterOperable(r)

	// fe follower not allowed scale down
	if *oldCluster.Spec.FeSpec.ElectionNumber > *r.Spec.FeSpec.ElectionNumber {
		errors = append(errors, fmt.Errorf("'FeSpec.ElectionNumber' error: scale down in the number of ElectionNumber are not allowed"))
	}

	// fe FeSpec.Replicas must greater than or equal to FeSpec.ElectionNumber
	if *r.Spec.FeSpec.Replicas < *r.Spec.FeSpec.ElectionNumber {
		errors = append(errors, fmt.Errorf("'FeSpec.Replicas' error: changes in the number of FeSpec.Replicas must greater than or equal to FeSpec.ElectionNumber"))
	}

	if oldCluster.Status.ClusterPhase.Phase != PHASE_OPERABLE && oldCluster.Status.ClusterPhase.Retry != RETRY_OPERATOR_FE && *r.Spec.FeSpec.Replicas != *oldCluster.Spec.FeSpec.Replicas {
		errors = append(errors, fmt.Errorf("ClusterOperationalConflicts error: there is a conflict in CRD modify. currently, cluster Phase is %+v ", oldCluster.Status.ClusterPhase))
	}

	if oldCluster.Status.ClusterPhase.Phase != PHASE_OPERABLE &&
		oldCluster.Status.ClusterPhase.Retry != RETRY_OPERATOR_BE &&
		*r.Spec.BeSpec.Replicas != *oldCluster.Spec.BeSpec.Replicas {
		errors = append(errors, fmt.Errorf("there is a conflict in CRD modify. currently, cluster Phase is %+v ", oldCluster.Status.ClusterPhase))
	}

	if len(errors) != 0 {
		return kerrors.NewAggregate(errors)
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a unnamedwatches will be registered for the type
func (r *DorisCluster) ValidateDelete() error {
	klog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
