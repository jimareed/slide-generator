package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jimareed/drawing"
)

var filePath = "./slides"

func drawingToHtml(path string, name string) (string, error) {

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

	s, err := drawing.ToSvg(d)

	return s, err
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		id = "default"
	}

	if id != "favicon.ico" {
		content, err := drawingToHtml(filePath, id)
		if err != nil {
			content = "Invalid File"
		}

		io.WriteString(w, "<html><body>"+content+"</body></html>\n")
	}
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

	log.Print("reading from ", filePath)

	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods("GET")
	r.HandleFunc("/{id}", getHandler).Methods("GET")

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
