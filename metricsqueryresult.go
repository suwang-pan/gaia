// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// MetricsQueryResult represents the model of a metricsqueryresult
type MetricsQueryResult struct {
	// The data of the query.
	Data map[string]interface{} `json:"data,omitempty" msgpack:"data,omitempty" bson:"-" mapstructure:"data,omitempty"`

	// The error message from the query attempt.
	Error string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// The type of error that occurred.
	ErrorType string `json:"errorType,omitempty" msgpack:"errorType,omitempty" bson:"-" mapstructure:"errorType,omitempty"`

	// The status of the query.
	Status string `json:"status" msgpack:"status" bson:"-" mapstructure:"status,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewMetricsQueryResult returns a new *MetricsQueryResult
func NewMetricsQueryResult() *MetricsQueryResult {

	return &MetricsQueryResult{
		ModelVersion: 1,
		Data:         map[string]interface{}{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MetricsQueryResult) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesMetricsQueryResult{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MetricsQueryResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesMetricsQueryResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *MetricsQueryResult) BleveType() string {

	return "metricsqueryresult"
}

// DeepCopy returns a deep copy if the MetricsQueryResult.
func (o *MetricsQueryResult) DeepCopy() *MetricsQueryResult {

	if o == nil {
		return nil
	}

	out := &MetricsQueryResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *MetricsQueryResult.
func (o *MetricsQueryResult) DeepCopyInto(out *MetricsQueryResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy MetricsQueryResult: %s", err))
	}

	*out = *target.(*MetricsQueryResult)
}

// Validate valides the current information stored into the structure.
func (o *MetricsQueryResult) Validate() error {

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
func (*MetricsQueryResult) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := MetricsQueryResultAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return MetricsQueryResultLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*MetricsQueryResult) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MetricsQueryResultAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *MetricsQueryResult) ValueForAttribute(name string) any {

	switch name {
	case "data":
		return o.Data
	case "error":
		return o.Error
	case "errorType":
		return o.ErrorType
	case "status":
		return o.Status
	}

	return nil
}

// MetricsQueryResultAttributesMap represents the map of attribute for MetricsQueryResult.
var MetricsQueryResultAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `The data of the query.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `The error message from the query attempt.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"ErrorType": {
		AllowedChoices: []string{},
		ConvertedName:  "ErrorType",
		Description:    `The type of error that occurred.`,
		Exposed:        true,
		Name:           "errorType",
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `The status of the query.`,
		Exposed:        true,
		Name:           "status",
		Type:           "string",
	},
}

// MetricsQueryResultLowerCaseAttributesMap represents the map of attribute for MetricsQueryResult.
var MetricsQueryResultLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `The data of the query.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `The error message from the query attempt.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"errortype": {
		AllowedChoices: []string{},
		ConvertedName:  "ErrorType",
		Description:    `The type of error that occurred.`,
		Exposed:        true,
		Name:           "errorType",
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `The status of the query.`,
		Exposed:        true,
		Name:           "status",
		Type:           "string",
	},
}

type mongoAttributesMetricsQueryResult struct {
}
