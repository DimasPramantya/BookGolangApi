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

func InitUserRouter(rg *gin.RouterGroup, db *sql.DB, validator *validation.Validator) {
	// Repositories
	userRepo := pgrepository.NewUserPgRepository(db)

	// Usecases
	userUC := service.NewUserService(userRepo)

	// Controllers
	userCtrl := controller.NewUserController(userUC, validator)

	// Routes
	rg.POST("/register", userCtrl.Register)
	rg.POST("/login", userCtrl.Login)
	rg.GET("/profile", middleware.JwtMiddleware(), userCtrl.GetUserProfile)
}