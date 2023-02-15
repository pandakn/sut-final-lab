package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `valid:"required"`
	Email      string
	EmployeeID string `valid:"matches(^[JMS]+[0-9]{8}$)"`
}
