package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string
	OwnerID     uint
	Description string
	Account     int
	Members     []*User `gorm:"many2many:user_groups;"`
}
