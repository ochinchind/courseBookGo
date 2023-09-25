package main

import (
	"fmt"
	"net/http"
)

func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new course") 
}


func (app *application) showCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id) 
}