package apiversions

// TODO wkpo move to server pkg? or used in client too? and same for other files in this pkg

// TODO wkpo unit tests?

// TODO wkpo next break that up in client & server parts?

import (
	"fmt"
	"sort"

	"github.com/golang/protobuf/proto"
)

type Handler struct {
	// we keep definitions sorted
	definitions []Definition
	// versionIndexes maps version names to their index in definitions above.
	// This allows jumping to any given version.
	versionIndexes map[string]int
}

// Register adds a new definition to the handler.
// It only errors out if there are duplicates in the versions registered with the handler.
// This is not an especially efficient operation (n log(n) in the number of API versions
// we support) - but that's okay, since that only gets called at start time.
func (h *Handler) Register(definitions ...Definition) error {
	newDefinitions := append(h.definitions, definitions...)
	sort.Slice(h.definitions, func(i, j int) bool {
		return definitions[i].Version.Compare(definitions[j].Version) == Lesser
	})

	newVersionIndexes := make(map[string]int)
	for i, definition := range newDefinitions {
		if _, present := newVersionIndexes[definition.Version.String()]; present {
			return fmt.Errorf("duplicate versions: %v", definition.Version)
		}
		newVersionIndexes[definition.Version.String()] = i
	}

	h.definitions = newDefinitions
	h.versionIndexes = newVersionIndexes
	return nil
}

func (h *Handler) Definitions() []Definition {
	return h.definitions
}

type UnknownVersion struct {
	version Version
}

func (e UnknownVersion) Error() string {
	return fmt.Sprintf("unknown version: %v", e.version)
}

type RequestDoesNotExistAtThisVersion struct {
	version       Version
	introductedIn Version
}

func (e RequestDoesNotExistAtThisVersion) Error() string {
	return fmt.Sprintf("this request got introduced in version: %v > %v", e.introductedIn, e.version)
}

type DeprecatedRequest struct {
	version      Version
	deprecatedAt Version
}

func (e DeprecatedRequest) Error() string {
	return fmt.Sprintf("this request got deprecated in version: %v < %v", e.deprecatedAt, e.version)
}

// TODO wkpo comment?
func (h *Handler) WrapServerHandler(request proto.Message, version Version, serverHandler func() (proto.Message, error)) (proto.Message, error) {
	// TODO wkpo next
	softDeprecations, hardDeprecations, deprecatedRequest, err := h.HandleServerRequest(version, request)
	if err != nil {
		return nil, err
	}
	if hardDeprecations != nil {
		// TODO wkpo on pourrait pas juste faire une erreur en cas de hard deprecation?
		return nil, fmt.Errorf("wkpo bordel %v, %v", softDeprecations, deprecatedRequest)
	}
	// TODO wkpo log soft deprecations!

	response, err := serverHandler()
	if err != nil {
		return nil, err
	}

	err = h.HandleServerResponse(version, request, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// HandleServerRequest is meant to be called by the server with a top-level request message.
// 1. errors out if that API version is unknown (and returns an UnknownVersion error),
//    or if that request message did not exist yet for that API version (and returns a RequestDoesNotExistAtThisVersion error)
// 2. returns the lists of soft and hard deprecated fields for that API version and all earlier ones; note that it will
//    stop as soon as soon as it encounters a hard deprecation.
// 3. if there are no hard deprecations, then proceeds to run UpRequestTransformers in order.
// On top of deprecated fields, also returns a boolean indicating whether this request is marked as
// deprecated for this API version.
// TODO wkpo break up in smaller funcs?
func (h *Handler) HandleServerRequest(version Version, request proto.Message) (softDeprecations []DeprecatedField, hardDeprecations []DeprecatedField, deprecatedRequest bool, err error) {
	versionIndex, present := h.versionIndexes[version.String()]
	if !present {
		err = UnknownVersion{version: version}
		return
	}

	// does this request already exist at this API version?
	for i := versionIndex + 1; i < len(h.definitions); i++ {
		if h.definitions[i].IsNewRequest != nil && h.definitions[i].IsNewRequest(request) {
			err = RequestDoesNotExistAtThisVersion{
				version:       version,
				introductedIn: h.definitions[i].Version,
			}
			return
		}
	}

	// or has it been deprecated?
	for i := versionIndex; i > 0; i-- {
		if h.definitions[i].IsDeprecatedRequest == nil {
			continue
		}
		isDeprecated, deprecationType := h.definitions[i].IsDeprecatedRequest(request)

		if isDeprecated {
			switch deprecationType {
			case HardDeprecation:
				err = DeprecatedRequest{
					version:      version,
					deprecatedAt: h.definitions[i].Version,
				}
				return
			case SoftDeprecation:
				deprecatedRequest = true
			default:
				// there are no other types
				panic(fmt.Errorf("uknown deprecation type: %v", deprecationType))
			}
		}
	}

	// deprecated fields
	for i := versionIndex; i > 0; i-- {
		if h.definitions[i].DeprecatedFields == nil {
			continue
		}

		for _, deprecation := range h.definitions[i].DeprecatedFields(request) {
			switch deprecation.DeprecationType {
			case HardDeprecation:
				hardDeprecations = append(hardDeprecations, deprecation)
			case SoftDeprecation:
				softDeprecations = append(softDeprecations, deprecation)
			default:
				// there are no other types
				panic(fmt.Errorf("uknown deprecation type: %v", deprecation.DeprecationType))
			}
		}

		if len(hardDeprecations) > 0 {
			return
		}
	}

	// run UpRequestTransformers
	for i := versionIndex + 1; i < len(h.definitions); i++ {
		if h.definitions[i].UpRequestTransformer != nil {
			h.definitions[i].UpRequestTransformer(request)
		}
	}

	return
}

// HandleServerResponse is meant to be called by the server with top-level request and
// response message of corresponding types and runs DownResponseTransformers in reverse
// chronological order.
func (h *Handler) HandleServerResponse(version Version, request, response proto.Message) error {
	versionIndex, present := h.versionIndexes[version.String()]
	if !present {
		return UnknownVersion{version: version}
	}

	for i := len(h.definitions); i > versionIndex; i-- {
		if h.definitions[i].DownResponseTransformer != nil {
			h.definitions[i].DownResponseTransformer(request, response)
		}
	}

	return nil
}

// HandleClientRequest is meant to be called by the client with a top-level request message;
// it runs DownRequestTransformers in reverse chronological order to adapt it to the API version
// the server is running.
func (h *Handler) HandleClientRequest(version Version, request proto.Message) error {
	versionIndex, present := h.versionIndexes[version.String()]
	if !present {
		return UnknownVersion{version: version}
	}

	for i := len(h.definitions); i > versionIndex; i-- {
		if h.definitions[i].DownRequestTransformer != nil {
			h.definitions[i].DownRequestTransformer(request)
		}
	}

	return nil
}

// HandleClientResponse is meant to be called by the client with top-level request and
// response message of corresponding types and runs UpResponseTransformers in chronological
// order to adapt the response from the API version the server is running.
func (h *Handler) HandleClientResponse(version Version, request, response proto.Message) error {
	versionIndex, present := h.versionIndexes[version.String()]
	if !present {
		return UnknownVersion{version: version}
	}

	for i := versionIndex + 1; i < len(h.definitions); i++ {
		if h.definitions[i].UpResponseTransformer != nil {
			h.definitions[i].UpResponseTransformer(request, response)
		}
	}

	return nil
}
