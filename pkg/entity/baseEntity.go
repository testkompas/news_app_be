package entity

import (
	"gorm.io/gorm"
)

type Entity interface {
	TableName() string
}

type BaseEntity struct {
	ID        int            `json:"id" gorm:"type:int(10);autoIncrement;primaryKey;not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:datetime"`
}

type Pagination struct {
	Limit  int
	PageNo int
}

func (entity *BaseEntity) SetID(id int) {
	entity.ID = id
}
