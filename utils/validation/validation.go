package validation

import (
	"fmt"
	"net/http"

	"github.com/DimasPramantya/goMiniProject/internal/domain"
	"github.com/DimasPramantya/goMiniProject/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) ValidateRequest(ctx *gin.Context, req any) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Invalid request", "Invalid Content-Type, expected application/json")
		return domain.ErrBadRequest
	}

	if errs := v.Validate(req); errs != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Bad Request", errs)
		return domain.ErrBadRequest
	}

	return nil
}

func (v *Validator) Validate(i any) []string {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}

	var data []string
	for _, e := range err.(validator.ValidationErrors) {
		data = append(data, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
	}
	return data
}
