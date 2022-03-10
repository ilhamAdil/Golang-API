package book

import "time"

//kalau struct singular (book) maka tabel plural (books)
type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
