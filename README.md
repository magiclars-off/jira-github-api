# Jira-Github releasenotes

this is a simple command line tool written in Golang. It looks for jira identifiers in github commmits, then finds the corresponding tickets in Jira and creates releasenotes style response.

## Build and Run
from root folder: 
    
    make build

before you run fill in the config.toml or set respective env variables:

    JIRA_URL               validate:"required"
	JIRA_USER              validate:"required"
    JIRA_TOKEN             validate:"required"
    JIRA_TICKET_IDENTIFIER validate:"required"

	GIT_OWNER     validate:"required"
	GIT_USERNAME  validate:"required"
	GIT_TOKEN     validate:"required"

To run

    ./bin/jira-api -c=config.toml -b=tags/v0.3 -h=master -r=deliberative-debate

with:

    -c      = config location
    -b      = base of github tree
    -h      = head of github tree
    -r      = repository name