package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ConnectionsResultIdentity represents the Identity of the object.
var ConnectionsResultIdentity = elemental.Identity{
	Name:     "connectionsresult",
	Category: "connectionsresults",
	Package:  "guy",
	Private:  false,
}

// ConnectionsResultsList represents a list of ConnectionsResults
type ConnectionsResultsList []*ConnectionsResult

// Identity returns the identity of the objects in the list.
func (o ConnectionsResultsList) Identity() elemental.Identity {

	return ConnectionsResultIdentity
}

// Copy returns a pointer to a copy the ConnectionsResultsList.
func (o ConnectionsResultsList) Copy() elemental.Identifiables {

	copy := append(ConnectionsResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ConnectionsResultsList.
func (o ConnectionsResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ConnectionsResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ConnectionsResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ConnectionsResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ConnectionsResultsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ConnectionsResultsList converted to SparseConnectionsResultsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ConnectionsResultsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseConnectionsResultsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseConnectionsResult)
	}

	return out
}

// Version returns the version of the content.
func (o ConnectionsResultsList) Version() int {

	return 1
}

// ConnectionsResult represents the model of a connectionsresult
type ConnectionsResult struct {
	// Contains a batch of connections.
	Connections []*CurrentConnection `json:"connections" msgpack:"connections" bson:"-" mapstructure:"connections,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Unique ID generated for each request.
	RequestID string `json:"requestID" msgpack:"requestID" bson:"-" mapstructure:"requestID,omitempty"`

	// Contains the total number of connections for the connection request.
	TotalConnections int `json:"totalConnections,omitempty" msgpack:"totalConnections,omitempty" bson:"-" mapstructure:"totalConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewConnectionsResult returns a new *ConnectionsResult
func NewConnectionsResult() *ConnectionsResult {

	return &ConnectionsResult{
		ModelVersion: 1,
		Connections:  []*CurrentConnection{},
	}
}

// Identity returns the Identity of the object.
func (o *ConnectionsResult) Identity() elemental.Identity {

	return ConnectionsResultIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ConnectionsResult) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ConnectionsResult) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionsResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesConnectionsResult{}

	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionsResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesConnectionsResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ConnectionsResult) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ConnectionsResult) BleveType() string {

	return "connectionsresult"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ConnectionsResult) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ConnectionsResult) Doc() string {

	return `A result from a connections request.`
}

func (o *ConnectionsResult) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *ConnectionsResult) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ConnectionsResult) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ConnectionsResult) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseConnectionsResult{
			Connections:      &o.Connections,
			Namespace:        &o.Namespace,
			RequestID:        &o.RequestID,
			TotalConnections: &o.TotalConnections,
		}
	}

	sp := &SparseConnectionsResult{}
	for _, f := range fields {
		switch f {
		case "connections":
			sp.Connections = &(o.Connections)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "requestID":
			sp.RequestID = &(o.RequestID)
		case "totalConnections":
			sp.TotalConnections = &(o.TotalConnections)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseConnectionsResult to the object.
func (o *ConnectionsResult) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseConnectionsResult)
	if so.Connections != nil {
		o.Connections = *so.Connections
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.RequestID != nil {
		o.RequestID = *so.RequestID
	}
	if so.TotalConnections != nil {
		o.TotalConnections = *so.TotalConnections
	}
}

// DeepCopy returns a deep copy if the ConnectionsResult.
func (o *ConnectionsResult) DeepCopy() *ConnectionsResult {

	if o == nil {
		return nil
	}

	out := &ConnectionsResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ConnectionsResult.
func (o *ConnectionsResult) DeepCopyInto(out *ConnectionsResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ConnectionsResult: %s", err))
	}

	*out = *target.(*ConnectionsResult)
}

// Validate valides the current information stored into the structure.
func (o *ConnectionsResult) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Connections {
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
func (*ConnectionsResult) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ConnectionsResultAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ConnectionsResultLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ConnectionsResult) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ConnectionsResultAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ConnectionsResult) ValueForAttribute(name string) interface{} {

	switch name {
	case "connections":
		return o.Connections
	case "namespace":
		return o.Namespace
	case "requestID":
		return o.RequestID
	case "totalConnections":
		return o.TotalConnections
	}

	return nil
}

// ConnectionsResultAttributesMap represents the map of attribute for ConnectionsResult.
var ConnectionsResultAttributesMap = map[string]elemental.AttributeSpecification{
	"Connections": {
		AllowedChoices: []string{},
		ConvertedName:  "Connections",
		Description:    `Contains a batch of connections.`,
		Exposed:        true,
		Name:           "connections",
		SubType:        "currentconnection",
		Type:           "refList",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"RequestID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequestID",
		Description:    `Unique ID generated for each request.`,
		Exposed:        true,
		Name:           "requestID",
		ReadOnly:       true,
		Type:           "string",
	},
	"TotalConnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TotalConnections",
		Description:    `Contains the total number of connections for the connection request.`,
		Exposed:        true,
		Name:           "totalConnections",
		Type:           "integer",
	},
}

// ConnectionsResultLowerCaseAttributesMap represents the map of attribute for ConnectionsResult.
var ConnectionsResultLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"connections": {
		AllowedChoices: []string{},
		ConvertedName:  "Connections",
		Description:    `Contains a batch of connections.`,
		Exposed:        true,
		Name:           "connections",
		SubType:        "currentconnection",
		Type:           "refList",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"requestid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequestID",
		Description:    `Unique ID generated for each request.`,
		Exposed:        true,
		Name:           "requestID",
		ReadOnly:       true,
		Type:           "string",
	},
	"totalconnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TotalConnections",
		Description:    `Contains the total number of connections for the connection request.`,
		Exposed:        true,
		Name:           "totalConnections",
		Type:           "integer",
	},
}

// SparseConnectionsResultsList represents a list of SparseConnectionsResults
type SparseConnectionsResultsList []*SparseConnectionsResult

// Identity returns the identity of the objects in the list.
func (o SparseConnectionsResultsList) Identity() elemental.Identity {

	return ConnectionsResultIdentity
}

// Copy returns a pointer to a copy the SparseConnectionsResultsList.
func (o SparseConnectionsResultsList) Copy() elemental.Identifiables {

	copy := append(SparseConnectionsResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseConnectionsResultsList.
func (o SparseConnectionsResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseConnectionsResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseConnectionsResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseConnectionsResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseConnectionsResultsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseConnectionsResultsList converted to ConnectionsResultsList.
func (o SparseConnectionsResultsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseConnectionsResultsList) Version() int {

	return 1
}

// SparseConnectionsResult represents the sparse version of a connectionsresult.
type SparseConnectionsResult struct {
	// Contains a batch of connections.
	Connections *[]*CurrentConnection `json:"connections,omitempty" msgpack:"connections,omitempty" bson:"-" mapstructure:"connections,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Unique ID generated for each request.
	RequestID *string `json:"requestID,omitempty" msgpack:"requestID,omitempty" bson:"-" mapstructure:"requestID,omitempty"`

	// Contains the total number of connections for the connection request.
	TotalConnections *int `json:"totalConnections,omitempty" msgpack:"totalConnections,omitempty" bson:"-" mapstructure:"totalConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseConnectionsResult returns a new  SparseConnectionsResult.
func NewSparseConnectionsResult() *SparseConnectionsResult {
	return &SparseConnectionsResult{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseConnectionsResult) Identity() elemental.Identity {

	return ConnectionsResultIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseConnectionsResult) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseConnectionsResult) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionsResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseConnectionsResult{}

	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionsResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseConnectionsResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseConnectionsResult) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseConnectionsResult) ToPlain() elemental.PlainIdentifiable {

	out := NewConnectionsResult()
	if o.Connections != nil {
		out.Connections = *o.Connections
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.RequestID != nil {
		out.RequestID = *o.RequestID
	}
	if o.TotalConnections != nil {
		out.TotalConnections = *o.TotalConnections
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseConnectionsResult) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseConnectionsResult) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseConnectionsResult.
func (o *SparseConnectionsResult) DeepCopy() *SparseConnectionsResult {

	if o == nil {
		return nil
	}

	out := &SparseConnectionsResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseConnectionsResult.
func (o *SparseConnectionsResult) DeepCopyInto(out *SparseConnectionsResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseConnectionsResult: %s", err))
	}

	*out = *target.(*SparseConnectionsResult)
}

type mongoAttributesConnectionsResult struct {
	Namespace string `bson:"namespace"`
}
type mongoAttributesSparseConnectionsResult struct {
	Namespace *string `bson:"namespace,omitempty"`
}