package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

const (
	defaultEnv        = "local"
	defaultServerAddr = "localhost"
	configPath        = "./internal/api/config"
	configType        = "yaml"
)

type Config struct {
	Port           int    `mapstructure:"port"`
	DBUser         string `mapstructure:"db_user"`
	DBPassword     string `mapstructure:"db_password"`
	DBHost         string `mapstructure:"db_host"`
	DBPort         int    `mapstructure:"db_port"`
	DBName         string `mapstructure:"db_name"`
	ServerAddress  string
	TrustedProxies []string `mapstructure:"trusted_proxies"`
}

var AppConfig *Config

func LoadConfig() error {
	env := getEnv("APP_ENV", defaultEnv)
	viper.SetConfigName(env)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file (%s): %w", env, err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Override config values with environment variables if available
	AppConfig.Port = getEnvAsInt("PORT", AppConfig.Port)
	AppConfig.DBUser = getEnv("DB_USER", AppConfig.DBUser)
	AppConfig.DBPassword = getEnv("DB_PASSWORD", AppConfig.DBPassword)
	AppConfig.DBHost = getEnv("DB_HOST", AppConfig.DBHost)
	AppConfig.DBPort = getEnvAsInt("DB_PORT", AppConfig.DBPort)
	AppConfig.DBName = getEnv("DB_NAME", AppConfig.DBName)

	// Set the ServerAddress based on environment
	serverAddr := getEnv("SERVER_ADDR", "local")
	if serverAddr == "docker" {
		AppConfig.ServerAddress = fmt.Sprintf("0.0.0.0:%d", AppConfig.Port)
	} else {
		AppConfig.ServerAddress = fmt.Sprintf("%s:%d", defaultServerAddr, AppConfig.Port)
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(name); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
