package main

import (
	"os"
	"testing"
)

func TestClientOptions_UnsetEnv(t *testing.T) {
	os.Unsetenv("profile")

	got := clientOptions()
	want := "mongodb://localhost:27017"

	if got.GetURI() != want {
		t.Errorf("clientOptions() URI = %q, want %q (default to localhost when profile is not set)", got.GetURI(), want)
	}
}

func TestClientOptions_EmptyProfile(t *testing.T) {
	os.Setenv("profile", "")

	got := clientOptions()
	want := "mongodb://localhost:27017"

	if got.GetURI() != want {
		t.Errorf("clientOptions() URI = %q, want %q", got.GetURI(), want)
	}
}

func TestClientOptions_AnyNonProdProfile(t *testing.T) {
	profiles := []string{"dev", "staging", "test", "local"}
	for _, p := range profiles {
		os.Setenv("profile", p)
		got := clientOptions()
		want := "mongodb://localhost:27017"
		if got.GetURI() != want {
			t.Errorf("clientOptions() with profile=%q: URI = %q, want %q", p, got.GetURI(), want)
		}
	}
}
