package config

import (
	"github.com/antonmarin/skeleton/config/params/osFamily"
)

//Config is a user configuration of usage
type Config struct {
	osFamily osFamily.OsFamily
}
