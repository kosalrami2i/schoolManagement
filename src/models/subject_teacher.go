package models

type SubjectTeacher struct {
	Subject   Subject
	Student   Student
	Group     Group
	SubjectID string `gorm:"type:uuid;"`
	TeacherID string `gorm:"type:uuid;"`
	GroupID   string `gorm:"type:uuid;"`
}
