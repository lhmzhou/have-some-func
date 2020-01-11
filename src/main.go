package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

var (
	db  *sql.DB
	err error
)

type Data struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func main() {
	db, err = sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	os.Setenv("PORT", "8080")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	/*
		resp, err := http.PostForm("http://localhost:5003/create",
			url.Values{"ID": {"1234"}, "Name": {"hello"}, " Description" : {"hello Project"}, "URL" : {"www.google.com"}})



		//resp, err := http.Get("http://localhost:5003/index1")
		if err != nil {
			log.Fatal(err)
		}


		// Print the HTTP Status Code and Status Name
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			fmt.Println("HTTP Status is in the 2xx range")
		} else {
			fmt.Println("Error ",resp.StatusCode,resp.Status)
		}

	*/

	http.HandleFunc("/index", index)
}
