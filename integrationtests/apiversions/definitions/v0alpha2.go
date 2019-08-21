package definitions

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/kubernetes-csi/csi-proxy/integrationtests/api"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

var v0alpha2 = apiversions.Definition{
	Version: apiversions.NewVersion("v0alpha2"),

	DeprecatedFields: func(requestMsg proto.Message) (deprecatedFields []apiversions.DeprecatedField) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			if request.Input32 != 0 {
				deprecatedFields = append(deprecatedFields, apiversions.DeprecatedField{
					DeprecationType: apiversions.SoftDeprecation,
					FieldName:       "input32",
					FieldValue:      request.Input32,
					Message:         "try using the new input64 field instead!",
				})
			}
		}

		return deprecatedFields
	},

	UpRequestTransformer: func(requestMsg proto.Message) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			if request.Input32 != 0 && request.Input64 == 0 {
				request.Input64 = int64(request.Input32)
			}
		}
	},

	DownRequestTransformer: func(requestMsg proto.Message) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			if request.Input64 != 0 && request.Input32 == 0 {
				request.Input32 = int32(request.Input64)
			}
		}
	},

	UpResponseTransformer: func(requestMsg, responseMsg proto.Message) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			response := responseMsg.(*pb.ComputeDoubleResponse)

			if request.Input64 != 0 {
				response.Response = int64(response.Response32)
			}

			abs := func(x int32) int32 {
				if x >= 0 {
					return x
				}
				return -x
			}
			if abs(response.Response32) > abs(int32(request.Input64)) {
				response.Success = true
			} else {
				response.Success = false
				// TODO wkpo re-use same error message!
			}
		}
	},

	DownResponseTransformer: func(requestMsg, responseMsg proto.Message) {
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
