package config

import "webservice/modules/http_server"

// readConfig .
func readConfig(em EnvMap) (*Config, error) {
	return &Config{
		HTTPServer: http_server.Config{
			Address:                  em.String(_HTTP_ADDRESS),
			ReadTimeoutSeconds:       em.MustInt(_HTTP_READ_TIMEOUT_SECONDS),
			ReadHeaderTimeoutSeconds: em.MustInt(_HTTP_READ_HEADER_TIMEOUT_SECONDS),
			WriteTimeoutSeconds:      em.MustInt(_HTTP_WRITE_TIMEOUT_SECONDS),
			IdleTimeoutSeconds:       em.MustInt(_HTTP_IDLE_TIMEOUT_SECONDS),
			// MaxHeaderBytes:           1 << 10, // 1024
		},
	}, nil
}
