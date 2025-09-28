# Go-Project with Gin Framework

A fast and clean **REST API** built with **Go (Golang)** using the **Gin framework**, implementing **JWT authentication**, **middlewares**, and following **clean code architecture** principles.

---

## Features

- **Clean Architecture**: Separation of concerns with structured folders (`models`, `routes`, `middlewares`, `utils`, etc.)  
- **JWT Authentication**: Secure login with JWT tokens for protected routes  
- **Middlewares**: Custom middlewares for authentication, logging, and request validation  
- **Fast and Lightweight**: Leveraging Gin for high-performance routing and HTTP handling  
- **Database Integration**: Using SQLite/Postgres/MySQL for persistence (adjust `db` package)  
- **API Testing**: Fully tested endpoints with proper JSON requests/responses  
- **Error Handling**: Clear and consistent error messages for API clients  

---

## Folder Structure

go-project/
│
├── api-test/ # API testing scripts (Postman collections, etc.)
├── db/ # Database setup and migrations
├── middlewares/ # JWT and custom middlewares
├── models/ # Database models and validation
├── routes/ # API routes and handlers
├── utils/ # Utility functions (e.g., password hashing)
├── main.go # Entry point
├── go.mod # Go modules
└── go.sum


Install dependencies:

go mod tidy


Run the server:

go run main.go
