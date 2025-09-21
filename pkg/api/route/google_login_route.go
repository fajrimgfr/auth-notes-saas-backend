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

func NewGoogleLoginRouter(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, group *gin.RouterGroup) {
	glc := controller.GoogleLoginController{
		Env: env,
	}
	group.GET("auth/google", glc.GoogleLogin)
}

func NewGoogleLoginCallbackRouter(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserCollection)
	glu := usecase.NewGoogleLoginUsecase(ur, timeout)
	glc := controller.GoogleLoginController{
		GoogleLoginUsecase: glu,
		Env:                env,
	}
	group.GET("auth/google/callback", glc.GoogleLoginCallback)
}
