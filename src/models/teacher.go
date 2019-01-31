package models

type Teacher struct {
	BaseModel
	FirstName       string `gorm:"type:varchar(100);"`
	LastName        string `gorm:"type:varchar(100);"`
	SubjectTeachers []SubjectTeacher
}
