package models

import "gorm.io/gorm"

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
	RoleVendor   UserRole = "vendor"
)

type User struct {
	gorm.Model
	ID    uint     `json:"id" gorm:"primaryKey"`
	Name  string   `json:"name"`
	Email string   `json:"email" gorm:"unique"`
	Role  UserRole `json:"role" gorm:"type:user_role"`
}
