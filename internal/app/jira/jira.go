package jira

import (
	"fmt"

	gojira "github.com/andygrunwald/go-jira"
)

type JiraApp struct {
	*gojira.Client
}

func CreateJiraApp(username, token, url string) (*JiraApp, error) {
	tp := gojira.BasicAuthTransport{
		Username: username,
		Password: token,
	}
	client, err := gojira.NewClient(tp.Client(), url)
	if err != nil {
		return nil, err
	}
	return &JiraApp{Client: client}, nil
}

func (j *JiraApp) PrintProjectList() {
	projectList, _, err := j.Project.GetList()
	if err != nil {
		fmt.Print("ERR: ")
		fmt.Println(err.Error())
		return
	}
	for _, project := range *projectList {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}
}
