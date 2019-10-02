package v1alpha2

import (
	v1alpha2 "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha2"
	internal "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
)

func autoConvert_v1alpha2_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(in *v1alpha2.ComputeDoubleRequest, out *internal.ComputeDoubleRequest) error {
	out.Input64 = in.Input64
	return nil
}

func autoConvert_internal_ComputeDoubleRequest_To_v1alpha2_ComputeDoubleRequest(in *internal.ComputeDoubleRequest, out *v1alpha2.ComputeDoubleRequest) error {
	out.Input64 = in.Input64
	return nil
}

func autoConvert_v1alpha2_ComputeDoubleResponse_To_internal_ComputeDoubleResponse(in *v1alpha2.ComputeDoubleResponse, out *internal.ComputeDoubleResponse) error {
	out.Response = in.Response
	return nil
}

func autoConvert_internal_ComputeDoubleResponse_To_v1alpha2_ComputeDoubleResponse(in *internal.ComputeDoubleResponse, out *v1alpha2.ComputeDoubleResponse) error {
	out.Response = in.Response
	return nil
}
