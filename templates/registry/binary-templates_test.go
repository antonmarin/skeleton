package registry

import (
	"io/ioutil"
	"testing"
)

func TestBinaryTemplatesRegistry_GetTemplateString_ShouldGetTestFileTemplateString(t *testing.T) {
	expectedTemplateString, err := ioutil.ReadFile("../files/testfile")
	if err != nil {
		t.Errorf("File for testing lost or not readable: %s", err)
	}

	registry := BinaryTemplatesRegistry{}
	templateString, err := registry.GetTemplateString("templates/files/testfile")
	if err != nil {
		t.Errorf("Error getting template string: %s", err)
	}

	if templateString != string(expectedTemplateString) {
		t.Errorf("template string should match from binary and file")
	}
}
