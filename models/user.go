package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	InternalId int64 `json:"internal_id" db:"internal_id" gorm:"primaryKey"`
	PublicId  uuid.UUID `json:"public_id" db:"public_id"`
	Name string `json:"name" db:"name"`
	Email string `json:"email" db:"email" gorm:"unique"`
	Password string `json:"password" db:"password" gorm:"column:password`
	Role string	`json:"role" db:"role"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
	DeleteAt gorm.DeletedAt `json:- gorm:"index"`
}