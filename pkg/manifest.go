package pkg

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"os"
)

type PackageJson struct {
	Version string `json:"version,omitempty"`
}

type CargoToml struct {
	Package CargoPackage
}

type CargoPackage struct {
	Version string `toml:"version"`
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

func ReadCargoToml(path string) (*CargoToml, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var m CargoToml
	err = toml.Unmarshal(file, &m)
	return &m, err
}
