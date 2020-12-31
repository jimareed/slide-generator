package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var nextId = 1
var specPath = "./slides"

func setSpecPath(path string) {
	specPath = path
}

func fileName(id string) string {

	return specPath + "/slideshow-" + id + ".draw"
}

func writeSpec(id string, spec string) error {

	d := []byte(spec)

	err := ioutil.WriteFile(fileName(id), d, 0644)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func createSpec(spec string) (string, error) {

	id := fmt.Sprintf("%d", nextId)

	err := writeSpec(id, spec)
	if err != nil {
		log.Print(err)
		return "", err
	}

	nextId++

	return id, nil
}

func readSpec(id string) (string, error) {

	spec := ""

	d, err := ioutil.ReadFile(fileName(id))
	if err != nil {
		log.Print(err)
		return spec, err
	}
	spec = string(d)

	return spec, nil
}

func updateSpec(id string, spec string) error {
	return writeSpec(id, spec)
}

func deleteSpec(id string) error {
	return os.Remove(fileName(id))
}
