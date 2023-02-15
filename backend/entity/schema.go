package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `valid:"required"`
	Email      string
	EmployeeID string `valid:"matches(^([JMS])+\\d{8}$)~Employee id is valid"`
}
