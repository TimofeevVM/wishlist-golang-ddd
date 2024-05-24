package configs

import "os"

type Config struct {
	addr          string
	databaseDsn   string
	stathouseHost string
}

func (c *Config) Addr() string {
	return c.addr
}

func (c *Config) DatabaseDsn() string {
	return c.databaseDsn
}

func (c *Config) StathouseHost() string {
	return c.stathouseHost
}

func LoadConfig() *Config {
	return &Config{
		addr:          os.Getenv("APP_ADDR"),
		databaseDsn:   os.Getenv("APP_DATABASE_DSN"),
		stathouseHost: os.Getenv("APP_STATHOUSE_HOST"),
	}
}
