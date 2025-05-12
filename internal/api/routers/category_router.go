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

func initCategoryRouter(rg *gin.RouterGroup, db *sql.DB, validator *validation.Validator) {
	// Repositories
	categoryRepo := pgrepository.NewCategoryPGRepo(db)
	bookRepo := pgrepository.NewBookPGRepo(db)

	// Usecases
	categoryUC := service.NewCategoryService(categoryRepo, bookRepo)

	// Controllers
	categoryCtrl := controller.NewCategoryController(categoryUC, validator)

	// Routes
	rg.POST("/", middleware.JwtMiddleware(), categoryCtrl.CreateCategory)
	rg.GET("/", middleware.JwtMiddleware(), categoryCtrl.GetAllCategory)
	rg.GET("/:id", middleware.JwtMiddleware(), categoryCtrl.GetCategoryByID)
	rg.PUT("/:id", middleware.JwtMiddleware(), categoryCtrl.UpdateCategory)
	rg.DELETE("/:id", categoryCtrl.DeleteCategory)
	rg.GET("/:id/books", middleware.JwtMiddleware(), categoryCtrl.GetAllBooksByCategoryID)
}