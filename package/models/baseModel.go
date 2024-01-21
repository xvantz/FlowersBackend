package models

import (

	"github.com/SamuelTissot/sqltime"
)

// Base define the model
type BaseModel struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt int `gorm:"autoCreateTime:false"`
	UpdatedAt int	`gorm:"autoUpdateTime:false"`
	DeletedAt *sqltime.Time `gorm:"type:timestamp"`
}