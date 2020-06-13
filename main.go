package main

import (
	"flag"
	"log"

	"github.com/jimareed/slide-generator/slides"
)

func main() {

	server := flag.Bool("server", false, "run in server mode")
	path  := flag.String("path", "", "path to source")

	flag.Parse()

	if *path == "" {
		log.Fatal("usage: slide-generator -path <path> [-server]")
	}

	output, err := slides.Execute(*path)		
	if err == nil {
		log.Print("output: ", output)
	} else {
		log.Fatal(err)
	}

	if *server {
		log.Print("server mode is not implemented yet.")
	}
}
