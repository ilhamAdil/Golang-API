# Golang API
## Library Used
- Gin
- GORM
- CORS

## Tools Used
- MySQL
- Postman
- GO Extension for Visual Studio Code

# Architecture
![lbexLhWvRfpexSV0lSIWczkHd5KdszeDy9a3](https://user-images.githubusercontent.com/66354919/157606557-1131b3c9-3816-4b1b-8985-fa8fad886e60.png)

![YIABVRTHRz58ZiT6W-emBkfNIQUHBelp8t6U](https://user-images.githubusercontent.com/66354919/157606594-8aeafbf1-8d94-4685-9b91-f83cfffea213.png)

An important goal of clean architecture is to provide developers with a way to organize code in such a way that it encapsulates the business logic but keeps it separate from the delivery mechanism.

### There are 4 layers in this project:
- Models</br>Created a Book Entity (Table) using go struct
- Repository</br>Created repository.go
- UseCase</br>Created service.go
- Delivery</br>Created handler.go

# Entity 
### Book Table 
Attributes | data type
--- | ---
ID | int: auto increment
Title | string
Description | string
Price | int
Rating | int
Discount | int

# API Testing
Endpoints:
```
v1 := router.Group("/v1")

v1.GET("/books", bookHandler.GetBooksHandler)
v1.GET("/books/:id", bookHandler.GetBookHandler)
v1.POST("/books", bookHandler.PostBookHandler)
v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

router.Run() //default run on port 8080
```
This API Test is using Postman
### GET BOOK
![resGET](https://user-images.githubusercontent.com/66354919/157624685-fbfea1d8-df3f-4565-80ae-06df11faa8d9.gif)

### GET BOOK ID
![resGETID](https://user-images.githubusercontent.com/66354919/157622478-c030ce37-64c3-4bff-91c2-5386aefada22.gif)

### POST BOOK
![resPOST](https://user-images.githubusercontent.com/66354919/157624808-a036e9d4-d03a-49b9-9959-f5faa4220d2a.gif)

### UPDATE BOOK
![resUPDATE](https://user-images.githubusercontent.com/66354919/157624940-3bc7c0b6-b6c8-412d-8053-5fa1b8546665.gif)

### DELETE BOOK
![resDELETE](https://user-images.githubusercontent.com/66354919/157625010-3eeb262e-c588-4038-95bd-f433671580f0.gif)

# API Consume
Created a frontend API Consumer using React.js on diffferent repository (Available on my github/ilhamAdil)
## Library used for API consume
- Axios

## Trouble
Cross-Origin Resource Sharing (CORS) errors occur when a server doesn't return the HTTP headers required by the CORS standard. 

To resolve a CORS error from an API Gateway REST API or HTTP API, reconfigure the API to meet the CORS standard. So, in Golang file main.go, added:
```
func handleFoo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	rw.Header().Set("Access-Control-Allow-Headers", "authentication, content-type")
}
```
localhost:3000 is where the API will be consumed by React.js frontend and Axios.

## Frontend DEMO
### Posting New Book
![APIpost](https://user-images.githubusercontent.com/66354919/157632010-af3356a2-56d7-4b52-9d2c-893f5f7e46c4.gif)

### Updating Book
![APIupdate](https://user-images.githubusercontent.com/66354919/157633922-b035ccb9-5399-4313-bd2a-08c6901de7ef.gif)

### Deleting Book
![APIdelete](https://user-images.githubusercontent.com/66354919/157634109-eb78dd6b-8d45-477a-8791-240444de7dc5.gif)

# Resources
- https://github.com/gin-gonic/gin
- https://github.com/go-gorm/gorm
- https://github.com/gin-contrib/cors





















