package v1alpha1

import (
	unsafe "unsafe"

	api "github.com/kubernetes-csi/csi-proxy/client/api"
	v1alpha1 "github.com/kubernetes-csi/csi-proxy/client/api/file_system/v1alpha1"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/file_system/internal"
)

func autoConvert_v1alpha1_PathExistsRequest_To_internal_PathExistsRequest(in *v1alpha1.PathExistsRequest, out *internal.PathExistsRequest) error {
	out.Path = in.Path
	return nil
}

func autoConvert_internal_PathExistsRequest_To_v1alpha1_PathExistsRequest(in *internal.PathExistsRequest, out *v1alpha1.PathExistsRequest) error {
	out.Path = in.Path
	return nil
}

func autoConvert_v1alpha1_PathExistsResponse_To_internal_PathExistsResponse(in *v1alpha1.PathExistsResponse, out *internal.PathExistsResponse) error {
	out.Success = in.Success
	out.CmdletError = (*api.CmdletError)(unsafe.Pointer(in.CmdletError))
	out.Exists = in.Exists
	return nil
}

func autoConvert_internal_PathExistsResponse_To_v1alpha1_PathExistsResponse(in *internal.PathExistsResponse, out *v1alpha1.PathExistsResponse) error {
	out.Success = in.Success
	out.CmdletError = (*api.CmdletError)(unsafe.Pointer(in.CmdletError))
	out.Exists = in.Exists
	return nil
}
