package app

type Config struct {
	JiraUsername string `envconfig:"JIRA_USERNAME"`
	JiraToken    string `envconfig:"JIRA_TOKEN"`
	JiraURL      string `envconfig:"JIRA_URL"`
}
