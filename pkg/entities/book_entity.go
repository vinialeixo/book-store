package entities

import (
	"time"

	"gorm.io/gorm"
)

type Status int

const (
	NOT_AVALIABLE Status = 1
	AVALIABLE     Status = 2
)

type Book struct {
	gorm.Model

	BookName         string    `json:"BookName" gorm:"type:varchar(100)"`
	PublisherRefer   uint      `json:"PublisherRefer"`
	Publisher        Publisher `gorm:"foreignKey:PublisherRefer"`
	BriefDescription string    `json:"BriefDescription" gorm:"type:varchar(100)"`
	Description      string    `json:"Description" gorm:"type:varchar(500)"`
	Author           string    `json:"Author" gorm:"type:varchar(50)"`
	Dimenssions      string    `json:"Dimenssions" gorm:"type:varchar(30)"`
	CodeISBM         string    `json:"CodeISBM" gorm:"type:varchar(50)"`
	NumberOfPages    uint      `json:"NumberOfPages"`
	PublicationDate  time.Time `json:"PublicationDate"`
	BookImage        string    `json:"BookImage"`
	BookStock        uint      `json:"BookStock"`
	Price            float64   `json:"Price"`
	Status           Status    `json:"Status"`
}
