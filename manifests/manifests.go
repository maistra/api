package manifests

import (
	"embed"
	"io/fs"
)

//go:embed *.yaml
var res embed.FS

// GetManifests returns a slice of the DirEntries embedded in this folder.
func GetManifests() ([]fs.DirEntry, error) {
	return res.ReadDir(".")
}

// ReadManifest reads embedded file content for the passed name.
func ReadManifest(name string) ([]byte, error) {
	return res.ReadFile(name)
}

// GetManifestsByName returns a map of all manifest names and their content.
// It propagates error if it wasn't able to read the content.
func GetManifestsByName() (map[string][]byte, error) {
	entries, err := res.ReadDir(".")
	if err != nil {
		return nil, err
	}

	manifestsByName := make(map[string][]byte)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		content, err := ReadManifest(name)
		if err != nil {
			return nil, err
		}
		manifestsByName[name] = content
	}

	return manifestsByName, nil
}
