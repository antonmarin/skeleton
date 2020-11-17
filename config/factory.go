package config

import (
	"fmt"
	"github.com/antonmarin/skeleton/config/params/osFamily"
	"os"
	"runtime"
)

//Factory creates Config
type Factory struct {
}

//NewFactory is a constructor of Factory
func NewFactory() *Factory {
	return &Factory{}
}

//ConstructConfig creates Config
func (factory Factory) ConstructConfig() (Config, error) {
	osStringValue := os.Getenv("GOOS")
	if len(osStringValue) == 0 {
		osStringValue = runtime.GOOS
	}
	osTypeValue := osFamily.NewFromGOOS(osStringValue)
	if osTypeValue == osFamily.Undefined {
		return Config{}, fmt.Errorf("undefined OS from env: GOOS=%s", osStringValue)
	}
	workingDirectory, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("error defining working directory: %s", err)
	}

	return Config{
		osFamily:         osTypeValue,
		WorkingDirectory: workingDirectory,
	}, nil
}
