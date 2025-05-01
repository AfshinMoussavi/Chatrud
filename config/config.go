package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port string
}

type LoggerConfig struct {
	Development       string
	DisableCaller     string
	DisableStacktrace string
	Encoding          string
	Level             string
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  string
	PgDriver           string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func InitConfig(filename string) (*Config, error) {
	loadConfig, err := LoadConfig(filename)
	if err != nil {
		return nil, err
	}
	parseConfig, err := ParseConfig(loadConfig)
	if err != nil {
		return nil, err
	}
	return parseConfig, nil
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(filename + ".yml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	fmt.Println("Raw config map:", v.AllSettings())

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
