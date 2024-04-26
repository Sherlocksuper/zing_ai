package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	sid     string   // 服务运行ID
	Cos     *Cos     `json:"cos" yaml:"cos"`
	Redis   *Redis   `json:"redis" yaml:"redis"`
	MySQL   *Mysql   `json:"mysql" yaml:"mysql"`
	Cors    *Cos     `json:"cors" yaml:"cors"`
	Email   *Email   `json:"email" yaml:"email"`
	OpenAI  *OpenAI  `json:"openai" yaml:"openai"`
	Qianfai *Qianfai `json:"qianfai" yaml:"qianfai"`
}

var myConfig *Config = nil

func GetConfig() *Config {
	if myConfig != nil {
		return myConfig
	}
	var config Config

	var configStr, _ = os.ReadFile("./config.yaml")

	err := yaml.Unmarshal(configStr, &config)

	if err != nil {
		log.Error().Msg("config.yaml 解析失败：" + err.Error())
		return nil
	}
	myConfig = &config
	return &config
}
