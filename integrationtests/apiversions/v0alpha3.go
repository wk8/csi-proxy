package apiversions

import (
	"github.com/golang/protobuf/proto"
	"strconv"

	pb "github.com/kubernetes-csi/csi-proxy/integrationtests/api"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

var v0alpha3 = apiversions.Definition{
	Version: apiversions.NewVersion("v0alpha3"),

	IsNewRequest: func(requestMsg proto.Message) bool {
		switch requestMsg.(type) {
		case *pb.TellMeAPoemRequest:
			return true
		}
		return false
	},

	Deprecations: func(requestMsg proto.Message) (deprecations []apiversions.Deprecation) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			if request.Input64 != 0 {
				deprecations = append(deprecations, apiversions.Deprecation{
					DeprecationType: apiversions.SoftDeprecation,
					FieldName:       "input64",
					FieldValue:      request.Input64,
					Message:         "try using the new input field instead!",
				})
			}
		}

		return deprecations
	},

	RequestTransformer: func(requestMsg proto.Message) {
		switch request := requestMsg.(type) {
		case *pb.ComputeDoubleRequest:
			if request.Input64 != 0 && request.Input == "" {
				request.Input = strconv.FormatInt(request.Input64, 10)
			}
		}
	},
}
