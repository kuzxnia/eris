package config

import (
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	HttpClient HttpClientConfig `mapstructure:"http_client"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type HttpClientConfig struct {
	Debug              bool `mapstructure:"debug"`
	RetryCount         int  `mapstructure:"retry_count"`
	RetryWaitTimeMs    int  `mapstructure:"retry_wait_time_ms"`
	RetryMaxWaitTimeMs int  `mapstructure:"retry_wait_max_time_ms"`
	TimeoutMs          int  `mapstructure:"timeout_ms"`
}

func ProvideConfig() *Config {
	viper.AutomaticEnv()
	viper.SetConfigName("properties")
	viper.AddConfigPath("./pkg/config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		exception.PanicLogging(err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		exception.PanicLogging(err)
	}

	return &configuration
}

var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
