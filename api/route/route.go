package route

import (
	"time"

	"github.com/fajrimgfr/auth-notes-saas-backend/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, router *gin.Engine) {
	publicRouter := router.Group("")
	NewSignupRouter(env, db, timeout, publicRouter)
	NewLoginRouter(env, db, timeout, publicRouter)
}
