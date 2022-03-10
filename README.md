# Golang API
## Library Used
- Gin
- GORM

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
--- | ---
Title | string
--- | ---
Description | string
--- | ---
Price | int
--- | ---
Rating | int
--- | ---
Discount | int

# API Testing
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















