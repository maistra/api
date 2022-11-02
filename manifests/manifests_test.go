package manifests_test

import (
	"bytes"
	"fmt"
	"maistra.io/api/manifests"
	"testing"
)

func TestGettingManifestDirEntries(t *testing.T) {
	manifests, err := manifests.GetManifests()
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	if len(manifests) == 0 {
		t.Fatalf("Expected to have non-zero amount of nested files")
	}
}

func TestGettingManifestsByName(t *testing.T) {
	manifestsByName, err := manifests.GetManifestsByName()
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	if len(manifestsByName) == 0 {
		t.Fatalf("Expected to have non-zero amount of nested files")
	}
}

func TestLoadingManifestByName(t *testing.T) {
	manifestsByName, err := manifests.GetManifestsByName()
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	if len(manifestsByName) == 0 {
		t.Fatalf("Expected to have non-zero amount of nested files")
	}
	smmYAML := "maistra.io_servicemeshmembers.yaml"

	smmContent, err := manifests.ReadManifest(smmYAML)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	if bytes.Compare(manifestsByName[smmYAML], smmContent) != 0 {
		t.Fatalf("Expected to have identical content")
	}
}
