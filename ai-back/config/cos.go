package config

// Cos 项目基础配置

// Cos 项目基础配置

type Cos struct {
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	SecretID  string `yaml:"secretid"`
	SecretKey string `yaml:"secretkey"`
	BaseURL   string `yaml:"baseurl"`
}
