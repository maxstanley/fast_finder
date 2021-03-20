package models

import "gorm.io/gorm"

// Shortcut contains information linking a keyword to a link.
type Shortcut struct {
	gorm.Model
	Keyword string `json:"keyword" gorm:"keyword,primary_key"`
	Link    string `json:"link" gorm:"link"`
}
