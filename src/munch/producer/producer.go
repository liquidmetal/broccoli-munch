package producer

import (
	"bytes"
	"fmt"
	"html/template"
	"munch/newsletters"
	"munch/stories"
)

type all_information struct {
	Title      string
	NumStories int
}

// Produces a good looking HTML for the given newsletter and stories
func Produce(newsletter *newsletters.Newsletter, stories []*stories.Story) string {
	var output bytes.Buffer

	var everything all_information
	everything.Title = newsletter.GetTitle()
	everything.NumStories = len(stories)

	tmpl, err := template.ParseFiles("./src/munch/producer/email.html")
	if err != nil {
		fmt.Printf("There was an error loading the template file: %s\n", err)
	}
	tmpl.Execute(&output, everything)

	return output.String()
}
