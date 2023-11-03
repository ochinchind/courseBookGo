package data

import ( 
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found") 
)

type Models struct {
	Courses CourseModel 
}

func NewModels(db *sql.DB) Models {
	return Models {
		Courses: CourseModel{DB: db},
	} 
}