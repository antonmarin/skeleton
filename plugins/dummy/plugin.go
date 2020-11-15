package dummy

import (
	"github.com/antonmarin/skeleton/config"
)

type Plugin struct {
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p Plugin) Accept(config config.Config) error {
	return nil
}
