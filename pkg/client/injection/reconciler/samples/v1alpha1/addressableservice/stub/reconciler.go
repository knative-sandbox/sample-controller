/*
Copyright 2020 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package addressableservice

import (
	"context"

	v1 "k8s.io/api/core/v1"
	reconciler "knative.dev/pkg/reconciler"
	v1alpha1 "knative.dev/sample-controller/pkg/apis/samples/v1alpha1"
	addressableservice "knative.dev/sample-controller/pkg/client/injection/reconciler/samples/v1alpha1/addressableservice"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// newReconciledNormal makes a new reconciler event with event type Normal, and
// reason AddressableServiceReconciled.
func newReconciledNormal(namespace, name string) reconciler.Event {
	return reconciler.NewEvent(v1.EventTypeNormal, "AddressableServiceReconciled", "AddressableService reconciled: \"%s/%s\"", namespace, name)
}

// Reconciler implements controller.Reconciler for AddressableService resources.
type Reconciler struct {
	// TODO: add additional requirements here.
}

// Check that our Reconciler implements Interface
var _ addressableservice.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *v1alpha1.AddressableService) reconciler.Event {
	if o.GetDeletionTimestamp() != nil {
		// Check for a DeletionTimestamp.  If present, elide the normal reconcile logic.
		// When a controller needs finalizer handling, it would go here.
		return nil
	}
	o.Status.InitializeConditions()

	// TODO: add custom reconciliation logic here.

	o.Status.ObservedGeneration = o.Generation
	return newReconciledNormal(o.Namespace, o.Name)
}
