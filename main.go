package main

import (
	"log"

	"github.com/imantung/jira_to_gsheet/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
