package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingPair represents the model of a pingpair
type PingPair struct {
	// Contains the request probe information.
	Request *PingProbe `json:"request" msgpack:"request" bson:"request" mapstructure:"request,omitempty"`

	// Contains the response probe information.
	Response *PingProbe `json:"response" msgpack:"response" bson:"response" mapstructure:"response,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingPair returns a new *PingPair
func NewPingPair() *PingPair {

	return &PingPair{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingPair) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingPair{}

	s.Request = o.Request
	s.Response = o.Response

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingPair) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingPair{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Request = s.Request
	o.Response = s.Response

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingPair) BleveType() string {

	return "pingpair"
}

// DeepCopy returns a deep copy if the PingPair.
func (o *PingPair) DeepCopy() *PingPair {

	if o == nil {
		return nil
	}

	out := &PingPair{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingPair.
func (o *PingPair) DeepCopyInto(out *PingPair) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingPair: %s", err))
	}

	*out = *target.(*PingPair)
}

// Validate valides the current information stored into the structure.
func (o *PingPair) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Request != nil {
		elemental.ResetDefaultForZeroValues(o.Request)
		if err := o.Request.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if o.Response != nil {
		elemental.ResetDefaultForZeroValues(o.Response)
		if err := o.Response.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*PingPair) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PingPairAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PingPairLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PingPair) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PingPairAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PingPair) ValueForAttribute(name string) interface{} {

	switch name {
	case "request":
		return o.Request
	case "response":
		return o.Response
	}

	return nil
}

// PingPairAttributesMap represents the map of attribute for PingPair.
var PingPairAttributesMap = map[string]elemental.AttributeSpecification{
	"Request": {
		AllowedChoices: []string{},
		BSONFieldName:  "request",
		ConvertedName:  "Request",
		Description:    `Contains the request probe information.`,
		Exposed:        true,
		Name:           "request",
		Stored:         true,
		SubType:        "pingprobe",
		Type:           "ref",
	},
	"Response": {
		AllowedChoices: []string{},
		BSONFieldName:  "response",
		ConvertedName:  "Response",
		Description:    `Contains the response probe information.`,
		Exposed:        true,
		Name:           "response",
		Stored:         true,
		SubType:        "pingprobe",
		Type:           "ref",
	},
}

// PingPairLowerCaseAttributesMap represents the map of attribute for PingPair.
var PingPairLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"request": {
		AllowedChoices: []string{},
		BSONFieldName:  "request",
		ConvertedName:  "Request",
		Description:    `Contains the request probe information.`,
		Exposed:        true,
		Name:           "request",
		Stored:         true,
		SubType:        "pingprobe",
		Type:           "ref",
	},
	"response": {
		AllowedChoices: []string{},
		BSONFieldName:  "response",
		ConvertedName:  "Response",
		Description:    `Contains the response probe information.`,
		Exposed:        true,
		Name:           "response",
		Stored:         true,
		SubType:        "pingprobe",
		Type:           "ref",
	},
}

type mongoAttributesPingPair struct {
	Request  *PingProbe `bson:"request"`
	Response *PingProbe `bson:"response"`
}
