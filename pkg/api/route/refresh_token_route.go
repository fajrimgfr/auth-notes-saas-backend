package route

import (
	"time"

	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/api/controller"
	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/bootstrap"
	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/domain"
	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/repository"
	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRefreshTokenRouter(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserCollection)
	rtu := usecase.NewRefreshTokenUsecase(ur, timeout)
	rtc := controller.RefreshTokenController{
		RefreshTokenUsecase: rtu,
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
