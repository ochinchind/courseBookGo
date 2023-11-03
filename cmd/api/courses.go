package main

import (
	"coursego/internal/data"
	"coursego/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title	 string		 `json:"title"`
		Year	 int32		 `json:"year"`
		Runtime  data.Runtime `json:"runtime"` 
		Subjects []string 	 `json:"subjects"`
	}
	

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	course := &data.Course { 
		Title: input.Title,
		Year: input.Year,
		Runtime: input.Runtime,
		Subjects: input.Subjects,
	}

	v := validator.New()

	if data.ValidateCourse(v, course); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Courses.Insert(course)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/courses/%d", course.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"course": course}, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Fprintf(w, "%+v\n", input) 
}

func (app *application) showCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	course, err := app.models.Courses.Get(id)
	if err != nil { 
		switch {
			case errors.Is(err, data.ErrRecordNotFound): 
				app.notFoundResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	course, err := app.models.Courses.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r) 
		default:
			app.serverErrorResponse(w, r, err) 
		}
		return
	}

	var input struct {
		Title string `json:"title"`
		Year int32 `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Subjects []string `json:"subjects"`
	}

	err = app.readJSON(w, r, &input) 
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	course.Title = input.Title
	course.Year = input.Year
	course.Runtime = input.Runtime
	course.Subjects = input.Subjects

	v := validator.New()
	if data.ValidateCourse(v, course); !v.Valid() { 
		app.failedValidationResponse(w, r, v.Errors) 
		return
	}

	err = app.models.Courses.Update(course) 
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil) 
	if err != nil {
		app.serverErrorResponse(w, r, err) 
	}
}