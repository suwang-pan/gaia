package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

const (
	// HealthReportAttributeNameAPIVersion represents the attribute APIVersion.
	HealthReportAttributeNameAPIVersion elemental.AttributeSpecificationNameKey = "APIVersion"

	// HealthReportAttributeNameID represents the attribute ID.
	HealthReportAttributeNameID elemental.AttributeSpecificationNameKey = "ID"

	// HealthReportAttributeNameBahamutVersion represents the attribute bahamutVersion.
	HealthReportAttributeNameBahamutVersion elemental.AttributeSpecificationNameKey = "BahamutVersion"

	// HealthReportAttributeNameElementalVersion represents the attribute elementalVersion.
	HealthReportAttributeNameElementalVersion elemental.AttributeSpecificationNameKey = "ElementalVersion"

	// HealthReportAttributeNameGaiaVersion represents the attribute gaiaVersion.
	HealthReportAttributeNameGaiaVersion elemental.AttributeSpecificationNameKey = "GaiaVersion"

	// HealthReportAttributeNameSquallVersion represents the attribute squallVersion.
	HealthReportAttributeNameSquallVersion elemental.AttributeSpecificationNameKey = "SquallVersion"

	// HealthReportAttributeNameStatus represents the attribute status.
	HealthReportAttributeNameStatus elemental.AttributeSpecificationNameKey = "Status"
)

// HealthReportStatusValue represents the possible values for attribute "status".
type HealthReportStatusValue string

const (
	// HealthReportStatusDegraded represents the value Degraded.
	HealthReportStatusDegraded HealthReportStatusValue = "Degraded"

	// HealthReportStatusFailure represents the value Failure.
	HealthReportStatusFailure HealthReportStatusValue = "Failure"

	// HealthReportStatusOk represents the value Ok.
	HealthReportStatusOk HealthReportStatusValue = "Ok"
)

// HealthReportIdentity represents the Identity of the object
var HealthReportIdentity = elemental.Identity{
	Name:     "healthreport",
	Category: "healthreports",
}

// HealthReportsList represents a list of HealthReports
type HealthReportsList []*HealthReport

// HealthReport represents the model of a healthreport
type HealthReport struct {
	// APIVersion is the API version served by the server.
	APIVersion string `json:"APIVersion,omitempty" cql:"-"`

	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"-"`

	// bahamutVersion is the version of Bahamut used by the server.
	BahamutVersion string `json:"bahamutVersion,omitempty" cql:"-"`

	// elementalVersion is the version of Elemental used by the server.
	ElementalVersion string `json:"elementalVersion,omitempty" cql:"-"`

	// gaiaVersion is the version of Gaia used by the server.
	GaiaVersion string `json:"gaiaVersion,omitempty" cql:"-"`

	// SquallVersion is the version of server.
	SquallVersion string `json:"squallVersion,omitempty" cql:"-"`

	// Status is the overall health status of the server.
	Status HealthReportStatusValue `json:"status,omitempty" cql:"-"`
}

// NewHealthReport returns a new *HealthReport
func NewHealthReport() *HealthReport {

	return &HealthReport{}
}

// Identity returns the Identity of the object.
func (o *HealthReport) Identity() elemental.Identity {

	return HealthReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HealthReport) Identifier() string {

	return o.ID
}

func (o *HealthReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HealthReport) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *HealthReport) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Degraded", "Failure", "Ok"}); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o HealthReport) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return HealthReportAttributesMap[name]
}

// HealthReportAttributesMap represents the map of attribute for HealthReport.
var HealthReportAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	HealthReportAttributeNameAPIVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "APIVersion",
		Orderable:      true,
		Type:           "string",
	},
	HealthReportAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Type:           "string",
		Unique:         true,
	},
	HealthReportAttributeNameBahamutVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "bahamutVersion",
		Orderable:      true,
		Type:           "string",
	},
	HealthReportAttributeNameElementalVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "elementalVersion",
		Orderable:      true,
		Type:           "string",
	},
	HealthReportAttributeNameGaiaVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "gaiaVersion",
		Orderable:      true,
		Type:           "string",
	},
	HealthReportAttributeNameSquallVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "squallVersion",
		Orderable:      true,
		Type:           "string",
	},
	HealthReportAttributeNameStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"Degraded", "Failure", "Ok"},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Type:           "enum",
	},
}
