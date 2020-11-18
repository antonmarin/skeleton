package registry

import "errors"

type MemoryTemplatesRegistry struct {
	templates map[string]string
}

func NewMemoryTemplatesRegistry(templates map[string]string) *MemoryTemplatesRegistry {
	return &MemoryTemplatesRegistry{templates: templates}
}

func (reg MemoryTemplatesRegistry) GetTemplateString(filename string) (string, error) {
	if v, found := reg.templates[filename]; found {
		return v, nil
	}
	return "", errors.New("template for requested file not found")
}
