package manifests

import (
	"bytes"
	"testing"
)

func TestGettingManifestDirEntries(t *testing.T) {
	manifests, err := GetManifests()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(manifests) == 0 {
		t.Fatalf("Expected to have non-zero amount of nested files")
	}
}

func TestGettingManifestsByName(t *testing.T) {
	manifestsByName, err := GetManifestsByName()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(manifestsByName) == 0 {
		t.Fatalf("Expected to have non-zero amount of nested files")
	}
}

func TestLoadingManifestByName(t *testing.T) {
	manifestsByName, err := GetManifestsByName()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(manifestsByName) == 0 {
		t.Fatalf("Expected to have non-zero amount of nested files")
	}
	smmYAML := "maistra.io_servicemeshmembers.yaml"

	smmContent, err := ReadManifest(smmYAML)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bytes.Compare(manifestsByName[smmYAML], smmContent) != 0 {
		t.Fatalf("Expected to have identical content")
	}
}
