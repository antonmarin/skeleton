package config

import (
	"fmt"
	"github.com/antonmarin/skeleton/config/enum/osFamily"
	"os"
)

//Factory creates Config
type Factory struct {
}

//NewFactory is a constructor of Factory
func NewFactory() *Factory {
	return &Factory{}
}

//CreateConfig creates Config
func (factory Factory) CreateConfig() (Config, error) {
	osTypeValue := osFamily.NewFromGOOS(os.Getenv("GOOS"))
	if osTypeValue == osFamily.Undefined {
		return Config{}, fmt.Errorf("undefined OS")
	}

	return Config{
		osFamily: osTypeValue,
	}, nil
}
