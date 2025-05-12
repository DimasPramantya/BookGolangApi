package helper

import (
	"time"

	"github.com/DimasPramantya/goMiniProject/internal/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TimeToString(t *time.Time) *string {
	if(t == nil) {
		return nil
	}
	formattedTime := t.Format("2006-01-02 15:04:05")
	return &formattedTime
}

func WriteError(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(code, dto.ErrorResponse{
		Error: msg,
		Code:  code,
		Data:  data,
	})
}

func WriteResponse(ctx *gin.Context, code int, message string,  data any) {
	ctx.JSON(code, dto.BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

