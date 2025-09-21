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

func NewLoginRouter(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserCollection)
	lu := usecase.NewSignupUsecase(ur, timeout)
	lc := controller.LoginController{
		LoginUsecase: lu,
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
