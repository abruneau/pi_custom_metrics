package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	//Conf is the global configuration object
	Conf         Config
	overrideVars = make(map[string]interface{})
)

func init() {
	// osinit()
	// Configure global configuration
	Conf = NewConfig("datadog", "", strings.NewReplacer(".", "_"))
	// Configuration defaults
	InitConfig(Conf)
}

// InitConfig initializes the config defaults on a config
func InitConfig(config Config) {
	config.BindEnvAndSetDefault("log_level", "info")
	config.BindEnvAndSetDefault("log_file", DefaultLogFile)
	config.BindEnvAndSetDefault("log_to_console", true)
	config.SetDefault("disable_file_logging", true)
}

// GetLogLevel returns the log level
func GetLogLevel() string {
	return Conf.GetString("log_level")
}

// SetupConfig fires up the configuration system
func SetupConfig(confFilePath string, configName string, failOnMissingFile bool) error {
	if configName != "" {
		Conf.SetConfigName(configName)
	}
	// set the paths where a config file is expected
	if len(confFilePath) != 0 {
		// if the configuration file path was supplied on the command line,
		// add that first so it's first in line
		Conf.AddConfigPath(confFilePath)
		// If they set a config file directly, let's try to honor that
		if strings.HasSuffix(confFilePath, ".yaml") {
			Conf.SetConfigFile(confFilePath)
		}
	}
	Conf.AddConfigPath(DefaultConfPath) // path to look for the config file in
	Conf.AddConfigPath(".")             // optionally look for config in the working directory
	// load the configuration
	var err error

	err = Load()

	// If `!failOnMissingFile`, do not issue an error if we cannot find the default config file.
	var e viper.ConfigFileNotFoundError
	if err != nil && (failOnMissingFile || !errors.As(err, &e) || confFilePath != "") {
		return fmt.Errorf("unable to load Datadog config file: %w", err)
	}
	return nil
}

// Load reads configs files and initializes the config module
func Load() error {
	return load(Conf)
}

func load(config Config) error {

	if err := config.ReadInConfig(); err != nil {
		if errors.Is(err, os.ErrPermission) {
			return fmt.Errorf("Error loading config: %v (check config file permissions)", err)
		}
		return fmt.Errorf("Error loading config: %v", err)
	}

	applyOverrides(config)
	return nil
}

// AddOverrides provides an externally accessible method for
// overriding config variables.
// This method must be called before Load() to be effective.
func AddOverrides(vars map[string]interface{}) {
	for k, v := range vars {
		overrideVars[k] = v
	}
}

// applyOverrides overrides config variables.
func applyOverrides(config Config) {
	for k, v := range overrideVars {
		config.Set(k, v)
	}
}
