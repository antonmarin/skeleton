package apply

type apply struct {
	configFactory  configFactory
	configObserver configObserver
}

//Constructs Apply use case.
// ConfigFactory creates configuration and configObserver dispatches config to handlers
func NewApply(configFactory configFactory, configObserver configObserver) *apply {
	return &apply{configFactory: configFactory, configObserver: configObserver}
}

//Execute use case.
func (useCase apply) Execute() error {
	config, err := useCase.configFactory.ConstructConfig()
	if err != nil {
		return err
	}

	if err = useCase.configObserver.Update(config); err != nil {
		return err
	}

	return nil
}

type configFactory interface {
	ConstructConfig() (Config, error)
}

type Config interface {
}

type configObserver interface {
	Update(Config) error
}
