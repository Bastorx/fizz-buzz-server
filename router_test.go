package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestFizzBuzzServer(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	data := map[string]interface{}{
		"int1": 3,
		"int2": 5,
		"limit": 50,
		"str1": "fizz",
		"str2": "buzz",
	}
	expected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz"

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to JSON Marshal, got %v", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/fizz-buzz", ts.URL), "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error to read the body response, got %v", err)
	}
	if string(body) != expected {
		t.Fatalf("Expected body to be %v, got %v", expected, string(body))
	}
}