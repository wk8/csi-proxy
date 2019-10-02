package v1

import (
	unsafe "unsafe"

	v1 "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1"
	internal "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
)

func autoConvert_v1_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(in *v1.ComputeDoubleRequest, out *internal.ComputeDoubleRequest) error {
	out.Input64 = in.Input64
	return nil
}

func autoConvert_internal_ComputeDoubleRequest_To_v1_ComputeDoubleRequest(in *internal.ComputeDoubleRequest, out *v1.ComputeDoubleRequest) error {
	out.Input64 = in.Input64
	return nil
}

func autoConvert_v1_ComputeDoubleResponse_To_internal_ComputeDoubleResponse(in *v1.ComputeDoubleResponse, out *internal.ComputeDoubleResponse) error {
	out.Response = in.Response
	return nil
}

func autoConvert_internal_ComputeDoubleResponse_To_v1_ComputeDoubleResponse(in *internal.ComputeDoubleResponse, out *v1.ComputeDoubleResponse) error {
	out.Response = in.Response
	return nil
}

func autoConvert_v1_TellMeAPoemRequest_To_internal_TellMeAPoemRequest(in *v1.TellMeAPoemRequest, out *internal.TellMeAPoemRequest) error {
	out.IWantATitle = in.IWantATitle
	return nil
}

func autoConvert_internal_TellMeAPoemRequest_To_v1_TellMeAPoemRequest(in *internal.TellMeAPoemRequest, out *v1.TellMeAPoemRequest) error {
	out.IWantATitle = in.IWantATitle
	return nil
}

func autoConvert_v1_TellMeAPoemResponse_To_internal_TellMeAPoemResponse(in *v1.TellMeAPoemResponse, out *internal.TellMeAPoemResponse) error {
	out.Title = in.Title
	out.Lines = *(*[]string)(unsafe.Pointer(&in.Lines))
	return nil
}

func autoConvert_internal_TellMeAPoemResponse_To_v1_TellMeAPoemResponse(in *internal.TellMeAPoemResponse, out *v1.TellMeAPoemResponse) error {
	out.Title = in.Title
	out.Lines = *(*[]string)(unsafe.Pointer(&in.Lines))
	return nil
}
