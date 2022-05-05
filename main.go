package main

import (
	"flag"
	"fmt"
	"log"

	jira "github.com/andygrunwald/go-jira"
)

var (
	username = flag.String("username", "some@email.com", "Jira Username")
	password = flag.String("password", "abcded", "Jira Token")
	baseURL  = flag.String("url", "https://your-org.atlassian.net", "Jira Base URL")
)

func main() {
	flag.Parse()

	fmt.Printf("Create client for '%s'\n", *baseURL)
	client, err := createClient()
	fatalIfError(err)

	projectList, _, err := client.Project.GetList()
	fatalIfError(err)

	// Range over the projects and print the key and name
	for _, project := range *projectList {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}
}

func createClient() (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: *username,
		Password: *password,
	}
	return jira.NewClient(tp.Client(), *baseURL)
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
