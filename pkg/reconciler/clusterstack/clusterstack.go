package clusterstack

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/controller"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	v1alpha1Informers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha1"
	v1alpha1Listers "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha1"
	"github.com/pivotal/kpack/pkg/reconciler"
	"github.com/pivotal/kpack/pkg/registry"
)

const (
	ReconcilerName = "Stacks"
	Kind           = "Stack"
)

//go:generate counterfeiter . ClusterStackReader
type ClusterStackReader interface {
	Read(keychain *authn.Keychain, clusterStackSpec v1alpha1.ClusterStackSpec) (v1alpha1.ResolvedClusterStack, error)
}

func NewController(
	opt reconciler.Options,
	keychainFactory registry.KeychainFactory,
	clusterStackInformer v1alpha1Informers.ClusterStackInformer,
	clusterStackReader ClusterStackReader) *controller.Impl {
	c := &Reconciler{
		Client:             opt.Client,
		ClusterStackLister: clusterStackInformer.Lister(),
		ClusterStackReader: clusterStackReader,
		KeychainFactory:    keychainFactory,
	}
	impl := controller.NewImpl(c, opt.Logger, ReconcilerName)
	clusterStackInformer.Informer().AddEventHandler(reconciler.Handler(impl.Enqueue))
	return impl
}

type Reconciler struct {
	Client             versioned.Interface
	ClusterStackLister v1alpha1Listers.ClusterStackLister
	ClusterStackReader ClusterStackReader
	KeychainFactory    registry.KeychainFactory
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_, clusterStackName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	clusterStack, err := c.ClusterStackLister.Get(clusterStackName)
	if k8serrors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	clusterStack = clusterStack.DeepCopy()

	clusterStack, err = c.reconcileClusterStackStatus(ctx, clusterStack)

	updateErr := c.updateClusterStackStatus(ctx, clusterStack)
	if updateErr != nil {
		return updateErr
	}

	if err != nil {
		return controller.NewPermanentError(err)
	}
	return nil
}

func (c *Reconciler) reconcileClusterStackStatus(ctx context.Context, clusterStack *v1alpha1.ClusterStack) (*v1alpha1.ClusterStack, error) {
	var keychain *authn.Keychain
	var err error

	if clusterStack.Spec.ServiceAccountRef != nil {
		k, err := c.KeychainFactory.KeychainForSecretRef(ctx, registry.SecretRef{
			ServiceAccount: clusterStack.Spec.ServiceAccountRef.Name,
			Namespace:      clusterStack.Spec.ServiceAccountRef.Namespace,
		})
		keychain = &k
		if err != nil {
			clusterStack.Status = v1alpha1.ClusterStackStatus{
				Status: corev1alpha1.Status{
					ObservedGeneration: clusterStack.Generation,
					Conditions: corev1alpha1.Conditions{
						{
							Type:               corev1alpha1.ConditionReady,
							Status:             corev1.ConditionFalse,
							LastTransitionTime: corev1alpha1.VolatileTime{Inner: metav1.Now()},
							Message:            err.Error(),
						},
					},
				},
			}
			return clusterStack, err
		}

	}

	resolvedClusterStack, err := c.ClusterStackReader.Read(keychain, clusterStack.Spec)
	if err != nil {
		clusterStack.Status = v1alpha1.ClusterStackStatus{
			Status: corev1alpha1.Status{
				ObservedGeneration: clusterStack.Generation,
				Conditions: corev1alpha1.Conditions{
					{
						Type:               corev1alpha1.ConditionReady,
						Status:             corev1.ConditionFalse,
						LastTransitionTime: corev1alpha1.VolatileTime{Inner: metav1.Now()},
						Message:            err.Error(),
					},
				},
			},
		}
		return clusterStack, err
	}

	clusterStack.Status = v1alpha1.ClusterStackStatus{
		Status: corev1alpha1.Status{
			ObservedGeneration: clusterStack.Generation,
			Conditions: corev1alpha1.Conditions{
				{
					LastTransitionTime: corev1alpha1.VolatileTime{Inner: metav1.Now()},
					Type:               corev1alpha1.ConditionReady,
					Status:             corev1.ConditionTrue,
				},
			},
		},
		ResolvedClusterStack: resolvedClusterStack,
	}
	return clusterStack, nil
}

func (c *Reconciler) updateClusterStackStatus(ctx context.Context, desired *v1alpha1.ClusterStack) error {
	desired.Status.ObservedGeneration = desired.Generation

	original, err := c.ClusterStackLister.Get(desired.Name)
	if err != nil {
		return err
	}

	if equality.Semantic.DeepEqual(desired.Status, original.Status) {
		return nil
	}

	_, err = c.Client.KpackV1alpha1().ClusterStacks().UpdateStatus(ctx, desired, metav1.UpdateOptions{})
	return err
}
