package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudEndpointDataProductInfo represents the model of a cloudendpointdataproductinfo
type CloudEndpointDataProductInfo struct {
	// The ID of the corresponding product.
	ProductID string `json:"productID,omitempty" msgpack:"productID,omitempty" bson:"productid,omitempty" mapstructure:"productID,omitempty"`

	// The type of the product.
	Type string `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudEndpointDataProductInfo returns a new *CloudEndpointDataProductInfo
func NewCloudEndpointDataProductInfo() *CloudEndpointDataProductInfo {

	return &CloudEndpointDataProductInfo{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointDataProductInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudEndpointDataProductInfo{}

	s.ProductID = o.ProductID
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointDataProductInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudEndpointDataProductInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ProductID = s.ProductID
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudEndpointDataProductInfo) BleveType() string {

	return "cloudendpointdataproductinfo"
}

// DeepCopy returns a deep copy if the CloudEndpointDataProductInfo.
func (o *CloudEndpointDataProductInfo) DeepCopy() *CloudEndpointDataProductInfo {

	if o == nil {
		return nil
	}

	out := &CloudEndpointDataProductInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudEndpointDataProductInfo.
func (o *CloudEndpointDataProductInfo) DeepCopyInto(out *CloudEndpointDataProductInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudEndpointDataProductInfo: %s", err))
	}

	*out = *target.(*CloudEndpointDataProductInfo)
}

// Validate valides the current information stored into the structure.
func (o *CloudEndpointDataProductInfo) Validate() error {

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
func (*CloudEndpointDataProductInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudEndpointDataProductInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudEndpointDataProductInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudEndpointDataProductInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudEndpointDataProductInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudEndpointDataProductInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "productID":
		return o.ProductID
	case "type":
		return o.Type
	}

	return nil
}

// CloudEndpointDataProductInfoAttributesMap represents the map of attribute for CloudEndpointDataProductInfo.
var CloudEndpointDataProductInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"ProductID": {
		AllowedChoices: []string{},
		BSONFieldName:  "productid",
		ConvertedName:  "ProductID",
		Description:    `The ID of the corresponding product.`,
		Exposed:        true,
		Name:           "productID",
		Stored:         true,
		Type:           "string",
	},
	"Type": {
		AllowedChoices: []string{},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `The type of the product.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "string",
	},
}

// CloudEndpointDataProductInfoLowerCaseAttributesMap represents the map of attribute for CloudEndpointDataProductInfo.
var CloudEndpointDataProductInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"productid": {
		AllowedChoices: []string{},
		BSONFieldName:  "productid",
		ConvertedName:  "ProductID",
		Description:    `The ID of the corresponding product.`,
		Exposed:        true,
		Name:           "productID",
		Stored:         true,
		Type:           "string",
	},
	"type": {
		AllowedChoices: []string{},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `The type of the product.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudEndpointDataProductInfo struct {
	ProductID string `bson:"productid,omitempty"`
	Type      string `bson:"type,omitempty"`
}
