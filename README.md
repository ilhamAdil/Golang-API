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

There are 4 layers, in this project:
- Models 

Book Entity (Table) using golang struct
- Repository

Created: repository.go
- UseCase

Created: service.go
- Delivery

Created: handler.go




