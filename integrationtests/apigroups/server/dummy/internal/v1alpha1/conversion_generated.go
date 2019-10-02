package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1"
	internal "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
)

func autoConvert_v1alpha1_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(in *v1alpha1.ComputeDoubleRequest, out *internal.ComputeDoubleRequest) error {
	return nil
}

func autoConvert_internal_ComputeDoubleRequest_To_v1alpha1_ComputeDoubleRequest(in *internal.ComputeDoubleRequest, out *v1alpha1.ComputeDoubleRequest) error {
	return nil
}

func autoConvert_v1alpha1_ComputeDoubleResponse_To_internal_ComputeDoubleResponse(in *v1alpha1.ComputeDoubleResponse, out *internal.ComputeDoubleResponse) error {
	return nil
}

func autoConvert_internal_ComputeDoubleResponse_To_v1alpha1_ComputeDoubleResponse(in *internal.ComputeDoubleResponse, out *v1alpha1.ComputeDoubleResponse) error {
	return nil
}
