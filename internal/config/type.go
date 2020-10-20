package config

// Jira contains jira settings
type Jira struct {
	BaseURL  string `env:"JIRA_URL"   validate:"required"`
	Username string `env:"JIRA_USER"  validate:"required"`
	Password string `env:"JIRA_TOKEN" validate:"required"`
}

// Git contains git settings
type Git struct {
	BaseURL string `env:"GIT_URL"   validate:"required"`
	Token   string `env:"GIT_TOKEN" validate:"required"`
}
