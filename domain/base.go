package domain

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID string `json:"id" gorm:"uuid;primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"datetime"`
}

func (base Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())

	return nil
}

func (base Base) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())

	return nil
}