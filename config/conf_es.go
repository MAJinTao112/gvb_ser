package config

import "fmt"

type Es struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (e Es) Url() string {
	return fmt.Sprintf("http://%s:%s", e.Host, e.Port)
}
