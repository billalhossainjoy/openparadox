package config

import "time"


type Config struct {
	Addr string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	ReqTimeout time.Duration
	ShutdownTimeout time.Duration
}

func Load() Config {
	return Config {
		Addr: ":8080",
		ReqTimeout: 5 * time.Second,
	}
}