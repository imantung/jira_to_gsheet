package app

import (
	"github.com/imantung/jira_to_gsheet/internal/app/jira"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	// fmt.Printf("%+v\n", cfg)

	jiraApp, err := jira.CreateJiraApp(cfg.JiraUsername, cfg.JiraToken, cfg.JiraURL)
	if err != nil {
		return err
	}

	jiraApp.PrintProjectList() // debug

	return nil
}
