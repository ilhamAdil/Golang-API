package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"golang-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	itemBook, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := converToBookResponse(itemBook)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, itemBook := range books {
		bookResponse := converToBookResponse(itemBook)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get succeed",
		"data":    booksResponse,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	//post title, price
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	//jika error tidak kosong
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s ", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post succeed",
		"data":    converToBookResponse(book),
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	//post title, price
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	//jika error tidak kosong
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s ", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	//ambil id utk diupdate
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update succeed",
		"data":    converToBookResponse(book),
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	itemBook, err := h.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := converToBookResponse(itemBook)

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted succeed",
		"data":    bookResponse,
	})

}

func converToBookResponse(itemBook book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          itemBook.ID,
		Title:       itemBook.Title,
		Price:       itemBook.Price,
		Description: itemBook.Description,
		Rating:      itemBook.Rating,
		Discount:    itemBook.Discount,
	}
}
