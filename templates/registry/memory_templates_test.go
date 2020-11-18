package registry

import "testing"

func TestMemoryTemplatesRegistry_GetTemplateStringShouldReturnSetTemplatesWhenRequested(t *testing.T) {
	expectedTemplateString := "Some template"
	templateFilename := "templates/files/readme/README.md"
	registry := NewMemoryTemplatesRegistry(map[string]string{templateFilename: expectedTemplateString})

	templateString, err := registry.GetTemplateString(templateFilename)
	if err != nil {
		t.Errorf("Error running test")
	}

	if templateString != expectedTemplateString {
		t.Errorf("Registry returned wrong template")
	}
}
