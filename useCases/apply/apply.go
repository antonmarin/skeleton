package apply

import "github.com/antonmarin/skeleton/config"

type apply struct {
	configFactory configFactory
	plugin        Plugin
}

//Constructs Accept use case.
func NewApply(configFactory configFactory, plugin Plugin) *apply {
	return &apply{configFactory: configFactory, plugin: plugin}
}

//Execute use case.
func (useCase apply) Execute() error {
	currentConfig, err := useCase.configFactory.ConstructConfig()
	if err != nil {
		return err
	}

	return useCase.plugin.Accept(currentConfig)
}

type configFactory interface {
	ConstructConfig() (config.Config, error)
}

type Plugin interface {
	Accept(config.Config) error
}
