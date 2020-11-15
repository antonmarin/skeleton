package config

import (
	"github.com/antonmarin/skeleton/config/params/osFamily"
	"os"
	"testing"
)

func TestFactory_ShouldCreateConfigWhenOsTypePassedAsEnvVariable(t *testing.T) {
	var config Config
	factory := NewFactory()

	_ = os.Setenv("GOOS", "darwin")
	config, _ = factory.ConstructConfig()
	if config.osFamily != osFamily.MacOs {
		t.Errorf("Should set macOS OsFamily if GOOS=darwin. Got: %s", config.osFamily)
	}

	_ = os.Setenv("GOOS", "unknown_os")
	_, err := factory.ConstructConfig()
	if err == nil {
		t.Errorf("Should return error if GOOS=unknown_os. Got: %s", config.osFamily)
	}
}
