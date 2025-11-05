package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Config   string `yaml:"config"` //高级配置，例如charset等
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	LogLevel string `yaml:"log_level"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.User, m.Password, m.Host, m.Port, m.Db, m.Config)
}
