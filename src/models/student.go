package models

type Student struct {
	BaseModel
	FirstName    string `gorm:"type:varchar(100);"`
	LastName     string `gorm:"type:varchar(100);"`
	MobileNumber string `gorm:"type:varchar(10);unique_index"`
	Email        string `gorm:"type:varchar(100);unique_index"`
	Group        Group
	GroupID      string `gorm:"type:uuid;"`
	Marks        []Mark
}
