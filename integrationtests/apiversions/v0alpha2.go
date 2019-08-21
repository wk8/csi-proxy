package apiversions

import (
	"github.com/golang/protobuf/proto"

	pb "github.com/kubernetes-csi/csi-proxy/integrationtests/api"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

var v0alpha2 = apiversions.Definition{
	Version: apiversions.NewVersion("v0alpha2"),

	Deprecations: func(requestMsg proto.Message) (deprecations []apiversions.Deprecation) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			if request.Input32 != 0 {
				deprecations = append(deprecations, apiversions.Deprecation{
					DeprecationType: apiversions.SoftDeprecation,
					FieldName:       "input32",
					FieldValue:      request.Input32,
					Message:         "try using the new input64 field instead!",
				})
			}
		}

		return deprecations
	},

	RequestTransformer: func(requestMsg proto.Message) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			if request.Input32 != 0 && request.Input64 == 0 {
				request.Input64 = int64(request.Input32)
			}
		}
	},

	ResponseTransformer: func(requestMsg, responseMsg proto.Message) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			response := responseMsg.(*pb.ComputeDoubleResponse)

			if request.Input32 != 0 {
				response.Response32 = int32(response.Response)
			}

			response.Success = false
			response.ErrorMessage = ""
		}
	},
}
