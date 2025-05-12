package routers

import (
	"database/sql"

	"github.com/DimasPramantya/goMiniProject/internal/api/controller"
	"github.com/DimasPramantya/goMiniProject/internal/api/middleware"
	pgrepository "github.com/DimasPramantya/goMiniProject/internal/repository/pgRepository"
	"github.com/DimasPramantya/goMiniProject/internal/service"
	"github.com/DimasPramantya/goMiniProject/utils/validation"
	"github.com/gin-gonic/gin"
)

func InitBookRouter(rg *gin.RouterGroup, db *sql.DB, validator *validation.Validator) {
	// Repositories
	bookRepo := pgrepository.NewBookPGRepo(db)
	categoryRepo := pgrepository.NewCategoryPGRepo(db)

	// Usecases
	bookUC := service.NewBookService(bookRepo, categoryRepo)

	// Controllers
	bookCtrl := controller.NewBookController(bookUC, validator)

	// Routes
	rg.POST("/", middleware.JwtMiddleware(), bookCtrl.CreateBook)
	rg.GET("/", middleware.JwtMiddleware(), bookCtrl.GetAllBooks)
	rg.GET("/:id", middleware.JwtMiddleware(), bookCtrl.GetBookByID)
	rg.DELETE("/:id", middleware.JwtMiddleware(), bookCtrl.DeleteBook)
}