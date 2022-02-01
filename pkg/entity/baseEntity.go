package entity

import (
	"gorm.io/gorm"
)

type Entity interface {
	TableName() string
}

type BaseEntity struct {
	ID        uint           `json:"id" gorm:"autoIncrement;primaryKey;not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:datetime"`
}

type Pagination struct {
	Limit  int
	PageNo int
}

func (entity *BaseEntity) SetID(id uint) {
	entity.ID = id
}
