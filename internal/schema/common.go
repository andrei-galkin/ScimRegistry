package schema

import "time"

// Meta holds the resource metadata (RFC 7643 Section 3.1)
type Meta struct {
	ResourceType string    `json:"resourceType"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	Location     string    `json:"location"`
	Version      string    `json:"version,omitempty"` // ETag
}

// MultiValued is a generic struct for attributes like emails, phoneNumbers, etc.
type MultiValued struct {
	Value   string `json:"value"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Display string `json:"display,omitempty"`
}
