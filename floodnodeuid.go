// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FloodNodeUID represents the model of a floodnodeuid
type FloodNodeUID struct {
	// The network address of the node described by the NodeUID.
	NetworkAddress string `json:"networkAddress,omitempty" msgpack:"networkAddress,omitempty" bson:"-" mapstructure:"networkAddress,omitempty"`

	// The node identifier.
	NodeID string `json:"nodeID" msgpack:"nodeID" bson:"-" mapstructure:"nodeID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFloodNodeUID returns a new *FloodNodeUID
func NewFloodNodeUID() *FloodNodeUID {

	return &FloodNodeUID{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodNodeUID) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFloodNodeUID{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FloodNodeUID) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFloodNodeUID{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *FloodNodeUID) BleveType() string {

	return "floodnodeuid"
}

// DeepCopy returns a deep copy if the FloodNodeUID.
func (o *FloodNodeUID) DeepCopy() *FloodNodeUID {

	if o == nil {
		return nil
	}

	out := &FloodNodeUID{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FloodNodeUID.
func (o *FloodNodeUID) DeepCopyInto(out *FloodNodeUID) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FloodNodeUID: %s", err))
	}

	*out = *target.(*FloodNodeUID)
}

// Validate valides the current information stored into the structure.
func (o *FloodNodeUID) Validate() error {

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
func (*FloodNodeUID) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FloodNodeUIDAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FloodNodeUIDLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FloodNodeUID) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FloodNodeUIDAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FloodNodeUID) ValueForAttribute(name string) any {

	switch name {
	case "networkAddress":
		return o.NetworkAddress
	case "nodeID":
		return o.NodeID
	}

	return nil
}

// FloodNodeUIDAttributesMap represents the map of attribute for FloodNodeUID.
var FloodNodeUIDAttributesMap = map[string]elemental.AttributeSpecification{
	"NetworkAddress": {
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAddress",
		Description:    `The network address of the node described by the NodeUID.`,
		Exposed:        true,
		Name:           "networkAddress",
		Type:           "string",
	},
	"NodeID": {
		AllowedChoices: []string{},
		ConvertedName:  "NodeID",
		Description:    `The node identifier.`,
		Exposed:        true,
		Name:           "nodeID",
		Type:           "string",
	},
}

// FloodNodeUIDLowerCaseAttributesMap represents the map of attribute for FloodNodeUID.
var FloodNodeUIDLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"networkaddress": {
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAddress",
		Description:    `The network address of the node described by the NodeUID.`,
		Exposed:        true,
		Name:           "networkAddress",
		Type:           "string",
	},
	"nodeid": {
		AllowedChoices: []string{},
		ConvertedName:  "NodeID",
		Description:    `The node identifier.`,
		Exposed:        true,
		Name:           "nodeID",
		Type:           "string",
	},
}

type mongoAttributesFloodNodeUID struct {
}
