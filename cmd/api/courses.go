package main

import (
	"coursego/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new course") 
}

func (app *application) showCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	course := data.Course{
		ID: id,
		CreatedAt: time.Now(),
		Title: "The Complete Go Course",
		Year: 2023,
		Runtime: 30,
		Subject: []string{"Go", "Programming"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}