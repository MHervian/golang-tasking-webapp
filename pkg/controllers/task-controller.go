// Module for book controller
package controllers

import (
	"net/http"
	"text/template"
	// "github.com/MHervian/golang-tasking-webapp/pkg/models"
)

// Controller halaman utama
// https://dasarpemrogramangolang.novalagung.com/B-template-render-html.html
func Index(w http.ResponseWriter, r *http.Request) {
	var filePath = "C:/go-projects/src/golang-tasking-webapp/pkg/views/index.html"
	var tmpl, err = template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Controller halaman buat task
func Form(w http.ResponseWriter, r *http.Request) {
	var filePath = "C:/go-projects/src/golang-tasking-webapp/pkg/views/form.html"
	var tmpl, err = template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Controller halaman detail
func DisplayTaskById(w http.ResponseWriter, r *http.Request) {
	var filePath = "C:/go-projects/src/golang-tasking-webapp/pkg/views/detail.html"
	var tmpl, err = template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Controller halaman edit task
func EditTaskById(w http.ResponseWriter, r *http.Request) {
	var filePath = "C:/go-projects/src/golang-tasking-webapp/pkg/views/edit.html"
	var tmpl, err = template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Controller CREATE task

// Controller UPDATE status task

// Controller DELETE task
