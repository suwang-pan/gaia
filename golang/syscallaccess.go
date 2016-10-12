package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// SyscallAccessIdentity represents the Identity of the object
var SyscallAccessIdentity = elemental.Identity{
	Name:     "syscallaccess",
	Category: "syscallaccesses",
}

// SyscallAccessList represents a list of SyscallAccess
type SyscallAccessList []*SyscallAccess

// SyscallAccess represents the model of a syscallaccess
type SyscallAccess struct {
	// PID is the PID of the process that used the system call.
	PID int `json:"PID" cql:"-" bson:"-"`

	// Actions tells if the system call has been allowed.
	Action string `json:"action" cql:"-" bson:"-"`

	// Count tells how many times the syscall has been sent.
	Count int `json:"count" cql:"-" bson:"-"`

	// Name represents the name of the system call.
	Name string `json:"name" cql:"-" bson:"-"`

	// ProcessName is the name of the process that used the system call.
	ProcessName string `json:"processName" cql:"-" bson:"-"`
}

// NewSyscallAccess returns a new *SyscallAccess
func NewSyscallAccess() *SyscallAccess {

	return &SyscallAccess{}
}

// Identity returns the Identity of the object.
func (o *SyscallAccess) Identity() elemental.Identity {

	return SyscallAccessIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SyscallAccess) Identifier() string {

	return ""
}

func (o *SyscallAccess) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SyscallAccess) SetIdentifier(ID string) {

}

// Validate valides the current information stored into the structure.
func (o *SyscallAccess) Validate() error {

	errors := elemental.Errors{}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o SyscallAccess) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return SyscallAccessAttributesMap[name]
}

// SyscallAccessAttributesMap represents the map of attribute for SyscallAccess.
var SyscallAccessAttributesMap = map[string]elemental.AttributeSpecification{
	"PID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "PID",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "action",
		ReadOnly:       true,
		Type:           "string",
	},
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"ProcessName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "processName",
		ReadOnly:       true,
		Type:           "string",
	},
}
