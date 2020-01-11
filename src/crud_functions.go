package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type ID struct {
	id int64
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
	rows, err := db.Query("SELECT * FROM database")
	checkInternalServerError(err, w)
	var funcMap = template.FuncMap{
		"addOne": func(n int) int {
			return n + 1
		},
	}

	var projects []Project
	var project Project

	for rows.Next() {
		err = rows.Scan(&project.id, &project.name,
			&project.description, &project.url)
		checkInternalServerError(err, w)
		projects = append(projects, project)
	}
	t, err := template.New("index.html").Funcs(funcMap).ParseFiles("tmpl/index.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, projects)
	checkInternalServerError(err, w)

}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/index", 301)
	}
	var project Project
	project.id, err = strconv.ParseInt(r.FormValue("id"), 10, 64)

	fmt.Print(err)

	if err != nil {
		http.Redirect(w, r, "/index", 301)
		return
	}

	project.name = r.FormValue("name")
	project.description = r.FormValue("description")
	project.url = r.FormValue("url")
	fmt.Println(project)

	// Save to database
	stmt, err := db.Prepare(`
		INSERT INTO database(id, name, description, url) VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		fmt.Println("Prepare query error")
		panic(err)
	}
	_, err = stmt.Exec(project.id, project.name,
		project.description, project.url)

	if err != nil {
		fmt.Println("Execute query error")
		panic(err)
	}
	http.Redirect(w, r, "/index", 301)
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Redirect(w, r, "/index", 301)
	}
	var project Project
	project.id, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	project.name = r.FormValue("name")
	project.description = r.FormValue("description")
	project.url = r.FormValue("url")
	fmt.Println(project)

	stmt, err := db.Prepare(`
		UPDATE database SET name=?, description=?, url=?
		WHERE id=?
	`)
	checkInternalServerError(err, w)
	res, err := stmt.Exec(project.name, project.description, project.url, project.id)
	checkInternalServerError(err, w)
	_, err = res.RowsAffected()
	checkInternalServerError(err, w)
	http.Redirect(w, r, "/index", 301)
}

func delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/index", 301)
	}
	var costId, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	fmt.Println(costId)

	stmt, err := db.Prepare("DELETE FROM database WHERE id=?")
	checkInternalServerError(err, w)
	res, err := stmt.Exec(costId)
	checkInternalServerError(err, w)
	_, err = res.RowsAffected()
	checkInternalServerError(err, w)
	http.Redirect(w, r, "/index", 301)

}

func checkInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createNew(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/create.html")
}

func edit(w http.ResponseWriter, r *http.Request) {
	var x, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	td := ID{x}

	t, err := template.New("edit.html").ParseFiles("tmpl/edit.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, td)
	checkInternalServerError(err, w)
}
