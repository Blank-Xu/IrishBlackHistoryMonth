package config

import (
	"sync"

	"webservice/modules/http_server"
)

var (
	appConfig = new(Config)

	once sync.Once
)

type Config struct {
	HTTPServer http_server.Config
}

// Parse .
func Parse(filename string) error {
	em, err := NewEnvMap(filename)
	if err != nil {
		return err
	}

	cfg, err := readConfig(em)
	if err != nil {
		return err
	}

	once.Do(func() {
		appConfig = cfg
	})

	return nil
}

// GetConfig .
func GetConfig() Config {
	return *appConfig
}
