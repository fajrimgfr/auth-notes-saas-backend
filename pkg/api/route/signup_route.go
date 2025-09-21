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

func NewSignupRouter(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserCollection)
	su := usecase.NewSignupUsecase(ur, timeout)
	sc := controller.SignupController{
		SignupUsecase: su,
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
