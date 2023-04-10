package initializers

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServicePort    string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// Loads parameters from environment
func LoadFromEnv() (config Config, err error) {

	config.DBHost = os.Getenv("POSTGRES_HOST")
	if config.DBHost == "" {
		err = errors.New("DBHost is empty")
		return
	}

	config.DBName = os.Getenv("POSTGRES_DB")
	if config.DBName == "" {
		err = errors.New("DBName is empty")
		return
	}

	config.DBUserName = os.Getenv("POSTGRES_USER")
	if config.DBUserName == "" {
		err = errors.New("DBUserName is empty")
		return
	}

	config.DBUserPassword = os.Getenv("POSTGRES_PASSWORD")
	if config.DBUserPassword == "" {
		err = errors.New("DBUserPassword is empty")
		return
	}

	config.DBPort = os.Getenv("POSTGRES_PORT")
	if config.DBPort == "" {
		config.DBPort = strconv.Itoa(5432)
	}

	config.ServicePort = os.Getenv("SERVICE_PORT")
	if config.DBPort == "" {
		config.DBPort = strconv.Itoa(8000)
	}
	config.ClientOrigin = os.Getenv("CLIENT_ORIGIN")

	return
}
