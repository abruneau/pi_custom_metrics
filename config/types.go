package config

import (
	"io"
	"strings"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config represents an object that can load and store configuration parameters
// coming from different kind of sources:
// - defaults
// - files
// - environment variables
// - flags
type Config interface {

	// API implemented by viper.Viper

	Set(key string, value interface{})
	SetDefault(key string, value interface{})
	SetFs(fs afero.Fs)
	IsSet(key string) bool

	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetSizeInBytes(key string) uint

	SetEnvPrefix(in string)
	BindEnv(input ...string) error
	SetEnvKeyReplacer(r *strings.Replacer)

	UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
	UnmarshalExact(rawVal interface{}, opts ...viper.DecoderConfigOption) error

	ReadInConfig() error
	ReadConfig(in io.Reader) error
	MergeConfig(in io.Reader) error

	AllSettings() map[string]interface{}
	AllKeys() []string

	AddConfigPath(in string)
	SetConfigName(in string)
	SetConfigFile(in string)
	SetConfigType(in string)
	ConfigFileUsed() string

	BindPFlag(key string, flag *pflag.Flag) error

	// API not implemented by viper.Viper and that have proven useful for our config usage

	// BindEnvAndSetDefault sets the default value for a config parameter and adds an env binding
	// in one call, used for most config options.
	//
	// If env is provided, it will override the name of the environment variable used for this
	// config key
	BindEnvAndSetDefault(key string, val interface{}, env ...string)
	// GetEnvVars returns a list of the non-sensitive env vars that the config supports
	GetEnvVars() []string
}