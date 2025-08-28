package route

import (
	"github.com/fajrimgfr/auth-notes-saas-backend/bootstrap"
	"database/sql"
	"time"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, db *sql.DB, timeout time.Duration, router *gin.Engine) {
	publicRouter := router.Group("")
	NewSignupRouter(env, db, timeout, publicRouter)
}