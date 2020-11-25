package config

// Jira contains jira settings
type Jira struct {
	BaseURL          string `env:"JIRA_URL"               validate:"required"`
	Username         string `env:"JIRA_USER"              validate:"required"`
	Password         string `env:"JIRA_TOKEN"             validate:"required"`
	TicketIdentifier string `env:"JIRA_TICKET_IDENTIFIER" validate:"required"`
}

// Git contains git settings
type Git struct {
	Owner    string `env:"GIT_OWNER"     validate:"required"`
	Username string `env:"GIT_USERNAME"  validate:"required"`
	Token    string `env:"GIT_TOKEN"     validate:"required"`
}

// Modus is the modus_type enum
type Modus uint8

// Enum values for modus_type
const (
	ModusSupportInstructions Modus = iota
	ModusReleasenotes
)

var modes = map[Modus]string{
	ModusSupportInstructions: `support-instructions`,
	ModusReleasenotes:        `releasenotes`,
}

// String returns the string value of the modus_type
func (m Modus) String() string {
	return modes[m]
}
