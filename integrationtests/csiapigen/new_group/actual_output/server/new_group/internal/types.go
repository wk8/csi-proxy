package internal

import (
	v1alpha1 "github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/api/v1alpha1"
)

type FooRequest struct {
	SubMessage *v1alpha1.FooRequestSubMessage
	Blob       []byte
}

type FooRequestSubMessage struct {
	Blah          int32
	SubSubMessage []*v1alpha1.FooRequestSubSubMessage
}

type FooRequestSubSubMessage struct {
	Bools []bool
}

type FooResponse struct {
	Response32 int32
	SubMessage *v1alpha1.FooRequestSubMessage
}
