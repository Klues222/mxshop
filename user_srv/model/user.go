package model

import "time"

type BaseModel struct {
	ID        int32     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:""`
}
