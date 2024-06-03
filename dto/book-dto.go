package dto

type Books struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Writer      string `json:"writer"`
	Page        int    `json:"page"`
	Publication string `json:"publication"`
}

type BorrowedBook struct {
	ID     int    `json:"id"`
	BookId int    `json:"book_id"`
	Date   string `json:"date"`
}

type BorrowedBookResponse struct {
	ID          int    `json:"id"`
	BookID      int    `json:"book_id"`
	BookName    string `json:"book_name"`
	Writer      string `json:"writer"`
	Page        int    `json:"page"`
	Date        string `json:"date"`
	Publication string `json:"publication"`
}
