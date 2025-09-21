package route

import (
	"time"

	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/api/middleware"
	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, router *gin.Engine) {
	publicRouter := router.Group("")
	NewSignupRouter(env, db, timeout, publicRouter)
	NewLoginRouter(env, db, timeout, publicRouter)
	NewGoogleLoginRouter(env, db, timeout, publicRouter)
	NewGoogleLoginCallbackRouter(env, db, timeout, publicRouter)
	NewRefreshTokenRouter(env, db, timeout, publicRouter)
	privateRouter := router.Group("")
	privateRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
}
