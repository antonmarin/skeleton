package registry

import (
	"fmt"
	"github.com/antonmarin/skeleton/templates"
)

type BinaryTemplatesRegistry struct{}

func (reg BinaryTemplatesRegistry) GetTemplateString(filename string) (string, error) {
	templateBytes, err := templates.Asset(filename)
	if err != nil {
		return "", fmt.Errorf("error getting template: %s", err)
	}

	return string(templateBytes), nil
}
