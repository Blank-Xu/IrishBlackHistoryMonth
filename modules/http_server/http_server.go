package http_server

import (
	"net/http"
	"time"
)

// NewServer create a http server by config
func NewServer(config Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr: config.Address,

		Handler: handler,

		// these times should be adjusted in real situations
		ReadTimeout:       time.Second * time.Duration(config.ReadTimeoutSeconds),
		ReadHeaderTimeout: time.Second * time.Duration(config.ReadHeaderTimeoutSeconds),
		WriteTimeout:      time.Second * time.Duration(config.WriteTimeoutSeconds),
		IdleTimeout:       time.Second * time.Duration(config.IdleTimeoutSeconds),

		MaxHeaderBytes: config.MaxHeaderBytes,
	}
}
