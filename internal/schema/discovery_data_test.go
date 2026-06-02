package schema

import (
	"encoding/json"
	"testing"
)

func TestRawSchemasJSONValid(t *testing.T) {
	if !json.Valid(RawSchemasJSON) {
		t.Fatal("RawSchemasJSON is not valid JSON")
	}
}

func TestSchemasCount(t *testing.T) {
	if len(Schemas) != 2 {
		t.Fatalf("expected 2 schemas, got %d", len(Schemas))
	}
}
