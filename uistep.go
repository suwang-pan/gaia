package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UIStep represents the model of a uistep
type UIStep struct {
	// Defines if the step is an advanced one.
	Advanced bool `json:"advanced" msgpack:"advanced" bson:"advanced" mapstructure:"advanced,omitempty"`

	// Description of the step.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Name of the step.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// List of parameters for this step.
	Parameters []*UIParameter `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUIStep returns a new *UIStep
func NewUIStep() *UIStep {

	return &UIStep{
		ModelVersion: 1,
		Parameters:   []*UIParameter{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UIStep) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesUIStep{}

	s.Advanced = o.Advanced
	s.Description = o.Description
	s.Name = o.Name
	s.Parameters = o.Parameters

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UIStep) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesUIStep{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Advanced = s.Advanced
	o.Description = s.Description
	o.Name = s.Name
	o.Parameters = s.Parameters

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *UIStep) BleveType() string {

	return "uistep"
}

// DeepCopy returns a deep copy if the UIStep.
func (o *UIStep) DeepCopy() *UIStep {

	if o == nil {
		return nil
	}

	out := &UIStep{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *UIStep.
func (o *UIStep) DeepCopyInto(out *UIStep) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy UIStep: %s", err))
	}

	*out = *target.(*UIStep)
}

// Validate valides the current information stored into the structure.
func (o *UIStep) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	for _, sub := range o.Parameters {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
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
func (*UIStep) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := UIStepAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return UIStepLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*UIStep) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return UIStepAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *UIStep) ValueForAttribute(name string) interface{} {

	switch name {
	case "advanced":
		return o.Advanced
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "parameters":
		return o.Parameters
	}

	return nil
}

// UIStepAttributesMap represents the map of attribute for UIStep.
var UIStepAttributesMap = map[string]elemental.AttributeSpecification{
	"Advanced": {
		AllowedChoices: []string{},
		BSONFieldName:  "advanced",
		ConvertedName:  "Advanced",
		Description:    `Defines if the step is an advanced one.`,
		Exposed:        true,
		Name:           "advanced",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the step.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the step.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Parameters": {
		AllowedChoices: []string{},
		BSONFieldName:  "parameters",
		ConvertedName:  "Parameters",
		Description:    `List of parameters for this step.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "uiparameter",
		Type:           "refList",
	},
}

// UIStepLowerCaseAttributesMap represents the map of attribute for UIStep.
var UIStepLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"advanced": {
		AllowedChoices: []string{},
		BSONFieldName:  "advanced",
		ConvertedName:  "Advanced",
		Description:    `Defines if the step is an advanced one.`,
		Exposed:        true,
		Name:           "advanced",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the step.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the step.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"parameters": {
		AllowedChoices: []string{},
		BSONFieldName:  "parameters",
		ConvertedName:  "Parameters",
		Description:    `List of parameters for this step.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "uiparameter",
		Type:           "refList",
	},
}

type mongoAttributesUIStep struct {
	Advanced    bool           `bson:"advanced"`
	Description string         `bson:"description"`
	Name        string         `bson:"name"`
	Parameters  []*UIParameter `bson:"parameters"`
}
