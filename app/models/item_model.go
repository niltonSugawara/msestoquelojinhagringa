package models

import (
	"github.com/google/uuid"
)

type Item struct{
	ID uuid.UUID `gorm:"primary_key;type:uuid;column:id"`
	DataCriacao string `gorm:"type:date;column:data_criacao"`
	Nome string `gorm:"type:varchar(50);column:nome"`
	Marca string `gorm:"type:varchar(50);column:marca"`
	Valor float32 `gorm:"type:real;column:valor"`
	Descricao string `gorm:"type:text;column:descricao"`

}