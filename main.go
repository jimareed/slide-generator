package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jimareed/drawing"
)

var filePath = "./slides"
var autoPlay = true

func drawingToHtml(path string, name string, autoPlay bool) (string, error) {

	filename := path + "/" + name + ".draw"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return "error opening file", err
	}
	text := string(content)

	d, err := drawing.FromString(text)
	if err != nil {
		log.Print(err)
		return "invalid drawing", err
	}

	s, err := drawing.ToHtml(d, autoPlay)

	return s, err
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		id = "default"
	}

	if id != "favicon.ico" {
		content, err := drawingToHtml(filePath, id, autoPlay)
		if err != nil {
			content = "Invalid File"
		}

		io.WriteString(w, content)
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	input := flag.String("input", "", "path to source")
	help := flag.Bool("help", false, "help")

	flag.Parse()

	if *help {
		log.Fatal("usage: slide-generator [-input <path>][-output <path>][-server][-help]")
	}

	if *input != "" {
		filePath = *input
	}

	log.Print("reading from ", filePath)

	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods("GET")
	r.HandleFunc("/{id}", getHandler).Methods("GET")

	log.Printf("Server started at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
