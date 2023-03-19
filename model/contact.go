package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerID     uint
	TargerID    uint
	Type        int
	Description string
}
