package schema

import "errors"

const UserSchema = "urn:ietf:params:scim:schemas:core:2.0:User"
const EnterpriseSchema = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"

type User struct {
	Schemas     []string      `json:"schemas"`
	ID          string        `json:"id"`
	ExternalID  string        `json:"externalId,omitempty"`
	UserName    string        `json:"userName"` // Required
	DisplayName string        `json:"displayName,omitempty"`
	Active      bool          `json:"active"`
	Emails      []MultiValued `json:"emails,omitempty"`

	// Enterprise Extension Handling
	EnterpriseUser *EnterpriseExtension `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`

	Meta Meta `json:"meta"`
}

type EnterpriseExtension struct {
	EmployeeNumber string `json:"employeeNumber,omitempty"`
	CostCenter     string `json:"costCenter,omitempty"`
	Department     string `json:"department,omitempty"`
	Division       string `json:"division,omitempty"`
}

// Validate handles the Validation Logic for User creation (RFC 7643)
func (u *User) Validate() error {
	if u.UserName == "" {
		return errors.New("attribute 'userName' is required")
	}

	foundCore := false
	for _, s := range u.Schemas {
		if s == UserSchema {
			foundCore = true
			break
		}
	}
	if !foundCore {
		return errors.New("missing core user schema urn")
	}
	return nil
}
