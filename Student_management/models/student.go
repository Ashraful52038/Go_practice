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
	if newGPA<0 || newGPA>5.0{
		return fmt.Errorf("Invalid GPA: must be between 0 and 5.")
	}
	s.GPA=newGPA;
	return nil
}

// METHOD - Add new subject
func (s *Student) AddSubject(sunject string ){
	for i, existingSubject := range s.Subjects{
		if existingSubject == subject {
			return
		}
	}
	s.Subjects = append(s.Subjects, subject)
}

// METHOD - Mark add
func (s *Student) AddMark(subject string, mark int) {
	// CONDITION: Mark validation
	if mark >= 0 && mark <= 100 {
		s.Marks[subject] = mark
		s.calculateGPAFromMarks()
	}
}

// PRIVATE FUNCTION - Marks to GPA calculate
func (s *Student) calculateGPAFromMarks() {
	if len(s.Marks) == 0 {
		return
	}
	
	// LOOP: Calculate average marks
	total := 0
	for _, mark := range s.Marks {
		total += mark
	}

	average := float64(total) / float64(len(s.Marks))
	
	// Convert marks to GPA (CONDITION: Grading system)
	if average >= 80 {
		s.GPA = 5.0
	} else if average >= 70 {
		s.GPA = 4.0
	} else if average >= 60 {
		s.GPA = 3.5
	} else if average >= 50 {
		s.GPA = 3.0
	} else if average >= 40 {
		s.GPA = 2.0
	} else {
		s.GPA = 0.0
	}
}

// METHOD - Grade
func (s *Student) GetGrade() string {
	// CONDITION: GPA to grade
	switch {
	case s.GPA >= 4.5:
		return "A+"
	case s.GPA >= 4.0:
		return "A"
	case s.GPA >= 3.5:
		return "B+"
	case s.GPA >= 3.0:
		return "B"
	case s.GPA >= 2.0:
		return "C"
	default:
		return "F"
	}
}

// METHOD - Student information print
func (s *Student) PrintInfo() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("Student ID: %s\n", s.ID)
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d | Class: %s\n", s.Age, s.Class)
	fmt.Printf("GPA: %.2f | Grade: %s\n", s.GPA, s.GetGrade())
	fmt.Printf("Active: %v\n", s.IsActive)
	
	// LOOP: Print subjects
	fmt.Print("Subjects: ")
	for i, subject := range s.Subjects {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(subject)
	}
	fmt.Println()
	
	// LOOP: Print marks
	if len(s.Marks) > 0 {
		fmt.Println("Marks:")
		for subject, mark := range s.Marks {
			fmt.Printf("  %s: %d\n", subject, mark)
		}
	}
	fmt.Println(strings.Repeat("=", 50))
}