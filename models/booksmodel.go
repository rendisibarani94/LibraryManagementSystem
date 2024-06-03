package models

type Book struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Writer      string
	Page        int
	Publication string
}