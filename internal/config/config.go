package config

import "os"

type Config struct {
	GRPCListen string
	Env        string
}

func New() *Config {
	return &Config{
		GRPCListen: getEnv("GRPC_LISTEN", ":8888"),
		Env:        getEnv("ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
