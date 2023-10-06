package main

import (
	"coursego/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title	string	`json:"title"`
		Year	int32	`json:"year"`
		Runtime int32	`json:"runtime"`
		Subjects []string `json:"subjects"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input) 
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
		Subjects: []string{"Go", "Programming"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}