package gql

import (
	"context"
	"errors"
	"testing"

	"github.com/shpota/skmz/model"
)

// emptyDB always returns an empty slice and no error.
type emptyDB struct{}

func (emptyDB) GetProgrammers(string) ([]*model.Programmer, error) {
	return []*model.Programmer{}, nil
}

// errDB always returns nil slice and an error.
type errDB struct{}

func (errDB) GetProgrammers(string) ([]*model.Programmer, error) {
	return nil, errors.New("connection refused")
}

func TestProgrammers_EmptyResult(t *testing.T) {
	r := &queryResolver{
		Resolver: &Resolver{DB: emptyDB{}},
	}

	programmers, err := r.Programmers(context.TODO(), "unknown")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(programmers) != 0 {
		t.Errorf("len(programmers) = %d, want 0", len(programmers))
	}
}

func TestProgrammers_DBError(t *testing.T) {
	r := &queryResolver{
		Resolver: &Resolver{DB: errDB{}},
	}

	programmers, err := r.Programmers(context.TODO(), "go")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "connection refused" {
		t.Errorf("error = %q, want %q", err.Error(), "connection refused")
	}
	if programmers != nil {
		t.Errorf("programmers should be nil on error, got %v", programmers)
	}
}

func TestProgrammers_MultipleProgrammers(t *testing.T) {
	icon := "go"
	multiDB := &multiMockDB{}
	r := &queryResolver{
		Resolver: &Resolver{DB: multiDB},
	}

	programmers, err := r.Programmers(context.TODO(), "go")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(programmers) != 2 {
		t.Fatalf("len(programmers) = %d, want 2", len(programmers))
	}
	if programmers[0].Name != "Alice" {
		t.Errorf("programmers[0].Name = %q, want %q", programmers[0].Name, "Alice")
	}
	if programmers[1].Skills[0].Icon != &icon {
		// Just verify icon is set (pointer comparison would fail, check value)
		if *programmers[1].Skills[0].Icon != "go" {
			t.Errorf("skill icon = %q, want %q", *programmers[1].Skills[0].Icon, "go")
		}
	}
}

type multiMockDB struct{}

func (multiMockDB) GetProgrammers(string) ([]*model.Programmer, error) {
	icon := "go"
	return []*model.Programmer{
		{ID: "1", Name: "Alice", Title: "Senior", Company: "ACME"},
		{
			ID: "2", Name: "Bob", Title: "Junior", Company: "Corp",
			Skills: []*model.Skill{{ID: "s1", Name: "Go", Icon: &icon, Importance: 3}},
		},
	}, nil
}
