package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jimareed/slides"
)

var mainDeck = slides.SlideDeck{}

func getHandler(w http.ResponseWriter, r *http.Request) {

	mainDeck, _ = slides.Read("./slides")
	content, err := slides.ToHtml(mainDeck)
	if err != nil {
		content = "Error"
	}

	io.WriteString(w, "<html><body>"+content+"</body></html>\n")
}

func main() {

	server := flag.Bool("server", false, "run in server mode")
	input := flag.String("input", "", "path to source")
	output := flag.String("output", "", "path to source")
	help := flag.Bool("help", false, "help")

	flag.Parse()

	if *help {
		log.Fatal("usage: slide-generator [-input <path>][-output <path>][-server][-help]")
	}

	if *input == "" {
		*input = "./slides"
	}

	log.Print("reading deck from ", *input)

	mainDeck, err := slides.Read(*input)
	if err == nil {
		log.Print(mainDeck.Title, " read successful.")
	} else {
		log.Fatal(err)
	}

	if *output != "" {
		log.Print("writing ", mainDeck.Title, " to ", *output)
		err = slides.Write(mainDeck, *output)
		if err == nil {
			log.Print(mainDeck.Title, " write successful.")
		} else {
			log.Fatal(err)
		}
	}

	if *server {
		r := mux.NewRouter()
		r.HandleFunc("/", getHandler).Methods("GET")

		log.Print("Server started on localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}
}
