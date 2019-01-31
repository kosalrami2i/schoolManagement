package models

type Group struct {
	BaseModel
	Name            string `gorm:"type:varchar(100);"`
	Students        []Student
	SubjectTeachers []SubjectTeacher
}
