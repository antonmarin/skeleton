package readme

import (
	"fmt"
	"github.com/antonmarin/skeleton/config"
	"github.com/spf13/afero"
	"os"
	"text/template"
)

type Plugin struct {
	filesystem afero.Fs
}

func NewPlugin(filesystem afero.Fs) *Plugin {
	return &Plugin{
		filesystem: filesystem,
	}
}

func (p Plugin) Accept(config config.Config) error {
	t, err := template.New("README.md").ParseFiles("templates/README.md")
	if err != nil {
		return fmt.Errorf("failed parsing template: %w", err)
	}

	filename := fmt.Sprintf("%s/README.md", config.WorkingDirectory)
	_, err = p.filesystem.Stat(filename)
	var readme afero.File
	if os.IsNotExist(err) {
		readme, err = p.filesystem.Create(filename)
	} else {
		readme, err = p.filesystem.Open(filename)
	}
	if err != nil {
		return fmt.Errorf("failed opening target file: %w", err)
	}

	err = t.Execute(readme, "")
	if err != nil {
		return fmt.Errorf("error templating: %w", err)
	}

	return nil
}
