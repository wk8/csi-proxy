package v1alpha1

import (
	"fmt"
	"math"

	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server"
)

func Convert_v1alpha1_ComputeDoubleRequest_To_server_ComputeDoubleRequest(in *ComputeDoubleRequest, out *server.ComputeDoubleRequest) error {
	out.Input64 = int64(in.Input32)
	return nil
}

func Convert_server_ComputeDoubleResponse_To_v1alpha1_ComputeDoubleResponse(in *server.ComputeDoubleResponse, out *ComputeDoubleResponse) error {
	i := in.Response
	if i > math.MaxInt32 || i < math.MinInt64 {
		return fmt.Errorf("overflow for %d", i)
	}
	out.Response32 = int32(i)
	return nil
}
