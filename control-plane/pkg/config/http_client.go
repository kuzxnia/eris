package config

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func ProvideHttpClient(config *Config) *resty.Client {
	client := resty.New().
		SetDebug(config.HttpClient.Debug).
		// retry
		SetRetryCount(config.HttpClient.RetryCount).
		SetRetryWaitTime(time.Duration(config.HttpClient.RetryWaitTimeMs) * time.Millisecond).
		SetRetryMaxWaitTime(time.Duration(config.HttpClient.RetryMaxWaitTimeMs) * time.Millisecond).
		// timeout
		SetTimeout(time.Duration(config.HttpClient.TimeoutMs) * time.Millisecond).
		// headers
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})

	return client
}
