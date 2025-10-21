package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	System   System   `yaml:"system"`
	Log      Logger   `yaml:"logger"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	QQ       QQ       `yaml:"qq"`
	Upload   Upload   `yaml:"upload"`
	Redis    Redis    `yaml:"redis"`
	Es       Es       `yaml:"es"`
}
