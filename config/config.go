package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	ServiceHost string
	HTTPPort    string

	Postgres struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	}
}

func Load() Config {

	config := Config{}

	config.ServiceHost = cast.ToString(GetOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.HTTPPort = cast.ToString(GetOrReturnDefaultValue("HTTP_PORT", "8001"))

	config.Postgres.Host = cast.ToString(GetOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.Postgres.Port = cast.ToInt(GetOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.Postgres.Username = cast.ToString(GetOrReturnDefaultValue("POSTGRES_USERNAME", "username"))
	config.Postgres.Password = cast.ToString(GetOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres"))
	config.Postgres.Database = cast.ToString(GetOrReturnDefaultValue("POSTGRES_DATABASE", "beeontime"))

	return config

}
func GetOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
