package models

type Mark struct {
	BaseModel
	Subject   Subject
	Student   Student
	StudentID string `gorm:"type:uuid;"`
	SubjectID string `gorm:"type:uuid;"`
	Score     int8   `gorm:"type:integer;"`
}
