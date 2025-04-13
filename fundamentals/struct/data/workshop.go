package data

import "time"

type Workshop struct {
	Course
	Instructor
	Date time.Time
}

func NewWorkShop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	return w
}
