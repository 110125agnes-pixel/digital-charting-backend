package models

import "time"

type User struct {
	DocCode    uint      `gorm:"column:doc_code;primaryKey;autoIncrement" json:"docCode"`
	Email      string    `gorm:"column:email;unique;not null" json:"email"`
	Password   string    `gorm:"column:password" json:"password"`
	FirstName  string    `gorm:"column:first_name" json:"firstName"`
	LastName   string    `gorm:"column:last_name" json:"lastName"`
	MiddleName string    `gorm:"column:middle_name" json:"middleName"`
	CellNo     string    `gorm:"column:cell_no" json:"cellNo"`
	DOB        time.Time `gorm:"column:dob;type:date" json:"dob"`
}
