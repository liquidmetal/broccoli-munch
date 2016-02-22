package main

import "fmt"
import "munch/sources"

// Worker process for fetching data from a given source

func main() {
    // Step 1: Establish connection to the message queue
    // Step 2: Wait for a task to show up in the queue
    // Step 3: Fetch the data for the given task, process it, generate stories
    // Step 4: Put the data into a DB somewhere
    // Step 5: Repeat

    fmt.Printf("This works!\n")
    f := sources.NewSourceRss("Techcrunch", "http://feeds.feedburner.com/TechCrunch/?fmt=xml", 0);
    //f := sources.NewSourceRss("Utkarsh Sinha", "http://utkarshsinha.com/index.xml", 0)
    f.FetchNewData()
}
