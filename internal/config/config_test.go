package config

import (
	"path/filepath"
	"testing"
)

func TestResolveDataDirDevelopment(t *testing.T) {
	t.Parallel()

	got, err := ResolveDataDir("", true)
	if err != nil {
		t.Fatalf("ResolveDataDir() error = %v", err)
	}
	want, err := filepath.Abs("data")
	if err != nil {
		t.Fatalf("filepath.Abs() error = %v", err)
	}
	if got != want {
		t.Fatalf("ResolveDataDir() = %q, want %q", got, want)
	}
}

func TestResolveDataDirExplicit(t *testing.T) {
	t.Parallel()

	got, err := ResolveDataDir("./custom-data", false)
	if err != nil {
		t.Fatalf("ResolveDataDir() error = %v", err)
	}
	want, err := filepath.Abs("custom-data")
	if err != nil {
		t.Fatalf("filepath.Abs() error = %v", err)
	}
	if got != want {
		t.Fatalf("ResolveDataDir() = %q, want %q", got, want)
	}
}
