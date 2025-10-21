package config

type Redis struct {
	IP       string `yaml:"ip" json:"ip"`
	Port     string `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	PoolSize int    `yaml:"pool_size" json:"pool_size"`
}

func (r *Redis) Addr() string {
	return r.IP + ":" + r.Port
}
