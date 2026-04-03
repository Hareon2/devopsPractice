package db

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNew_ReturnsNonNil(t *testing.T) {
	// Create a minimal mongo.Client (not connected — we only verify
	// that New() wires the collection reference without panic).
	client, _ := mongo.NewClient()
	mdb := New(client)
	if mdb == nil {
		t.Fatal("New() returned nil")
	}
	if mdb.collection == nil {
		t.Fatal("New() did not initialise the collection field")
	}
}

func TestFilter_EmptySkill(t *testing.T) {
	mdb := MongoDB{}
	got := mdb.filter("")
	// With an empty skill the regex should be "^.*$" which matches everything.
	if got[0].Key != "skills.name" {
		t.Errorf("filter key = %q, want %q", got[0].Key, "skills.name")
	}
}

func TestFilter_SpecialCharacters(t *testing.T) {
	mdb := MongoDB{}
	got := mdb.filter("C++")
	// Ensures the filter does not panic on regex-unsafe input.
	if got[0].Key != "skills.name" {
		t.Errorf("filter key = %q, want %q", got[0].Key, "skills.name")
	}
}
