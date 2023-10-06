package data

import (
	"encoding/json"
	"fmt"
	"time"
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