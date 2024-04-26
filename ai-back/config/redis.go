package config

// Redis 项目基础配置
type Redis struct {
	RedisAddress  string `json:"address" yaml:"address"`
	RedisPassword string `json:"password" yaml:"password"`
	RedisDb       int    `json:"db" yaml:"db"`
}
