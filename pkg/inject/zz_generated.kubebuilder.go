// TODO: add apache license boilerplate here
package inject

import (
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	idlingv1alpha2 "github.com/openshift/service-idler/pkg/apis/idling/v1alpha2"
	rscheme "github.com/openshift/service-idler/pkg/client/clientset/versioned/scheme"
	"github.com/openshift/service-idler/pkg/controller/idler"
	"github.com/openshift/service-idler/pkg/inject/args"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&idlingv1alpha2.Idler{}, arguments.Informers.Idling().V1alpha2().Idlers()); err != nil {
			return err
		}

		// Add Kubernetes informers
		if err := arguments.ControllerManager.AddInformerProvider(&corev1.Endpoints{}, arguments.KubernetesInformers.Core().V1().Endpoints()); err != nil {
			return err
		}

		if c, err := idler.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &idlingv1alpha2.IdlerCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"idling.openshift.io"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{
			"",
		},
		Resources: []string{
			"endpoints",
		},
		Verbs: []string{
			"get", "list", "watch",
		},
	})
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{
			"",
		},
		Resources: []string{
			"events",
		},
		Verbs: []string{
			"create", "patch", "update",
		},
	})
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{
			"*",
		},
		Resources: []string{
			"*/scale",
		},
		Verbs: []string{
			"get", "update",
		},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "idling.openshift.io",
		Version: "v1alpha2",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
