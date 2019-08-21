package apiversions

import (
	"github.com/golang/protobuf/proto"
)

type DeprecationType int

const (
	SoftDeprecation DeprecationType = iota
	HardDeprecation
)

// TODO wkpo integration test on starting to use another field!!

type Deprecation struct {
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
	message string
}

// TODO wkpo missing deprecated requests alltogether?
type Definition struct {
	Version Version

	// IsNewRequest should return true iff handling this request was introduced in this version.
	// message will be one of the top-level request message.
	IsNewRequest func(message *proto.Message) bool

	// IsDeprecatedRequest should return true iff we started deprecating this request in this version;
	// if it returns (true, SoftDeprecation), then we'll still process the request, with a warning.
	// it it returns (true, HardDeprecation), then the request won't be processed any more.
	IsDeprecatedRequest func(message *proto.Message) (bool, DeprecationType)

	// Deprecations should list deprecated fields used by the message, if any.
	// It should not transform the message in any way.
	// TODO wkpo enforce that ^ ?
	// Deprecations are guaranteed to be called in reverse chronological order.
	Deprecations func(*proto.Message) []Deprecation

	// ResponseTransformer should take a top-level request message, and transform it to fit older versions
	// requests; by e.g. resetting to their null value fields introduced in this version, or copying
	// values from old fields to new ones.
	// TODO wkpo vrai ca? the order?
	// RequestTransformers are guaranteed to be called in chronological order.
	RequestTransformer func(*proto.Message)

	// ResponseTransformer should take a top-level response message, and transform it to fit older versions
	// responses; by e.g. resetting to their null value fields introduced in this version, or copying
	// values from new fields to old ones.
	// TODO wkpo vrai ca? the order?
	// ResponseTransformers are guaranteed to be called in reverse chronological order.
	ResponseTransformer func(*proto.Message)
}
