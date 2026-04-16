package models

import "time"

type User struct {
	Doccode    uint      `gorm:"column:doccode;primaryKey;autoIncrement" json:"doccode"`
	Email      string    `gorm:"column:email;unique;not null" json:"email"`
	Password   string    `gorm:"column:password" json:"password"`
	FirstName  string    `gorm:"column:first_name" json:"first_name"`
	LastName   string    `gorm:"column:last_name" json:"last_name"`
	MiddleName string    `gorm:"column:middle_name" json:"middle_name"`
	CellNo     string    `gorm:"column:cellno" json:"cellno"`
	DOB        time.Time `gorm:"column:dob;type:date" json:"dob"`
}
