package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudLoadBalancerRouteData represents the model of a cloudloadbalancerroutedata
type CloudLoadBalancerRouteData struct {
	// The port for the target group.
	Port string `json:"port" msgpack:"port" bson:"-" mapstructure:"port,omitempty"`

	// The ID for the target group.
	TargetGroupID string `json:"targetGroupID" msgpack:"targetGroupID" bson:"targetgroupid" mapstructure:"targetGroupID,omitempty"`

	// The ID of the target object.
	TargetID string `json:"targetID" msgpack:"targetID" bson:"targetid" mapstructure:"targetID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudLoadBalancerRouteData returns a new *CloudLoadBalancerRouteData
func NewCloudLoadBalancerRouteData() *CloudLoadBalancerRouteData {

	return &CloudLoadBalancerRouteData{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerRouteData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudLoadBalancerRouteData{}

	s.TargetGroupID = o.TargetGroupID
	s.TargetID = o.TargetID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerRouteData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudLoadBalancerRouteData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.TargetGroupID = s.TargetGroupID
	o.TargetID = s.TargetID

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudLoadBalancerRouteData) BleveType() string {

	return "cloudloadbalancerroutedata"
}

// DeepCopy returns a deep copy if the CloudLoadBalancerRouteData.
func (o *CloudLoadBalancerRouteData) DeepCopy() *CloudLoadBalancerRouteData {

	if o == nil {
		return nil
	}

	out := &CloudLoadBalancerRouteData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudLoadBalancerRouteData.
func (o *CloudLoadBalancerRouteData) DeepCopyInto(out *CloudLoadBalancerRouteData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudLoadBalancerRouteData: %s", err))
	}

	*out = *target.(*CloudLoadBalancerRouteData)
}

// Validate valides the current information stored into the structure.
func (o *CloudLoadBalancerRouteData) Validate() error {

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
func (*CloudLoadBalancerRouteData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudLoadBalancerRouteDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudLoadBalancerRouteDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudLoadBalancerRouteData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudLoadBalancerRouteDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudLoadBalancerRouteData) ValueForAttribute(name string) interface{} {

	switch name {
	case "port":
		return o.Port
	case "targetGroupID":
		return o.TargetGroupID
	case "targetID":
		return o.TargetID
	}

	return nil
}

// CloudLoadBalancerRouteDataAttributesMap represents the map of attribute for CloudLoadBalancerRouteData.
var CloudLoadBalancerRouteDataAttributesMap = map[string]elemental.AttributeSpecification{
	"Port": {
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		Description:    `The port for the target group.`,
		Exposed:        true,
		Name:           "port",
		Type:           "string",
	},
	"TargetGroupID": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetgroupid",
		ConvertedName:  "TargetGroupID",
		Description:    `The ID for the target group.`,
		Exposed:        true,
		Name:           "targetGroupID",
		Stored:         true,
		Type:           "string",
	},
	"TargetID": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetid",
		ConvertedName:  "TargetID",
		Description:    `The ID of the target object.`,
		Exposed:        true,
		Name:           "targetID",
		Stored:         true,
		Type:           "string",
	},
}

// CloudLoadBalancerRouteDataLowerCaseAttributesMap represents the map of attribute for CloudLoadBalancerRouteData.
var CloudLoadBalancerRouteDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"port": {
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		Description:    `The port for the target group.`,
		Exposed:        true,
		Name:           "port",
		Type:           "string",
	},
	"targetgroupid": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetgroupid",
		ConvertedName:  "TargetGroupID",
		Description:    `The ID for the target group.`,
		Exposed:        true,
		Name:           "targetGroupID",
		Stored:         true,
		Type:           "string",
	},
	"targetid": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetid",
		ConvertedName:  "TargetID",
		Description:    `The ID of the target object.`,
		Exposed:        true,
		Name:           "targetID",
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudLoadBalancerRouteData struct {
	TargetGroupID string `bson:"targetgroupid"`
	TargetID      string `bson:"targetid"`
}
