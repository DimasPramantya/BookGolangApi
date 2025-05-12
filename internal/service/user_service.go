package service

import (
	"fmt"

	"github.com/DimasPramantya/goMiniProject/internal/api/middleware"
	"github.com/DimasPramantya/goMiniProject/internal/domain"
	"github.com/DimasPramantya/goMiniProject/internal/dto"
	"github.com/DimasPramantya/goMiniProject/utils/helper"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetByID(id int) (*dto.ResUser, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	if user == nil {
		return nil, domain.ErrNotFound
	}
	return &dto.ResUser{
		ID:       user.ID,
		Username: user.Username,
		CreatedAt: *helper.TimeToString(&user.CreatedAt),
		CreatedBy: user.CreatedBy,
		ModifiedAt: helper.TimeToString(user.ModifiedAt),
		ModifiedBy: user.ModifiedBy,
	}, nil
}

func (u *UserService) Login(req dto.ReqLogin) (*dto.ResLogin, error) {
	user, err := u.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}
	if user == nil {
		return nil, domain.ErrUnauthorized
	}
	if(!helper.CheckPasswordHash(req.Password, user.Password)) {
		return nil, domain.ErrUnauthorized
	}
	token, err := middleware.GenerateJwtToken(user.Username, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	return &dto.ResLogin{
		Token: token,
	}, nil
}

func (us *UserService) Register(req dto.ReqRegister) error {
	duplicateUser, err := us.userRepo.FindByUsername(req.Username)
	if err != nil {
		return fmt.Errorf("%s: %w", domain.FailedRegisterMessage, err)
	}
	if duplicateUser != nil {
		return domain.ErrDataAlreadyExists
	}
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", domain.FailedRegisterMessage, err)
	}
	reqUser := domain.User{
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedBy: req.Username,
	}
	_, err = us.userRepo.Create(reqUser)
	if err != nil {
		return fmt.Errorf("%s: %w", domain.FailedRegisterMessage, err)
	}
	return nil
}

func (u *UserService) GetProfile(id int) (*dto.ResUser, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}
	return &dto.ResUser{
		ID:       user.ID,
		Username: user.Username,
		CreatedAt: *helper.TimeToString(&user.CreatedAt),
		CreatedBy: user.CreatedBy,
		ModifiedAt: helper.TimeToString(user.ModifiedAt),
		ModifiedBy: user.ModifiedBy,
	}, nil
}
