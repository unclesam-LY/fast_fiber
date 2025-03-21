package config

type Config struct {
	System System `yaml:"system"`
	DB     DB     `yaml:"db"`
	Redis  Redis  `yaml:"redis"`
	Log    Log    `yaml:"log"`
	Jwt    Jwt    `yaml:"jwt"`
}
