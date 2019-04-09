package config

//Factory creates Config
type Factory struct {
}

//NewFactory is a constructor of Factory
func NewFactory() *Factory {
	return &Factory{}
}

//CreateConfig creates Config
func (factory Factory) CreateConfig() (Config, error) {
	return Config{}, nil
}
