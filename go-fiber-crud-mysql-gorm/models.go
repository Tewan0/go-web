package main

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Price    float32
	Quantity int32
}
