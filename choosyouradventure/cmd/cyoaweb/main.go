package main

import (
	"flag"
	cyoa "gophercises/choosyouradventure"
	"log"
	"net/http"
	"os"
)

func main() {
	fileName := flag.String("filename", "../../story.json", "Filename where the complete story is present")
	flag.Parse()

	//Check if the file is present on the server or not
	file, err := os.Open(*fileName)

	//In case  of error we have to log that fatal exception and go out from here
	if err != nil {
		log.Fatal("Not able to find the file Specified")
	}

	//var story cyoa.Story
	story, err := cyoa.JsonStory(file)

	if err != nil {
		log.Fatal("Error in conversion to the Map")
	}
	//fmt.Println(story)

	//Starting the server and calling the handler to do the task
	http.ListenAndServe(":8080", cyoa.NewHandler(story))

}
