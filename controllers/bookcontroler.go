package controllers

import (
	"encoding/json"
	"first-jwt/configs"
	"first-jwt/dto"
	"first-jwt/helpers"
	"first-jwt/models"
	"net/http"

	"github.com/gorilla/mux"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
var book dto.Books

if err := json.NewDecoder(r.Body).Decode(&book); err != nil{ // Decode function returns error (err)
	helpers.Response(w, 500, err.Error(), nil)
	return
}
defer r.Body.Close() // will execute at the end of the function

buku := models.Book {
	Name: book.Name,
	Writer: book.Writer,
	Page: book.Page,
	Publication: book.Publication,
}

// Create function do create data to database, and Error to catch the error
if err := configs.DB.Create(&buku).Error; err != nil{
	helpers.Response(w, 500, err.Error(), nil)
	return
}
helpers.Response(w, 201, "Book added Successfully", nil)

}

func ViewALlBooks (w http.ResponseWriter, r *http.Request) {
	var books []models.Book

    // Find all records in the 'books' table and store them in the 'books' variable
    if err := configs.DB.Find(&books).Error; err != nil {
        helpers.Response(w, 500, err.Error(), nil)
        return
    }

    // If books are found, return them in the response
    helpers.Response(w, 200, "Books retrieved successfully", books)
}

func ViewBooksById (w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// Extract the 'id' parameter from the URL
	vars := mux.Vars(r)
	id := vars["id"]


    // Find all records in the 'book' table and store them in the 'book' variable
	if err := configs.DB.First(&book, id).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

    // If book are found, return them in the response
    helpers.Response(w, 200, "Book retrieved successfully", book)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	// Extract the 'id' parameter from the URL path
	vars := mux.Vars(r)
	id := vars["id"]

	// Decode the request body to get the updated book information
	var updatedBook models.Book
	
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		helpers.Response(w, 400, "Invalid request payload", nil)
		return
	}
	defer r.Body.Close()

	var existingBook models.Book

	// Find the existing record in the 'books' table with the specified 'id'
	if err := configs.DB.First(&existingBook, id).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// Update the existing book with the new information
	existingBook.Name = updatedBook.Name
	existingBook.Writer = updatedBook.Writer
	existingBook.Page = updatedBook.Page
	existingBook.Publication = updatedBook.Publication

	// Save the updated book back to the database
	if err := configs.DB.Save(&existingBook).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// If update is successful, return the updated book in the response
	helpers.Response(w, 200, "Book updated successfully", existingBook)
}

func BorrowBook (w http.ResponseWriter, r *http.Request) {
	var book dto.BorrowedBook

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil{ // Decode function returns error (err)
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	buku := models.BorrowedBook {
		ID: book.ID,
		BookId: book.BookId,
		Date: book.Date,
	}

	if err := configs.DB.Create(&buku).Error; err != nil{
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	
helpers.Response(w, 201, "Book Borrowed Successfully", nil)
}



func GetAllBorrowedBook(w http.ResponseWriter, r *http.Request) {
	var borrowedBooks []models.BorrowedBook

	if err := configs.DB.Preload("Book").Find(&borrowedBooks).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Create a slice to hold the custom response
	var response []dto.BorrowedBookResponse
	// Convert each BorrowedBook to the custom response format
	for _, borrowedBook := range borrowedBooks {
		bookResponse := dto.BorrowedBookResponse{
			ID:          borrowedBook.ID,
			BookID:      borrowedBook.Book.ID,
			BookName:    borrowedBook.Book.Name,
			Writer:      borrowedBook.Book.Writer,
			Page:        borrowedBook.Book.Page,
			Date:        borrowedBook.Date,
			Publication: borrowedBook.Book.Publication,
		}

		response = append(response, bookResponse)
	}

	helpers.Response(w, 200, "Borrowed books retrieved successfully", response)
}