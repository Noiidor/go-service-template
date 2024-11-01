package config

import "github.com/Noiidor/go-service-template/pkg/config"

// Exported fields needed for env loader
type Config struct {
	AppHost string `env:"APP_HOST"`

	AppPlainHttpPort uint16 `env:"APP_PLAIN_HTTP_PORT,notEmpty"`

	DbHost string `env:"DB_HOST,notEmpty"`
	DbPort uint16 `env:"DB_PORT,notEmpty"`
	DbName string `env:"DB_NAME,notEmpty"`
	DbUser string `env:"DB_USER,notEmpty"`
	DbPass string `env:"DB_PASS,notEmpty"`
}

func Load() (*Config, error) {
	cfg, err := config.LoadStruct[Config]()
	return &cfg, err
}

func (c *Config) GetAppHost() string {
	return c.AppHost
}

func (c *Config) GetAppPlainHttpPort() uint16 {
	return c.AppPlainHttpPort
}

func (c *Config) GetDbName() string {
	return c.DbName
}

func (c *Config) GetDbPort() uint16 {
	return c.DbPort
}

func (c *Config) GetDbHost() string {
	return c.DbHost
}

func (c *Config) GetDbUser() string {
	return c.DbUser
}

func (c *Config) GetDbPass() string {
	return c.DbPass
}
