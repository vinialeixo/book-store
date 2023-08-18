package entities

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model

	PublisherName string
	Authors       string
}
