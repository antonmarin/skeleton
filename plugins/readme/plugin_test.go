package readme

import (
	"fmt"
	"github.com/antonmarin/skeleton/config"
	"github.com/spf13/afero"
	"os"
	"testing"
)

func TestAcceptShouldCreateReadmeInWorkdir(t *testing.T) {
	filesystem := afero.NewMemMapFs()
	plugin := NewPlugin(filesystem)
	configObj := config.Config{
		WorkingDirectory: t.TempDir(),
	}

	err := plugin.Accept(configObj)
	if err != nil {
		t.Errorf("error running test: %s", err)
	}

	filename := fmt.Sprintf("%s/README.md", configObj.WorkingDirectory)
	_, err = filesystem.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", filename)
	}
}
