package controller

import (
	"net/http"

	"github.com/DimasPramantya/goMiniProject/internal/domain"
	"github.com/DimasPramantya/goMiniProject/internal/dto"
	"github.com/DimasPramantya/goMiniProject/utils/helper"
	"github.com/DimasPramantya/goMiniProject/utils/validation"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUC domain.UserUsecase
	validator *validation.Validator
}

func NewUserController(userUC domain.UserUsecase, validator *validation.Validator) *UserController {
	return &UserController{
		UserUC: userUC,
		validator: validator,
	}
}

// Register godoc
// @Summary     Register
// @Description Register new user
// @Tags        auth
// @Param       request body dto.ReqRegister true "Register Payload"
// @Produce     json
// @Success     201 {object} dto.BaseResponse
// @Failure     400 {object} dto.ErrorResponse
// @Router      /users/register [POST]
func (uc *UserController) Register(ctx *gin.Context) {
	var req dto.ReqRegister
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		return
	}

	err = uc.UserUC.Register(req)
	if err != nil {
		if(err == domain.ErrDataAlreadyExists) {
			helper.WriteError(ctx, http.StatusBadRequest, "Username has registered", nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusCreated, "registration successful", nil)
}

// Login godoc
// @Summary     Login
// @Description User login
// @Tags        auth
// @Param       request body dto.ReqLogin true "Login Payload"
// @Produce     json
// @Success     200 {object} dto.ResLogin
// @Failure     400 {object} dto.ErrorResponse
// @Router      /users/login [POST]
func (uc *UserController) Login(ctx *gin.Context) {
	var req dto.ReqLogin
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		return
	}

	res, err := uc.UserUC.Login(req)
	if err != nil {
		if(err == domain.ErrUnauthorized) {
			helper.WriteError(ctx, http.StatusUnauthorized, "Invalid username or password", nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "login successful", res)
}

// GetUserProfile godoc
// @Summary     Get User Profile By JWT
// @Description Get user profile by JWT
// @Tags        users
// @Produce     json
// @Success     200 {array} dto.ResUser
// @Failure     400 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /users/profile [get]
func (uc *UserController) GetUserProfile(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	username, _ := ctx.Get("username")
	if userID == nil || username == nil {
		helper.WriteError(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	user, err := uc.UserUC.GetByID(int(userID.(float64)))
	if err != nil {
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if user == nil {
		helper.WriteError(ctx, http.StatusNotFound, "User not found", nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "success", user)
}
