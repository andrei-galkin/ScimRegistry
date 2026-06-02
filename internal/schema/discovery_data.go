package schema

import "encoding/json"

type SchemaAttribute struct {
	Name          string            `json:"name"`
	Type          string            `json:"type"`
	MultiValued   bool              `json:"multiValued"`
	Description   string            `json:"description"`
	Required      bool              `json:"required"`
	CaseExact     bool              `json:"caseExact,omitempty"`
	Mutability    string            `json:"mutability,omitempty"`
	Returned      string            `json:"returned,omitempty"`
	Uniqueness    string            `json:"uniqueness,omitempty"`
	SubAttributes []SchemaAttribute `json:"subAttributes,omitempty"`
}

type Schema struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Attributes  []SchemaAttribute `json:"attributes"`
}

var Schemas = []Schema{
	{
		ID:          "urn:ietf:params:scim:schemas:core:2.0:User",
		Name:        "User",
		Description: "Core User Schema",
		Attributes: []SchemaAttribute{
			{
				Name:        "userName",
				Type:        "string",
				MultiValued: false,
				Description: "Unique identifier for the User, typically a user log in name.",
				Required:    true,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "server",
			},
			{
				Name:        "displayName",
				Type:        "string",
				MultiValued: false,
				Description: "The name of the User, suitable for display to end-users.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			{
				Name:        "active",
				Type:        "boolean",
				MultiValued: false,
				Description: "A Boolean value indicating the User's administrative status.",
				Required:    false,
				Mutability:  "readWrite",
				Returned:    "default",
			},
		},
	},
	{
		ID:          "urn:ietf:params:scim:schemas:core:2.0:Group",
		Name:        "Group",
		Description: "Core Group Schema",
		Attributes: []SchemaAttribute{
			{
				Name:        "displayName",
				Type:        "string",
				MultiValued: false,
				Description: "A human-readable name for the Group.",
				Required:    true,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			{
				Name:        "members",
				Type:        "complex",
				MultiValued: true,
				Description: "A list of members of the Group.",
				Required:    false,
				Mutability:  "readWrite",
				Returned:    "default",
				SubAttributes: []SchemaAttribute{
					{
						Name:        "value",
						Type:        "string",
						MultiValued: false,
						Description: "Identifier of the member.",
						Required:    true,
						Mutability:  "immutable",
					},
					{
						Name:        "type",
						Type:        "string",
						MultiValued: false,
						Description: "A label indicating the type of resource, e.g., 'User' or 'Group'.",
						Required:    false,
						Mutability:  "immutable",
					},
				},
			},
		},
	},
}

var RawSchemasJSON []byte

func init() {
	var err error
	RawSchemasJSON, err = json.Marshal(Schemas)
	if err != nil {
		panic("Failed to marshal RawSchemasJSON: " + err.Error())
	}
}
