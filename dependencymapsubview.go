package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

const (
	// DependencyMapSubviewAttributeNameID represents the attribute ID.
	DependencyMapSubviewAttributeNameID elemental.AttributeSpecificationNameKey = "ID"

	// DependencyMapSubviewAttributeNameSelector represents the attribute selector.
	DependencyMapSubviewAttributeNameSelector elemental.AttributeSpecificationNameKey = "Selector"

	// DependencyMapSubviewAttributeNameSubSelectors represents the attribute subSelectors.
	DependencyMapSubviewAttributeNameSubSelectors elemental.AttributeSpecificationNameKey = "SubSelectors"

	// DependencyMapSubviewAttributeNameTonality represents the attribute tonality.
	DependencyMapSubviewAttributeNameTonality elemental.AttributeSpecificationNameKey = "Tonality"
)

// DependencyMapSubviewIdentity represents the Identity of the object
var DependencyMapSubviewIdentity = elemental.Identity{
	Name:     "dependencymapsubview",
	Category: "dependencymapsubviews",
}

// DependencyMapSubviewsList represents a list of DependencyMapSubviews
type DependencyMapSubviewsList []*DependencyMapSubview

// DependencyMapSubview represents the model of a dependencymapsubview
type DependencyMapSubview struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"-"`

	// Selector is the main selector for the DependencyMapSubview.
	Selector []string `json:"selector,omitempty" cql:"selector,omitempty"`

	// SubSelectors are the selector to apply inside the main selector.
	SubSelectors map[string][]string `json:"subSelectors,omitempty" cql:"subselectors,omitempty"`

	// Tonality sets the color tonality to use for the DependencyMapSubView.
	Tonality string `json:"tonality,omitempty" cql:"tonality,omitempty"`
}

// NewDependencyMapSubview returns a new *DependencyMapSubview
func NewDependencyMapSubview() *DependencyMapSubview {

	return &DependencyMapSubview{
		Selector:     []string{},
		SubSelectors: map[string][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMapSubview) Identity() elemental.Identity {

	return DependencyMapSubviewIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMapSubview) Identifier() string {

	return o.ID
}

func (o *DependencyMapSubview) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMapSubview) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *DependencyMapSubview) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o DependencyMapSubview) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return DependencyMapSubviewAttributesMap[name]
}

// DependencyMapSubviewAttributesMap represents the map of attribute for DependencyMapSubview.
var DependencyMapSubviewAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	DependencyMapSubviewAttributeNameID: elemental.AttributeSpecification{
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
	DependencyMapSubviewAttributeNameSelector: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "selector",
		Orderable:      true,
		Stored:         true,
		SubType:        "dependencymapview_selector",
		Type:           "external",
	},
	DependencyMapSubviewAttributeNameSubSelectors: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "subSelectors",
		Orderable:      true,
		Stored:         true,
		SubType:        "dependencymapview_subselector",
		Type:           "external",
	},
	DependencyMapSubviewAttributeNameTonality: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "tonality",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
