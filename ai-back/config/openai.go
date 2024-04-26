package config

// OpenAI 项目基础配置

type OpenAI struct {
	OpenAIToken          string `yaml:"openai_token"`
	BaseURL              string `yaml:"openai_baseurl"`
	Model                string `yaml:"model"`
	MaxTokens            int    `yaml:"max_tokens"`
	DefaultSystemMessage string `yaml:"default_system_message"`
}
