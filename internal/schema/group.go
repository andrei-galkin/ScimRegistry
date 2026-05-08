package schema

import "errors"

// Group defines the core SCIM Group attributes (RFC 7643 Section 4.2)
type Group struct {
	Schemas     []string `json:"schemas"`
	ID          string   `json:"id"`
	ExternalID  string   `json:"externalId,omitempty"`
	DisplayName string   `json:"displayName"`
	Members     []Member `json:"members,omitempty"`
	Meta        Meta     `json:"meta"`
}

// Member represents a user or another group inside a group
type Member struct {
	Value   string `json:"value"`   // The User or Group ID
	Ref     string `json:"$ref"`    // URL to the Resource
	Display string `json:"display"` // Name for easy reading
	Type    string `json:"type"`    // "User" or "Group"
}

// Validate logic for Group
func (g *Group) Validate() error {
	if g.DisplayName == "" {
		return errors.New("attribute 'displayName' is required")
	}
	return nil
}
