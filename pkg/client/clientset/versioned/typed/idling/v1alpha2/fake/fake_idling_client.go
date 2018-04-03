package fake

import (
	v1alpha2 "github.com/openshift/service-idler/pkg/client/clientset/versioned/typed/idling/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIdlingV1alpha2 struct {
	*testing.Fake
}

func (c *FakeIdlingV1alpha2) Idlers(namespace string) v1alpha2.IdlerInterface {
	return &FakeIdlers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIdlingV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
