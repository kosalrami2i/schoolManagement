package models

type Subject struct {
	BaseModel
	Title           string `gorm:"type:varchar(100);"`
	Marks           []Mark
	SubjectTeachers []SubjectTeacher
}
