package dummy

import "testing"

func TestConstructorShouldCreatePluginWhenCalledWithoutParams(t *testing.T) {
	plugin := NewPlugin()
	if plugin == nil {
		t.Error("There should be no dependencies")
	}
}
