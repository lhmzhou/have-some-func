package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestCreate(t *testing.T) {

	resp, err := http.PostForm("http://localhost:5003/create",
		url.Values{"ID": {"1234567"}, "Name": {"H.M."}, " Description": {"hello1 Project"}, "URL": {"www.google.com"}})

	//resp, err := http.Get("http://localhost:5003/index1")
	if err != nil {
		log.Fatal(err)
	}

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Error ", resp.StatusCode, resp.Status)
	}

}

func TestDelete(t *testing.T) {

	resp, err := http.PostForm("http://localhost:5003/delete",
		url.Values{"ID": {"7"}})

	//resp, err := http.Get("http://localhost:5003/index1")
	if err != nil {
		log.Fatal(err)
	}

	// print the HTTP status code and status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Error ", resp.StatusCode, resp.Status)
	}

}
