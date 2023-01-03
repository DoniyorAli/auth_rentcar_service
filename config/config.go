package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App string
	AppVersion string
	Environment string	// devlopment, staging, production
	GRPCPort string

	DefaultOffset string
	DefaultLimit string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	SecretKey string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "article"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":7001"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "article"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "1.0.0"))


	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "127.0.0.1"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "article_service_db"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "article_db_user"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "article_db_pswd"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "YourSecretKey"))
 
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}