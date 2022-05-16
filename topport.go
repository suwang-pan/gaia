// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TopPortIdentity represents the Identity of the object.
var TopPortIdentity = elemental.Identity{
	Name:     "topport",
	Category: "topPorts",
	Package:  "agrias",
	Private:  false,
}

// TopPortsList represents a list of TopPorts
type TopPortsList []*TopPort

// Identity returns the identity of the objects in the list.
func (o TopPortsList) Identity() elemental.Identity {

	return TopPortIdentity
}

// Copy returns a pointer to a copy the TopPortsList.
func (o TopPortsList) Copy() elemental.Identifiables {

	copy := append(TopPortsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TopPortsList.
func (o TopPortsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TopPortsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TopPort))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TopPortsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TopPortsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TopPortsList converted to SparseTopPortsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TopPortsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTopPortsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTopPort)
	}

	return out
}

// Version returns the version of the content.
func (o TopPortsList) Version() int {

	return 1
}

// TopPort represents the model of a topport
type TopPort struct {
	// The entry encountered.
	Entry string `json:"entry" msgpack:"entry" bson:"-" mapstructure:"entry,omitempty"`

	// Number of times the entry showed up.
	Occurrences int `json:"occurrences" msgpack:"occurrences" bson:"-" mapstructure:"occurrences,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTopPort returns a new *TopPort
func NewTopPort() *TopPort {

	return &TopPort{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *TopPort) Identity() elemental.Identity {

	return TopPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TopPort) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TopPort) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TopPort) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTopPort{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TopPort) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTopPort{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TopPort) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TopPort) BleveType() string {

	return "topport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TopPort) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TopPort) Doc() string {

	return `Represent a top port within a bucket.`
}

func (o *TopPort) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TopPort) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTopPort{
			Entry:       &o.Entry,
			Occurrences: &o.Occurrences,
		}
	}

	sp := &SparseTopPort{}
	for _, f := range fields {
		switch f {
		case "entry":
			sp.Entry = &(o.Entry)
		case "occurrences":
			sp.Occurrences = &(o.Occurrences)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTopPort to the object.
func (o *TopPort) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTopPort)
	if so.Entry != nil {
		o.Entry = *so.Entry
	}
	if so.Occurrences != nil {
		o.Occurrences = *so.Occurrences
	}
}

// DeepCopy returns a deep copy if the TopPort.
func (o *TopPort) DeepCopy() *TopPort {

	if o == nil {
		return nil
	}

	out := &TopPort{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TopPort.
func (o *TopPort) DeepCopyInto(out *TopPort) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TopPort: %s", err))
	}

	*out = *target.(*TopPort)
}

// Validate valides the current information stored into the structure.
func (o *TopPort) Validate() error {

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
func (*TopPort) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TopPortAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TopPortLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TopPort) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TopPortAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TopPort) ValueForAttribute(name string) interface{} {

	switch name {
	case "entry":
		return o.Entry
	case "occurrences":
		return o.Occurrences
	}

	return nil
}

// TopPortAttributesMap represents the map of attribute for TopPort.
var TopPortAttributesMap = map[string]elemental.AttributeSpecification{
	"Entry": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Entry",
		Description:    `The entry encountered.`,
		Exposed:        true,
		Name:           "entry",
		ReadOnly:       true,
		Type:           "string",
	},
	"Occurrences": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Occurrences",
		Description:    `Number of times the entry showed up.`,
		Exposed:        true,
		Name:           "occurrences",
		ReadOnly:       true,
		Type:           "integer",
	},
}

// TopPortLowerCaseAttributesMap represents the map of attribute for TopPort.
var TopPortLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"entry": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Entry",
		Description:    `The entry encountered.`,
		Exposed:        true,
		Name:           "entry",
		ReadOnly:       true,
		Type:           "string",
	},
	"occurrences": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Occurrences",
		Description:    `Number of times the entry showed up.`,
		Exposed:        true,
		Name:           "occurrences",
		ReadOnly:       true,
		Type:           "integer",
	},
}

// SparseTopPortsList represents a list of SparseTopPorts
type SparseTopPortsList []*SparseTopPort

// Identity returns the identity of the objects in the list.
func (o SparseTopPortsList) Identity() elemental.Identity {

	return TopPortIdentity
}

// Copy returns a pointer to a copy the SparseTopPortsList.
func (o SparseTopPortsList) Copy() elemental.Identifiables {

	copy := append(SparseTopPortsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTopPortsList.
func (o SparseTopPortsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTopPortsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTopPort))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTopPortsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTopPortsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTopPortsList converted to TopPortsList.
func (o SparseTopPortsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTopPortsList) Version() int {

	return 1
}

// SparseTopPort represents the sparse version of a topport.
type SparseTopPort struct {
	// The entry encountered.
	Entry *string `json:"entry,omitempty" msgpack:"entry,omitempty" bson:"-" mapstructure:"entry,omitempty"`

	// Number of times the entry showed up.
	Occurrences *int `json:"occurrences,omitempty" msgpack:"occurrences,omitempty" bson:"-" mapstructure:"occurrences,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTopPort returns a new  SparseTopPort.
func NewSparseTopPort() *SparseTopPort {
	return &SparseTopPort{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTopPort) Identity() elemental.Identity {

	return TopPortIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTopPort) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTopPort) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTopPort) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTopPort{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTopPort) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTopPort{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTopPort) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTopPort) ToPlain() elemental.PlainIdentifiable {

	out := NewTopPort()
	if o.Entry != nil {
		out.Entry = *o.Entry
	}
	if o.Occurrences != nil {
		out.Occurrences = *o.Occurrences
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTopPort.
func (o *SparseTopPort) DeepCopy() *SparseTopPort {

	if o == nil {
		return nil
	}

	out := &SparseTopPort{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTopPort.
func (o *SparseTopPort) DeepCopyInto(out *SparseTopPort) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTopPort: %s", err))
	}

	*out = *target.(*SparseTopPort)
}

type mongoAttributesTopPort struct {
}
type mongoAttributesSparseTopPort struct {
}
