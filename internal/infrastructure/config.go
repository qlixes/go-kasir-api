package infrastructure

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig AppConfig
	DbConfig  DbConfig
}

type AppConfig struct {
	Host string `mapstructure:"APP_HOST"`
	Port int    `mapstructure:"APP_PORT"`
}

type DbConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	DbName   string `mapstructure:"DB_NAME"`
}

func NewConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed load .env")
	}

	config := &Config{
		AppConfig: AppConfig{
			Host: viper.GetString("APP_HOST"),
			Port: viper.GetInt("APP_PORT"),
		},

		DbConfig: DbConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASS"),
			DbName:   viper.GetString("DB_NAME"),
		},
	}

	return config
}

func (c *DbConfig) PgSQLConnStr() string {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
	)

	return connStr
}
