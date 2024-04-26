package config

// Qianfai 项目基础配置
type Qianfai struct {
	DiffussionXLBaseURL string `yaml:"diffussion_xl_baseurl"`
	QFBaseURL           string `yaml:"qf_baseurl"`
	QFTokenURL          string `yaml:"qf_tokenurl"`
	QFAPIKey            string `yaml:"qf_apikey"`
	QFAPISecretKey      string `yaml:"qf_apisecretkey"`
}
