package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Winner struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Season      string `gorm:"type:varchar(50);not null" json:"season"`
	Game        string `gorm:"type:varchar(100);not null" json:"game"`
	Position    string `gorm:"type:varchar(10);not null" json:"position"`
	TeamMember1 string `gorm:"type:varchar(100);not null" json:"teamMember1"`
	TeamMember2 string `gorm:"type:varchar(100);not null" json:"teamMember2"`
}

// BeforeCreate hook to generate UUID manually if not provided
func (w *Winner) BeforeCreate(tx *gorm.DB) (err error) {
	if w.ID == "" {
		w.ID = uuid.NewString() // Generates a valid UUID string
	}
	return
}
