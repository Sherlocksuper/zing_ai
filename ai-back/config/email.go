package config

type Email struct {
	EmailAuthorEmail string `yaml:"email_author_email"`
	EmailPassword    string `yaml:"email_password"`
	EmailTitle       string `yaml:"email_title"`
	EmailTemplate    string `yaml:"email_template"`
}
