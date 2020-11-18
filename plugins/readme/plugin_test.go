package readme

import (
	"fmt"
	"github.com/antonmarin/skeleton/config"
	"github.com/antonmarin/skeleton/templates/registry"
	"github.com/spf13/afero"
	"os"
	"testing"
)

func TestAcceptShouldCreateReadmeInWorkdir(t *testing.T) {
	configObj := config.Config{
		WorkingDirectory: t.TempDir(),
	}
	templateFilePath := "templates/files/plugins/readme/README.md"
	filesystem := afero.NewMemMapFs()
	plugin := NewPlugin(filesystem, registry.NewMemoryTemplatesRegistry(map[string]string{templateFilePath: "test"}))

	err := plugin.Accept(configObj)
	if err != nil {
		t.Errorf("error running test: %s", err)
	}

	expectedFilePath := fmt.Sprintf("%s/README.md", configObj.WorkingDirectory)
	_, err = filesystem.Stat(expectedFilePath)
	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", expectedFilePath)
	}
}
