package readme

import (
	"fmt"
	"github.com/antonmarin/skeleton/config"
	"github.com/spf13/afero"
	"os"
	"text/template"
)

type Plugin struct {
	filesystem        afero.Fs
	templatesRegistry TemplatesRegistry
}

func NewPlugin(filesystem afero.Fs, templatesRegistry TemplatesRegistry) *Plugin {
	return &Plugin{filesystem: filesystem, templatesRegistry: templatesRegistry}
}

func (p Plugin) Accept(config config.Config) error {
	templateString, err := p.templatesRegistry.GetTemplateString("templates/files/plugins/readme/README.md")
	if err != nil {
		return fmt.Errorf("failed reading template: %w", err)
	}

	t, err := template.New("README.md").Parse(templateString)
	if err != nil {
		return fmt.Errorf("failed parsing template: %w", err)
	}

	filename := fmt.Sprintf("%s/README.md", config.WorkingDirectory)
	var readme afero.File
	readme, err = p.filesystem.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed opening target file: %w", err)
	}

	err = t.Execute(readme, "")
	if err != nil {
		return fmt.Errorf("error templating: %w", err)
	}

	return nil
}

type TemplatesRegistry interface {
	GetTemplateString(filename string) (string, error)
}
