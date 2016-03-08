package producer

import (
	"bytes"
	"fmt"
	"github.com/dchest/htmlmin"
	"html/template"
	"munch/newsletters"
	"munch/stories"
)

type all_information struct {
	Title      string
	NumStories int

	TextStories     []*stories.Story
	VideoStories    []*stories.Story
	ActivityStories []*stories.Story
}

// Produces a good looking HTML for the given newsletter and stories
func Produce(newsletter *newsletters.Newsletter, all_stories []*stories.Story) string {
	var output bytes.Buffer

	var everything all_information
	everything.Title = newsletter.GetTitle()
	everything.NumStories = len(all_stories)

	// Group stories based on source type
	everything.TextStories = make([]*stories.Story, 0, 10)
	for _, story := range all_stories {
		everything.TextStories = append(everything.TextStories, story)
	}

	tmpl, err := template.ParseFiles("./src/munch/producer/email.html")
	if err != nil {
		fmt.Printf("There was an error loading the template file: %s\n", err)
	}
	tmpl.Execute(&output, everything)

	b := output.Bytes()
	len_total := len(b)

	minified, err := htmlmin.Minify(b, &htmlmin.Options{MinifyScripts: true, MinifyStyles: true, UnquoteAttrs: true})

	len_min := len(minified)

	if err != nil {
		fmt.Printf("There was an error in minifying... using the unminified version\n")
		return output.String()
	}
	fmt.Printf("Total saving of %d bytes\n", (len_total - len_min))
	return string(minified)
}
