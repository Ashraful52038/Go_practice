package models

import "fmt"

// Student struct - Variable দিয়ে Student define
type student struct{
	ID			string
	Name		string
	Class		string
	Age			int
	GPA			float64
	Subjects	[]string // Variable: slice (dynamic array)
	IsActive	bool
	Marks		map [String]int // Variable: map (key-value pair)
}

// CONSTRUCTOR FUNCTION
func NewStudent(id, name string, age int, class string) *Student{
	return &Student{
		ID:       id,
		Name:     name,
		Age:      age,
		Class:    class,
		GPA:      0.0,
		Subjects: []string{"Bangla", "English", "Math"},
		IsActive: true,
		Marks:    make(map[string]int),
	}
}

// METHOD - GPA update
function (s *Student) UpdateGPA(newGPA float64) error{

}