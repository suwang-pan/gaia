// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FloodNodeMakerConfig represents the model of a floodnodemakerconfig
type FloodNodeMakerConfig struct {
	// A list of cloud types involved in flooding.
	CloudTypes []string `json:"cloudTypes" msgpack:"cloudTypes" bson:"-" mapstructure:"cloudTypes,omitempty"`

	// A list of addresses which nodemakers will ignore when evaluating IP rules.
	OptionIgnoreIPRulesForGivenAddresses []string `json:"optionIgnoreIPRulesForGivenAddresses" msgpack:"optionIgnoreIPRulesForGivenAddresses" bson:"-" mapstructure:"optionIgnoreIPRulesForGivenAddresses,omitempty"`

	// If set, IP rule evaluation will only be considered if the target address is
	// fully covered by the IP rule.
	OptionSelectIPRulesWithFullCoverage bool `json:"optionSelectIPRulesWithFullCoverage" msgpack:"optionSelectIPRulesWithFullCoverage" bson:"-" mapstructure:"optionSelectIPRulesWithFullCoverage,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFloodNodeMakerConfig returns a new *FloodNodeMakerConfig
func NewFloodNodeMakerConfig() *FloodNodeMakerConfig {

	return &FloodNodeMakerConfig{
		ModelVersion:                         1,
		CloudTypes:                           []string{},
		OptionIgnoreIPRulesForGivenAddresses: []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodNodeMakerConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFloodNodeMakerConfig{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodNodeMakerConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFloodNodeMakerConfig{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *FloodNodeMakerConfig) BleveType() string {

	return "floodnodemakerconfig"
}

// DeepCopy returns a deep copy if the FloodNodeMakerConfig.
func (o *FloodNodeMakerConfig) DeepCopy() *FloodNodeMakerConfig {

	if o == nil {
		return nil
	}

	out := &FloodNodeMakerConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FloodNodeMakerConfig.
func (o *FloodNodeMakerConfig) DeepCopyInto(out *FloodNodeMakerConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FloodNodeMakerConfig: %s", err))
	}

	*out = *target.(*FloodNodeMakerConfig)
}

// Validate valides the current information stored into the structure.
func (o *FloodNodeMakerConfig) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("cloudTypes", o.CloudTypes); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*FloodNodeMakerConfig) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FloodNodeMakerConfigAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FloodNodeMakerConfigLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FloodNodeMakerConfig) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FloodNodeMakerConfigAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FloodNodeMakerConfig) ValueForAttribute(name string) interface{} {

	switch name {
	case "cloudTypes":
		return o.CloudTypes
	case "optionIgnoreIPRulesForGivenAddresses":
		return o.OptionIgnoreIPRulesForGivenAddresses
	case "optionSelectIPRulesWithFullCoverage":
		return o.OptionSelectIPRulesWithFullCoverage
	}

	return nil
}

// FloodNodeMakerConfigAttributesMap represents the map of attribute for FloodNodeMakerConfig.
var FloodNodeMakerConfigAttributesMap = map[string]elemental.AttributeSpecification{
	"CloudTypes": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudTypes",
		Description:    `A list of cloud types involved in flooding.`,
		Exposed:        true,
		Name:           "cloudTypes",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"OptionIgnoreIPRulesForGivenAddresses": {
		AllowedChoices: []string{},
		ConvertedName:  "OptionIgnoreIPRulesForGivenAddresses",
		Description:    `A list of addresses which nodemakers will ignore when evaluating IP rules.`,
		Exposed:        true,
		Name:           "optionIgnoreIPRulesForGivenAddresses",
		SubType:        "string",
		Type:           "list",
	},
	"OptionSelectIPRulesWithFullCoverage": {
		AllowedChoices: []string{},
		ConvertedName:  "OptionSelectIPRulesWithFullCoverage",
		Description: `If set, IP rule evaluation will only be considered if the target address is
fully covered by the IP rule.`,
		Exposed: true,
		Name:    "optionSelectIPRulesWithFullCoverage",
		Type:    "boolean",
	},
}

// FloodNodeMakerConfigLowerCaseAttributesMap represents the map of attribute for FloodNodeMakerConfig.
var FloodNodeMakerConfigLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"cloudtypes": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudTypes",
		Description:    `A list of cloud types involved in flooding.`,
		Exposed:        true,
		Name:           "cloudTypes",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"optionignoreiprulesforgivenaddresses": {
		AllowedChoices: []string{},
		ConvertedName:  "OptionIgnoreIPRulesForGivenAddresses",
		Description:    `A list of addresses which nodemakers will ignore when evaluating IP rules.`,
		Exposed:        true,
		Name:           "optionIgnoreIPRulesForGivenAddresses",
		SubType:        "string",
		Type:           "list",
	},
	"optionselectipruleswithfullcoverage": {
		AllowedChoices: []string{},
		ConvertedName:  "OptionSelectIPRulesWithFullCoverage",
		Description: `If set, IP rule evaluation will only be considered if the target address is
fully covered by the IP rule.`,
		Exposed: true,
		Name:    "optionSelectIPRulesWithFullCoverage",
		Type:    "boolean",
	},
}

type mongoAttributesFloodNodeMakerConfig struct {
}
