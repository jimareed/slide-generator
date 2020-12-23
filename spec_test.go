package main

import (
	"testing"
)

const checkMark = "\u2713"
const xMark = "\u2717"

const TEST_SPEC = "example spec"
const UPDATED_SPEC = "updated spec"

func TestCreateSpec(t *testing.T) {

	t.Log("a user:")

	id, err := createSpec(TEST_SPEC)

	if err == nil {
		t.Log("Should be able to create a spec.", checkMark)
	} else {
		t.Fatal("Should be able to create a spec.", xMark, err)
	}

	deleteSpec(id)
}

func TestReadSpec(t *testing.T) {

	id, err := createSpec(TEST_SPEC)
	if err != nil {
		t.Fatal("Should be able to create a spec.", xMark, err)
	}

	spec, err := readSpec(id)

	if err == nil {
		t.Log("Should be able to read a spec.", checkMark)
	} else {
		t.Fatal("Should be able to read a spec.", xMark, err)
	}

	if spec == TEST_SPEC {
		t.Log("Which should contain the same value as when created.", checkMark)
	} else {
		t.Fatal("Which should contain the same value as when created.", xMark, spec)
	}

	deleteSpec(id)
}

func TestReadEmptySpec(t *testing.T) {

	id, err := createSpec("")
	if err == nil {
		t.Log("Should be able to create an empty spec.", checkMark)
	} else {
		t.Fatal("Should be able to create an empty spec.", xMark, err)
	}

	spec, err := readSpec(id)

	if err == nil {
		t.Log("Should be able to read an empty spec.", checkMark)
	} else {
		t.Fatal("Should be able to read an empty spec.", xMark, err)
	}

	if spec == "" {
		t.Log("Which should be empty.", checkMark)
	} else {
		t.Fatal("Which should be empty.", xMark, spec)
	}

	deleteSpec(id)
}

func TestUpdateSpec(t *testing.T) {

	id, err := createSpec(TEST_SPEC)
	if err != nil {
		t.Fatal("Should be able to create a spec.", xMark, err)
	}

	err = updateSpec(id, UPDATED_SPEC)
	if err == nil {
		t.Log("Should be able to update a spec.", checkMark)
	} else {
		t.Fatal("Should be able to update a spec.", xMark, err)
	}

	spec, err := readSpec(id)
	if err != nil {
		t.Fatal("Should be able to read a spec.", xMark, err)
	}

	if spec == UPDATED_SPEC {
		t.Log("Should be able to read the updated value.", checkMark)
	} else {
		t.Fatal("Should be able to read the updated value.", xMark, spec)
	}

	deleteSpec(id)
}

func TestDeleteSpec(t *testing.T) {

	id, err := createSpec(TEST_SPEC)
	if err != nil {
		t.Fatal("Should be able to create a spec.", xMark, err)
	}

	err = deleteSpec(id)
	if err == nil {
		t.Log("Should be able to delete a spec.", checkMark)
	} else {
		t.Fatal("Should be able to delete a spec.", xMark, err)
	}

	_, err = readSpec(id)
	if err != nil {
		t.Log("Should not be able to read the deleted spec.", checkMark, err)
	} else {
		t.Fatal("Should not be able to read the deleted spec.", xMark)
	}
}
