package main

import (
	"golang-api/book"
	"golang-api/handler"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//main
//handler
//service
//repository
//db (mysql)

func handleFoo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	rw.Header().Set("Access-Control-Allow-Headers", "authentication, content-type")
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	// bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// versioning (api: v1/hello, v1/books, ..)
	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.Run() //default run on port 8080

	// bookRequest := book.BookRequest{
	// 	Title: "iron man",
	// 	Price: 90000,
	// }

	// bookService.Create(bookRequest)

	//inisialisasi input (dummy) CREATE
	// book := book.Book{}
	// book.Title = "Atomic Habits"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 4
	// book.Description = "Buku self development tentang membangun kebiasan baik dan menghilangkan kebiasaan buruk"

	// //kalau variabel sudah ada, assignment tak perlu :=
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("====================\n Error Creating Book Record =====================\n")
	// }

	// var book book.Book

	//First ngambil data pertama masukkin ke &book
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("===================== Error Creating Book Record =====================\n")
	// }

	// book.Title = "Man tiger"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================\n Error Updating Book Record =====================\n")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================\n Error Deleting Book Record =====================\n")
	// }

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// 	fmt.Println("book object: ", book)
	// }

}
