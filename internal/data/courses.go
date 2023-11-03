package data

import (
	"database/sql" 
	"encoding/json"
	"fmt"
	"time"
	"coursego/internal/validator"
	"github.com/lib/pq"
)

type Course struct {
	ID 			int64 				`json:"id"`	
	CreatedAt 	time.Time  			`json:"-"`
	Title 		string	 			`json:"title"`
	Year 		int32	 			`json:"year,omitempty"`
	Runtime     Runtime             `json:"-"`
	Subjects 	[]string 			`json:"subjects,omitempty"`
	Version 	int32 				`json:"version"`
}

func (m Course) MarshalJSON() ([]byte, error) {
	var runtime string

	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	type CourseAlias Course
	aux := struct { 
		CourseAlias
		Runtime string `json:"runtime,omitempty"` 
	}{
		CourseAlias: CourseAlias(m),
		Runtime: runtime, 
	}

	return json.Marshal(aux)
}

func ValidateCourse(v *validator.Validator, course *Course) {
	v.Check(course.Title != "", "title", "must be provided")
	v.Check(len(course.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(course.Year != 0, "year", "must be provided")
	v.Check(course.Year >= 1888, "year", "must be greater than 1888")
	v.Check(course.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(course.Runtime != 0, "runtime", "must be provided") 
	v.Check(course.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(course.Subjects != nil, "subjects", "must be provided")
	v.Check(len(course.Subjects) >= 1, "subjects", "must contain at least 1 genre") 
	v.Check(len(course.Subjects) <= 5, "subjects", "must not contain more than 5 subjects") 
	v.Check(validator.Unique(course.Subjects), "subjects", "must not contain duplicate values")
}

type CourseModel struct { 
	DB *sql.DB
}

func (m CourseModel) Insert(course *Course) error { 
	query := `
		INSERT INTO courses (title, year, runtime, subjects) VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`

	args := []interface{}{course.Title, course.Year, course.Runtime, pq.Array(course.Subjects)}

	return m.DB.QueryRow(query, args...).Scan(&course.ID, &course.CreatedAt, &course.Version)
}

func (m CourseModel) Get(id int64) (*Course, error) { 
	return nil, nil
}

func (m CourseModel) Update(course *Course) error { 
	return nil
}

func (m CourseModel) Delete(id int64) error { 
	return nil
}