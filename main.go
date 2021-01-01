package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jimareed/drawing"
	"github.com/rs/cors"
)

type Specification struct {
	Id   string `json:"id"`
	Spec string `json:"specification"`
}

var filePath = "./slides"
var autoPlay = true

func specId2Name(id string) string {
	if len(id) == 0 {
		return ""
	}
	if id[0] >= '0' && id[0] <= '9' {
		return "slideshow-" + id
	}

	return id
}

func drawingToHtml(path string, name string, autoPlay bool) (string, error) {

	filename := path + "/" + name + ".draw"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return "error opening file", err
	}
	text := string(content)

	d := drawing.Drawing{}

	if text != "" {
		d, err = drawing.FromString(text)
		if err != nil {
			log.Print(err)
			return "invalid drawing", err
		}
	}

	s, err := drawing.ToHtml(d, autoPlay)

	return s, err
}

func getSlideshowsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		id = "default"
	}

	if id != "favicon.ico" {
		content, err := drawingToHtml(filePath, specId2Name(id), autoPlay)
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
		log.Fatal("usage: slideshow [-input <path>][-help]")
	}

	if *input != "" {
		filePath = *input
		setSpecPath(filePath)
	}

	log.Print("reading from ", filePath)

	r := mux.NewRouter()
	r.HandleFunc("/slideshows/{id}", getSlideshowsHandler).Methods("GET")
	r.HandleFunc("/specs", postSpecsHandler).Methods("POST")
	r.HandleFunc("/specs/{id}", getSpecsHandler).Methods("GET")
	r.HandleFunc("/specs/{id}", putSpecsHandler).Methods("PUT")
	r.HandleFunc("/specs/{id}", deleteSpecsHandler).Methods("DELETE")

	// For dev only - Set up CORS so React client can consume our API
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	log.Printf("Server started at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsWrapper.Handler(r)))
}

func getSpecsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var spec = Specification{}
	spec.Id = id

	s, err := readSpec(id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	spec.Spec = s

	payload, _ := json.Marshal(spec)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func putSpecsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var spec Specification
	json.Unmarshal(reqBody, &spec)

	err := updateSpec(id, spec.Spec)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	payload, _ := json.Marshal(spec)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func postSpecsHandler(w http.ResponseWriter, r *http.Request) {

	id, err := createSpec("")
	if err != nil {
		log.Printf("Error creating spec:" + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var spec = Specification{}
	spec.Id = id
	spec.Spec = ""

	payload, _ := json.Marshal(spec)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func deleteSpecsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var spec = Specification{}
	spec.Id = id

	err := deleteSpec(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	payload, _ := json.Marshal(spec)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
