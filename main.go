package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jimareed/slides"
)

var filePath = "./slides"

func getHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		id = "default"
	}

	content := ""
	deck, err := slides.Read(filePath, id)
	if err != nil {
		content = "Invalid File"
	} else {
		content, err = slides.ToHtml(deck)
		if err != nil {
			content = "Error"
		}
	}

	io.WriteString(w, "<html><body>"+content+"</body></html>\n")
}

func main() {

	input := flag.String("input", "", "path to source")
	help := flag.Bool("help", false, "help")

	flag.Parse()

	if *help {
		log.Fatal("usage: slide-generator [-input <path>][-output <path>][-server][-help]")
	}

	if *input != "" {
		filePath = *input
	}

	log.Print("reading deck from ", filePath)

	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods("GET")
	r.HandleFunc("/{id}", getHandler).Methods("GET")

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
