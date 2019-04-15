package config

import (
	"github.com/antonmarin/skeleton/config/enum/osFamily"
	"os"
	"testing"
)

func TestFactory_ShouldCreateConfigWithOsType(t *testing.T) {
	var config Config
	factory := NewFactory()

	_ = os.Setenv("GOOS", "darwin")
	config, _ = factory.CreateConfig()
	if config.osFamily != osFamily.MacOs {
		t.Errorf("Should set macOS OsFamily if GOOS=darwin. Got: %s", config.osFamily)
	}

	_ = os.Setenv("GOOS", "unknown_os")
	_, err := factory.CreateConfig()
	if err == nil {
		t.Errorf("Should return error if GOOS=unknown_os. Got: %s", config.osFamily)
	}
}
