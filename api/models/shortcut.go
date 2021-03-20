package models

import "gorm.io/gorm"

// Shortcut contains information linking a keyword to a link.
type Shortcut struct {
	gorm.Model
	Keyword string `gorm:"keyword,primary_key"`
	Link    string `gorm:"link"`
}
