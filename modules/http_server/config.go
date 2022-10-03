package http_server

// Config for http server
type Config struct {
	Address                  string `json:"address" yaml:"address"`
	ReadTimeoutSeconds       int    `json:"read_timeout_seconds" yaml:"read_timeout_seconds"`
	ReadHeaderTimeoutSeconds int    `json:"read_header_timeout_seconds" yaml:"read_header_timeout_seconds"`
	WriteTimeoutSeconds      int    `json:"write_timeout_seconds" yaml:"write_timeout_seconds"`
	IdleTimeoutSeconds       int    `json:"idle_timeout_seconds" yaml:"idle_timeout_seconds"`
	MaxHeaderBytes           int    `json:"max_header_bytes" yaml:"max_header_bytes"`
}
