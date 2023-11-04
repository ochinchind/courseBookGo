package data

import ( 
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found") 
	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	Courses CourseModel 
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models {
		Courses: CourseModel{DB: db},
		Users: UserModel{DB: db}, 
	} 
}