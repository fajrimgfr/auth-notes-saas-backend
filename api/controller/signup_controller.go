package controller

import (
	"database/sql"
	"net/http"

	"time"

	"github.com/fajrimgfr/auth-notes-saas-backend/bootstrap"
	"github.com/fajrimgfr/auth-notes-saas-backend/domain"
	"github.com/fajrimgfr/auth-notes-saas-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if _, err := sc.SignupUsecase.GetUserByEmail(c, request.Email); err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(hasedPassword)

	user := domain.User{
		ID:        uuid.New().String(),
		Email:     request.Email,
		Password:  request.Password,
		Name:      sql.NullString{String: request.Name, Valid: true},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := sc.SignupUsecase.Create(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := util.CreateAccessToken(&user, sc.Env.AccessTokenExpiryHour, sc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := util.CreateRefreshToken(&user, sc.Env.RefreshTokenExpiryHour, sc.Env.RefreshTokenExpirySecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
