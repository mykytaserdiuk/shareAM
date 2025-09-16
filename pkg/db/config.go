package db

import (
	"fmt"
	"net"
)

type Config struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	DBName         string `yaml:"dbName"`
	SSLMode        string `yaml:"sslmode"`
	IsolationLevel string `yaml:"isolationLevel"`
}

// URL returns a postgres connection URL
func (c *Config) URL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		c.Username, c.Password, net.JoinHostPort(c.Host, c.Port), c.DBName, c.SSLMode,
	)
}
