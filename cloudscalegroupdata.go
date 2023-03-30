// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudScaleGroupData represents the model of a cloudscalegroupdata
type CloudScaleGroupData struct {
	// Availability Zones for the scale group.
	AvailabilityZones []string `json:"availabilityZones" msgpack:"availabilityZones" bson:"availabilityzones" mapstructure:"availabilityZones,omitempty"`

	// ID of associated instances with this scale group.
	Instances []string `json:"instances" msgpack:"instances" bson:"instances" mapstructure:"instances,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudScaleGroupData returns a new *CloudScaleGroupData
func NewCloudScaleGroupData() *CloudScaleGroupData {

	return &CloudScaleGroupData{
		ModelVersion:      1,
		AvailabilityZones: []string{},
		Instances:         []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudScaleGroupData) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudScaleGroupData{}

	s.AvailabilityZones = o.AvailabilityZones
	s.Instances = o.Instances

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudScaleGroupData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudScaleGroupData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AvailabilityZones = s.AvailabilityZones
	o.Instances = s.Instances

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudScaleGroupData) BleveType() string {

	return "cloudscalegroupdata"
}

// DeepCopy returns a deep copy if the CloudScaleGroupData.
func (o *CloudScaleGroupData) DeepCopy() *CloudScaleGroupData {

	if o == nil {
		return nil
	}

	out := &CloudScaleGroupData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudScaleGroupData.
func (o *CloudScaleGroupData) DeepCopyInto(out *CloudScaleGroupData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudScaleGroupData: %s", err))
	}

	*out = *target.(*CloudScaleGroupData)
}

// Validate valides the current information stored into the structure.
func (o *CloudScaleGroupData) Validate() error {

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
func (*CloudScaleGroupData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudScaleGroupDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudScaleGroupDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudScaleGroupData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudScaleGroupDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudScaleGroupData) ValueForAttribute(name string) any {

	switch name {
	case "availabilityZones":
		return o.AvailabilityZones
	case "instances":
		return o.Instances
	}

	return nil
}

// CloudScaleGroupDataAttributesMap represents the map of attribute for CloudScaleGroupData.
var CloudScaleGroupDataAttributesMap = map[string]elemental.AttributeSpecification{
	"AvailabilityZones": {
		AllowedChoices: []string{},
		BSONFieldName:  "availabilityzones",
		ConvertedName:  "AvailabilityZones",
		Description:    `Availability Zones for the scale group.`,
		Exposed:        true,
		Name:           "availabilityZones",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Instances": {
		AllowedChoices: []string{},
		BSONFieldName:  "instances",
		ConvertedName:  "Instances",
		Description:    `ID of associated instances with this scale group.`,
		Exposed:        true,
		Name:           "instances",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// CloudScaleGroupDataLowerCaseAttributesMap represents the map of attribute for CloudScaleGroupData.
var CloudScaleGroupDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"availabilityzones": {
		AllowedChoices: []string{},
		BSONFieldName:  "availabilityzones",
		ConvertedName:  "AvailabilityZones",
		Description:    `Availability Zones for the scale group.`,
		Exposed:        true,
		Name:           "availabilityZones",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"instances": {
		AllowedChoices: []string{},
		BSONFieldName:  "instances",
		ConvertedName:  "Instances",
		Description:    `ID of associated instances with this scale group.`,
		Exposed:        true,
		Name:           "instances",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesCloudScaleGroupData struct {
	AvailabilityZones []string `bson:"availabilityzones"`
	Instances         []string `bson:"instances"`
}