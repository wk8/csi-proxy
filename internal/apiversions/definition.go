package apiversions

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type DeprecationType int

const (
	SoftDeprecation DeprecationType = iota
	HardDeprecation
)

// TODO wkpo integration test on starting to use another field!!

type DeprecatedField struct {
	// if a soft deprecated field is used, a warning will be logged
	// if a hard deprecated field is used, the call will fail
	DeprecationType DeprecationType
	// This should be a fully qualified FieldName, starting from the message's root
	FieldName  string
	FieldValue interface{}
	// If this is a soft deprecation, can be used to communicate to the user e.g.
	// whether the field was ignored, or if they should start using another field for this;
	// If this is a hard deprecation, can be used to give the user a human-readable detailed
	// explanation.
	// In any case, this is optional.
	Message string
}

// TODO wkpo missing deprecated requests alltogether?
type Definition struct {
	Version Version

	BuildAndRegisterServers func(*grpc.Server)

	// IsNewRequest should return true iff handling this request was introduced in this version.
	// message will be one of the top-level request message.
	IsNewRequest func(requestMsg proto.Message) bool

	// IsDeprecatedRequest should return true iff we started deprecating this request in this version;
	// if it returns (true, SoftDeprecation), then we'll still process the request, with a warning.
	// it it returns (true, HardDeprecation), then the request won't be processed any more.
	IsDeprecatedRequest func(requestMsg proto.Message) (bool, DeprecationType)

	// DeprecatedFields should list deprecated fields used by the message, if any.
	// It should not transform the message in any way.
	// DeprecatedFields are guaranteed to be called in reverse chronological order.
	DeprecatedFields func(requestMsg proto.Message) []DeprecatedField

	// UpRequestTransformer should take a top-level request message from the previous API version
	// and do a best effort to turn it into a valid message for the newer API version;
	// e.g. by copying values from old fields to new ones.
	// This will be used by the server when receiving a request from an older API version.
	// This is really just a best effort, API handlers will still be told which API versions
	// a given message belongs to, and can thus handle breaking API changes this way.
	// UpRequestTransformers are guaranteed to be called in chronological order.
	UpRequestTransformer func(requestMsg proto.Message)

	// DownRequestTransformer should take a top-level request message from this API version
	// and do a best effort to turn it into a valid message for the previous API version;
	// e.g. by copying values from new fields to old ones or setting to their null values
	// fields introduced in this version.
	// This will be used by the client when sending a request to an older API version.
	// This is really just a best effort, the client will still be told which API version
	// the server is running, and can thus handle breaking API changes this way.
	// DownRequestTransformers are guaranteed to be called in reverse chronological order.
	DownRequestTransformer func(requestMsg proto.Message)

	// UpResponseTransformer should take a top-level response message from the previous API version
	// and do a best effort to turn it into a valid message for the newer API version;
	// e.g. by copying values from old fields to new ones.
	// This will be used by the client when receiving a response from a server running an older API version.
	// This is really just a best effort, the client will still be told which API version
	// the server is running, and can thus handle breaking API changes this way.
	// UpResponseTransformers are guaranteed to be called in chronological order.
	UpResponseTransformer func(requestMsg, responseMsg proto.Message)

	// DownResponseTransformer should take a top-level response message from this API version
	// and do a best effort to turn it into a valid message for the previous API version;
	// e.g. by copying values from new fields to old ones or setting to their null values
	// fields introduced in this version.
	// This will be used by the server when handling a request from an older API version.
	// This is really just a best effort, API handlers will still be told which API versions
	// a given message belongs to, and can thus handle breaking API changes this way.
	// DownResponseTransformers are guaranteed to be called in chronological order.
	DownResponseTransformer func(requestMsg, responseMsg proto.Message)
}
