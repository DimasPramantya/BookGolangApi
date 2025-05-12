package routers

import (
	"database/sql"

	"github.com/DimasPramantya/goMiniProject/utils/validation"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, db *sql.DB) {
	validator := validation.NewValidator()
	api := r.Group("/api")

	userRoute := api.Group("/users")
	InitUserRouter(userRoute, db, validator) 

	categoryRoute := api.Group("/categories")
	initCategoryRouter(categoryRoute, db, validator) 

	bookRoute := api.Group("/books")
	InitBookRouter(bookRoute, db, validator)
}