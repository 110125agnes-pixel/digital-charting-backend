package models

import "time"

type DocMaster struct {
	DocCode       string    `gorm:"column:doccode;primaryKey" json:"docCode"`
	LastName      string    `gorm:"column:lastname" json:"lastName"`
	FirstName     string    `gorm:"column:firstname" json:"firstName"`
	MiddleName    string    `gorm:"column:middlename" json:"middleName"`
	Suffix        string    `gorm:"column:suffix" json:"suffix"`
	Sex           string    `gorm:"column:sex" json:"sex"`
	CivilStatus   string    `gorm:"column:civilstatus" json:"civilStatus"`
	Tin           string    `gorm:"column:tin" json:"tin"`
	PhicNo        string    `gorm:"column:phicno" json:"phicNo"`
	PhicExpiry    time.Time `gorm:"column:phicexpiry;type:date" json:"phicExpiry"`
	DOB           time.Time `gorm:"column:dob;type:date" json:"dob"`
	TelNo         string    `gorm:"column:telno" json:"telNo"`
	ClinicTelNo   string    `gorm:"column:clinictelno" json:"clinicTelNo"`
	CellNo        string    `gorm:"column:cellno" json:"cellNo"`
	Address       string    `gorm:"column:address" json:"address"`
	Clinic        string    `gorm:"column:clinic" json:"clinic"`
	ClinicAddress string    `gorm:"column:clinicaddress" json:"clinicAddress"`
	Vat           string    `gorm:"column:vat" json:"vat"`
	Specialty     string    `gorm:"column:specialty" json:"specialty"`
	Specialist    string    `gorm:"column:specialist" json:"specialist"`
	BankAccount   string    `gorm:"column:bankaccount" json:"bankAccount"`
	Active        string    `gorm:"column:active" json:"active"`
	DocType       string    `gorm:"column:doctype" json:"docType"`
	TaxRate       int       `gorm:"column:taxrate" json:"taxRate"`
	User          string    `gorm:"column:user" json:"user"`
	DateEncoded   time.Time `gorm:"column:dateencoded;type:timestamp" json:"dateEncoded"`
	DocTitle      string    `gorm:"column:doctitle" json:"docTitle"`
	GlCode        string    `gorm:"column:glcode" json:"glCode"`
	Used          string    `gorm:"column:used" json:"used"`
	LedgerLink    string    `gorm:"column:ledgerlink" json:"ledgerLink"`
}

func (DocMaster) TableName() string {
	return "docmaster"
}
