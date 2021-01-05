package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config represents an object that can load and store configuration parameters
// coming from different kind of sources:
// - defaults
// - files
// - environment variables
type config struct {
	*viper.Viper
	configEnvVars []string
}

// GetEnvVars implements the Config interface
func (c *config) GetEnvVars() []string {
	return c.configEnvVars
}

func (c config) BindEnvAndSetDefault(key string, val interface{}, env ...string) {
	c.SetDefault(key, val)
	c.BindEnv(append([]string{key}, env...)...) //nolint:errcheck
}

// NewConfig returns a new Config object.
func NewConfig(name string, envPrefix string, envKeyReplacer *strings.Replacer) Config {
	config := config{
		Viper: viper.New(),
	}
	config.SetConfigName(name)
	config.SetEnvPrefix(envPrefix)
	config.SetEnvKeyReplacer(envKeyReplacer)
	config.SetTypeByDefaultValue(true)
	return &config
}
