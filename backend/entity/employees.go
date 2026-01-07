package entity

import (
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Name         string  `valid:"required~Name is required,stringlength(2|80)~Length name 2-80 character"`
	Salary       float64 `valid:"rang(15000|200000)~Salary must be between 15000 and 200000"`
	EmployeeCode string  `valid:"matches(^[A-Z]{2}-[0-9]{4}$)~EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"`
}
