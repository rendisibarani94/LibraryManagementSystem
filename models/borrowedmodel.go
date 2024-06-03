package models

type BorrowedBook struct {
	ID     int `gorm:"primaryKey"`
	BookId int
	Date   string
	Book   Book `gorm:"foreignKey:BookId"`
}

func (BorrowedBook) TableName() string {
	return "borrowed_books"
}