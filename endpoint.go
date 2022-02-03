package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// Endpoint represents the model of a endpoint
type Endpoint struct {
	// URI of the exposed API.
	URI string `json:"URI" msgpack:"URI" bson:"uri" mapstructure:"URI,omitempty"`

	// The scopes authorized to access the API.
	AllowedScopes [][]string `json:"allowedScopes" msgpack:"allowedScopes" bson:"allowedscopes" mapstructure:"allowedScopes,omitempty"`

	// Methods exposed to access the API.
	Methods []string `json:"methods" msgpack:"methods" bson:"methods" mapstructure:"methods,omitempty"`

	// If `true`, the API is public.
	Public bool `json:"public" msgpack:"public" bson:"public" mapstructure:"public,omitempty"`

	// Use `allowedScopes`.
	Scopes []string `json:"scopes" msgpack:"scopes" bson:"-" mapstructure:"scopes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEndpoint returns a new *Endpoint
func NewEndpoint() *Endpoint {

	return &Endpoint{
		ModelVersion:  1,
		AllowedScopes: [][]string{},
		Methods:       []string{},
		Scopes:        []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Endpoint) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEndpoint{}

	s.URI = o.URI
	s.AllowedScopes = o.AllowedScopes
	s.Methods = o.Methods
	s.Public = o.Public

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Endpoint) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEndpoint{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.URI = s.URI
	o.AllowedScopes = s.AllowedScopes
	o.Methods = s.Methods
	o.Public = s.Public

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *Endpoint) BleveType() string {

	return "endpoint"
}

// DeepCopy returns a deep copy if the Endpoint.
func (o *Endpoint) DeepCopy() *Endpoint {

	if o == nil {
		return nil
	}

	out := &Endpoint{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Endpoint.
func (o *Endpoint) DeepCopyInto(out *Endpoint) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Endpoint: %s", err))
	}

	*out = *target.(*Endpoint)
}

// Validate valides the current information stored into the structure.
func (o *Endpoint) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateHTTPMethods("methods", o.Methods); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Endpoint) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EndpointAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EndpointLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Endpoint) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EndpointAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Endpoint) ValueForAttribute(name string) interface{} {

	switch name {
	case "URI":
		return o.URI
	case "allowedScopes":
		return o.AllowedScopes
	case "methods":
		return o.Methods
	case "public":
		return o.Public
	case "scopes":
		return o.Scopes
	}

	return nil
}

// EndpointAttributesMap represents the map of attribute for Endpoint.
var EndpointAttributesMap = map[string]elemental.AttributeSpecification{
	"URI": {
		AllowedChoices: []string{},
		BSONFieldName:  "uri",
		ConvertedName:  "URI",
		Description:    `URI of the exposed API.`,
		Exposed:        true,
		Name:           "URI",
		Stored:         true,
		Type:           "string",
	},
	"AllowedScopes": {
		AllowedChoices: []string{},
		BSONFieldName:  "allowedscopes",
		ConvertedName:  "AllowedScopes",
		Description:    `The scopes authorized to access the API.`,
		Exposed:        true,
		Name:           "allowedScopes",
		Orderable:      true,
		Stored:         true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"Methods": {
		AllowedChoices: []string{},
		BSONFieldName:  "methods",
		ConvertedName:  "Methods",
		Description:    `Methods exposed to access the API.`,
		Exposed:        true,
		Name:           "methods",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Public": {
		AllowedChoices: []string{},
		BSONFieldName:  "public",
		ConvertedName:  "Public",
		Description:    `If ` + "`" + `true` + "`" + `, the API is public.`,
		Exposed:        true,
		Name:           "public",
		Stored:         true,
		Type:           "boolean",
	},
	"Scopes": {
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Deprecated:     true,
		Description:    `Use ` + "`" + `allowedScopes` + "`" + `.`,
		Exposed:        true,
		Name:           "scopes",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// EndpointLowerCaseAttributesMap represents the map of attribute for Endpoint.
var EndpointLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"uri": {
		AllowedChoices: []string{},
		BSONFieldName:  "uri",
		ConvertedName:  "URI",
		Description:    `URI of the exposed API.`,
		Exposed:        true,
		Name:           "URI",
		Stored:         true,
		Type:           "string",
	},
	"allowedscopes": {
		AllowedChoices: []string{},
		BSONFieldName:  "allowedscopes",
		ConvertedName:  "AllowedScopes",
		Description:    `The scopes authorized to access the API.`,
		Exposed:        true,
		Name:           "allowedScopes",
		Orderable:      true,
		Stored:         true,
		SubType:        "[][]string",
		Type:           "external",
	},
	"methods": {
		AllowedChoices: []string{},
		BSONFieldName:  "methods",
		ConvertedName:  "Methods",
		Description:    `Methods exposed to access the API.`,
		Exposed:        true,
		Name:           "methods",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"public": {
		AllowedChoices: []string{},
		BSONFieldName:  "public",
		ConvertedName:  "Public",
		Description:    `If ` + "`" + `true` + "`" + `, the API is public.`,
		Exposed:        true,
		Name:           "public",
		Stored:         true,
		Type:           "boolean",
	},
	"scopes": {
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Deprecated:     true,
		Description:    `Use ` + "`" + `allowedScopes` + "`" + `.`,
		Exposed:        true,
		Name:           "scopes",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesEndpoint struct {
	URI           string     `bson:"uri"`
	AllowedScopes [][]string `bson:"allowedscopes"`
	Methods       []string   `bson:"methods"`
	Public        bool       `bson:"public"`
}
