package config

import (
	"fmt"
	"github.com/antonmarin/skeleton/config/params/osFamily"
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
	goosEnvValue := runtime.GOOS
	osTypeValue := osFamily.NewFromGOOS(goosEnvValue)
	if osTypeValue == osFamily.Undefined {
		return Config{}, fmt.Errorf("undefined OS from env: GOOS=%s", goosEnvValue)
	}

	return Config{
		osFamily: osTypeValue,
	}, nil
}
