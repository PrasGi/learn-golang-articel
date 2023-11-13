package models

import "gorm.io/gorm"

type Articel struct {
	gorm.Model
	Content string
}
