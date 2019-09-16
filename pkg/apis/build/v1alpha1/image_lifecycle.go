package v1alpha1

import (
	"fmt"

	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

const (
	BuilderNotFound = "BuilderNotFound"
	BuilderNotReady = "BuilderNotReady"
)

func (im *Image) BuilderNotFound() duckv1alpha1.Conditions {
	return duckv1alpha1.Conditions{
		{
			Type:    duckv1alpha1.ConditionReady,
			Status:  corev1.ConditionFalse,
			Reason:  BuilderNotFound,
			Message: fmt.Sprintf("Unable to find builder %s.", im.Spec.Builder.Name),
		},
	}
}

func (im *Image) BuilderNotReady() duckv1alpha1.Conditions {
	return duckv1alpha1.Conditions{
		{
			Type:    duckv1alpha1.ConditionReady,
			Status:  corev1.ConditionFalse,
			Reason:  BuilderNotReady,
			Message: fmt.Sprintf("Builder %s is not ready.", im.Spec.Builder.Name),
		},
	}
}

func (b *Build) BuilderNotFound() duckv1alpha1.Conditions {
	return duckv1alpha1.Conditions{
		{
			Type:    duckv1alpha1.ConditionSucceeded,
			Status:  corev1.ConditionFalse,
			Reason:  BuilderNotFound,
			Message: fmt.Sprintf("Unable to find builder %s.", b.Spec.Builder.Image),
		},
	}
}
