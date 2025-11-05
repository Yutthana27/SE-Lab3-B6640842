package entity

import "gorm.io/gorm"

type Student2 struct {
    gorm.Model
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       int    `json:"age"`
    Email     string `json:"email"`
}