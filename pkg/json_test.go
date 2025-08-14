package pkg

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestReadPackageJson(t *testing.T) {
	dir, err := os.MkdirTemp("", "test-output-")
	assert.Nil(t, err)
	defer os.RemoveAll(dir)

	p := PackageJson{
		Version: "v25.9.0",
	}
	b, err := json.Marshal(p)
	assert.Nil(t, err)
	path := filepath.Join(dir, "package.json")
	err = os.WriteFile(path, b, 0777)
	assert.Nil(t, err)

	pkg, err := ReadPackageJson(path)
	assert.Nil(t, err)
	assert.Equal(t, pkg.Version, "v25.9.0")
}
