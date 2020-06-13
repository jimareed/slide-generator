package main

import (
	"flag"
	"log"

	"github.com/jimareed/slide-generator/slides"
)

func main() {

	server := flag.Bool("server", false, "run in server mode")
	input := flag.String("input", "", "path to source")
	output := flag.String("output", "", "path to source")

	flag.Parse()

	if *input == "" {
		log.Fatal("usage: slide-generator -input <path> [-output <path>][-server]")
	}


	log.Print("reading deck from ", *input)

	deck, err := slides.Read(*input)
	if err == nil {
		log.Print(deck.Title, " read successful.")
	} else {
		log.Fatal(err)
	}

	if *output != "" {
		log.Print("writing ", deck.Title, " to ", *output)
		err = slides.Write(deck, *output)
		if err == nil {
			log.Print(deck.Title, " write successful.")
		} else {
			log.Fatal(err)
		}
	}

	if *server {
		log.Print("server mode is not implemented yet.")
	}
}
