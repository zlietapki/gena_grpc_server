// start name:top
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	//start name:struct type:add
	GRPCListen string
	//start name:body
	Env string
}

func New() *Config {
	godotenv.Load()

	return &Config{
		//start name:return type:add
		GRPCListen: getEnv("GRPC_LISTEN", ":8888"),
		//start name:post_return
		Env: getEnv("ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
