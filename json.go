package main

import (
	"encoding/json"
	"os"
)

type PackageJson struct {
	Version string `json:"version,omitempty"`
}

func ReadPackageJson(path string) (*PackageJson, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg PackageJson
	err = json.Unmarshal(file, &pkg)
	if err != nil {
		return nil, err
	}

	return &pkg, nil
}
