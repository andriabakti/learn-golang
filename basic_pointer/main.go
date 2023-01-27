package main

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// method dengan pointer receiver
func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
	// println("student name >", student.Name)
}

func mainPointer() {
	student := Student{1, "Argos", 3.27}
	println(student.Name)
	student.graduate()
	println(student.Name)
}
