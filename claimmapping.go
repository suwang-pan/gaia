package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ClaimMapping represents the model of a claimmapping
type ClaimMapping struct {
	// The name of the claim to map to the HTTP header. header.
	ClaimName string `json:"claimName" msgpack:"claimName" bson:"claimname" mapstructure:"claimName,omitempty"`

	// The HTTP header that will be the destination of the mapped claim.
	TargetHTTPHeader string `json:"targetHTTPHeader" msgpack:"targetHTTPHeader" bson:"targethttpheader" mapstructure:"targetHTTPHeader,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewClaimMapping returns a new *ClaimMapping
func NewClaimMapping() *ClaimMapping {

	return &ClaimMapping{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ClaimMapping) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesClaimMapping{}

	s.ClaimName = o.ClaimName
	s.TargetHTTPHeader = o.TargetHTTPHeader

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ClaimMapping) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesClaimMapping{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ClaimName = s.ClaimName
	o.TargetHTTPHeader = s.TargetHTTPHeader

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *ClaimMapping) BleveType() string {

	return "claimmapping"
}

// DeepCopy returns a deep copy if the ClaimMapping.
func (o *ClaimMapping) DeepCopy() *ClaimMapping {

	if o == nil {
		return nil
	}

	out := &ClaimMapping{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ClaimMapping.
func (o *ClaimMapping) DeepCopyInto(out *ClaimMapping) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ClaimMapping: %s", err))
	}

	*out = *target.(*ClaimMapping)
}

// Validate valides the current information stored into the structure.
func (o *ClaimMapping) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("claimName", o.ClaimName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidatePattern("claimName", o.ClaimName, `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`, `must be an alpha numerical character or '-', '_', '/', '*', '#', '&', '@', '_', '$' ~ or ':'`, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("targetHTTPHeader", o.TargetHTTPHeader); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidatePattern("targetHTTPHeader", o.TargetHTTPHeader, `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`, `must be an alpha numerical character or '-', '_', '/', '*', '#', '&', '@', '_', '$' ~ or ':'`, true); err != nil {
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
func (*ClaimMapping) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ClaimMappingAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ClaimMappingLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ClaimMapping) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ClaimMappingAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ClaimMapping) ValueForAttribute(name string) interface{} {

	switch name {
	case "claimName":
		return o.ClaimName
	case "targetHTTPHeader":
		return o.TargetHTTPHeader
	}

	return nil
}

// ClaimMappingAttributesMap represents the map of attribute for ClaimMapping.
var ClaimMappingAttributesMap = map[string]elemental.AttributeSpecification{
	"ClaimName": {
		AllowedChars:   `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`,
		AllowedChoices: []string{},
		BSONFieldName:  "claimname",
		ConvertedName:  "ClaimName",
		Description:    `The name of the claim to map to the HTTP header. header.`,
		Exposed:        true,
		Name:           "claimName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"TargetHTTPHeader": {
		AllowedChars:   `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`,
		AllowedChoices: []string{},
		BSONFieldName:  "targethttpheader",
		ConvertedName:  "TargetHTTPHeader",
		Description:    `The HTTP header that will be the destination of the mapped claim.`,
		Exposed:        true,
		Name:           "targetHTTPHeader",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// ClaimMappingLowerCaseAttributesMap represents the map of attribute for ClaimMapping.
var ClaimMappingLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claimname": {
		AllowedChars:   `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`,
		AllowedChoices: []string{},
		BSONFieldName:  "claimname",
		ConvertedName:  "ClaimName",
		Description:    `The name of the claim to map to the HTTP header. header.`,
		Exposed:        true,
		Name:           "claimName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"targethttpheader": {
		AllowedChars:   `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`,
		AllowedChoices: []string{},
		BSONFieldName:  "targethttpheader",
		ConvertedName:  "TargetHTTPHeader",
		Description:    `The HTTP header that will be the destination of the mapped claim.`,
		Exposed:        true,
		Name:           "targetHTTPHeader",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesClaimMapping struct {
	ClaimName        string `bson:"claimname"`
	TargetHTTPHeader string `bson:"targethttpheader"`
}
