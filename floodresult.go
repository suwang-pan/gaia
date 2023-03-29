// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FloodResult represents the model of a floodresult
type FloodResult struct {
	// Errors that were encountered while flooding.
	Errors []string `json:"errors" msgpack:"errors" bson:"-" mapstructure:"errors,omitempty"`

	// Set if flooder found at least one trail.
	PathExists bool `json:"pathExists" msgpack:"pathExists" bson:"-" mapstructure:"pathExists,omitempty"`

	// A list of trails the flooder found. WARNING: this will eventually go away as we
	// should transmit the tree. We keep it this way for backwards compatibility with
	// existing code for the sake of speed.
	Trails [][]*DetachedDecoy `json:"trails" msgpack:"trails" bson:"-" mapstructure:"trails,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFloodResult returns a new *FloodResult
func NewFloodResult() *FloodResult {

	return &FloodResult{
		ModelVersion: 1,
		Errors:       []string{},
		Trails:       [][]*DetachedDecoy{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodResult) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFloodResult{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFloodResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *FloodResult) BleveType() string {

	return "floodresult"
}

// DeepCopy returns a deep copy if the FloodResult.
func (o *FloodResult) DeepCopy() *FloodResult {

	if o == nil {
		return nil
	}

	out := &FloodResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FloodResult.
func (o *FloodResult) DeepCopyInto(out *FloodResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FloodResult: %s", err))
	}

	*out = *target.(*FloodResult)
}

// Validate valides the current information stored into the structure.
func (o *FloodResult) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*FloodResult) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FloodResultAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FloodResultLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FloodResult) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FloodResultAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FloodResult) ValueForAttribute(name string) any {

	switch name {
	case "errors":
		return o.Errors
	case "pathExists":
		return o.PathExists
	case "trails":
		return o.Trails
	}

	return nil
}

// FloodResultAttributesMap represents the map of attribute for FloodResult.
var FloodResultAttributesMap = map[string]elemental.AttributeSpecification{
	"Errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `Errors that were encountered while flooding.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"PathExists": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PathExists",
		Description:    `Set if flooder found at least one trail.`,
		Exposed:        true,
		Name:           "pathExists",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"Trails": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Trails",
		Description: `A list of trails the flooder found. WARNING: this will eventually go away as we
should transmit the tree. We keep it this way for backwards compatibility with
existing code for the sake of speed.`,
		Exposed:  true,
		Name:     "trails",
		ReadOnly: true,
		SubType:  "[][]detacheddecoy",
		Type:     "external",
	},
}

// FloodResultLowerCaseAttributesMap represents the map of attribute for FloodResult.
var FloodResultLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `Errors that were encountered while flooding.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"pathexists": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PathExists",
		Description:    `Set if flooder found at least one trail.`,
		Exposed:        true,
		Name:           "pathExists",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"trails": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Trails",
		Description: `A list of trails the flooder found. WARNING: this will eventually go away as we
should transmit the tree. We keep it this way for backwards compatibility with
existing code for the sake of speed.`,
		Exposed:  true,
		Name:     "trails",
		ReadOnly: true,
		SubType:  "[][]detacheddecoy",
		Type:     "external",
	},
}

type mongoAttributesFloodResult struct {
}
