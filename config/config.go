package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string
	ServiceName string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	LogLevel string
	HttpPort int

	PlayMobileUsername string
	PlayMobilePassword string
	PlayMobileBaseURL  string
}

func Load() Config {
	envFileName := cast.ToString(".env")

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "info"))
	config.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", ""))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "root"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", ""))

	config.PlayMobileBaseURL = cast.ToString(getOrReturnDefault("PlayMobileBaseURL", "http://91.204.239.44"))
	config.PlayMobileUsername = cast.ToString(getOrReturnDefault("PlayMobileUsername", "beetoexpress"))
	config.PlayMobilePassword = cast.ToString(getOrReturnDefault("PlayMobilePassword", "jf91Tq4"))

	config.HttpPort = cast.ToInt(getOrReturnDefault("HTTP_PORT", "8008"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
