package apply

import "testing"

func TestApply_Execute(t *testing.T) {
	factory := new(ConfigFactorySpy)
	observer := new(observerSpy)
	useCase := NewApply(factory, observer)

	_ = useCase.Execute()

	if factory.CallCounter != 1 {
		t.Errorf("Factory should be called to create Config")
	}
	if observer.CallCounter != 1 {
		t.Errorf("Observer should be called to create skeleton")
	}
}

type ConfigFactorySpy struct {
	CallCounter int
}

func (f *ConfigFactorySpy) ConstructConfig() (Config, error) {
	f.CallCounter++

	config := new(testConfig)
	return config, nil
}

type testConfig struct {
}

type observerSpy struct {
	CallCounter int
}

func (o *observerSpy) Update(Config) error {
	o.CallCounter++
	return nil
}
