# BookGolangApi

This is a RESTful API built in Go for managing books and categories with JWT authentication and authorization. It follows a **Clean Architecture** approach inspired by [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

## 🧱 Project Architecture

This project implements a layered structure:

- `cmd/` – entry point of the application
- `internal/` – contains core application logic, divided into:
  - `domain/` – interfaces and domain models
  - `usecase/` – business logic
  - `repository/` – data access
  - `api/controller/` – HTTP handlers using Gin
- `utils/` – helper packages (e.g., validation, response writing)
- `docs/` – auto-generated Swagger documentation

## 📄 API Documentation

Access the interactive API documentation via Swagger:

👉 [View Swagger Docs](https://bookgolangapi-production.up.railway.app/api/swagger/index.html#/)

This documentation includes all available endpoints for:
- Authentication (register/login)
- Book and Category CRUD operations
- Authorization via Bearer Token

To authorize requests in Swagger:
1. Click on the "Authorize" button.
2. Enter your token in the format: `Bearer <your_token>`.

## 🚀 Getting Started

To run the project locally:

```bash
git clone https://github.com/your-username/BookGolangApi.git
cd BookGolangApi
go run cmd/goMiniProject/main.go
