package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name    string
	Code    string
	Flag    string
	Rtl     bool
	Default bool
}
