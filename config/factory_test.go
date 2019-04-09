package config

import "testing"

func TestFactory_CreateConfig(t *testing.T) {
	factory := NewFactory()

	_, err := factory.CreateConfig()
	if err != nil {
		t.Errorf("Factory should create config without error. Error: %s", err)
	}
}
