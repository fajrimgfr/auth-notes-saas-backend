package main

import (
	"fmt"
	"time"

	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/api/route"
	"github.com/fajrimgfr/auth-notes-saas-backend/pkg/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	env := bootstrap.NewEnv()

	db := bootstrap.NewPostgresDatabase(env)
	defer db.Close()

	bootstrap.NewGoogleOAuthConfig(env)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	route.Setup(env, db, timeout, router)

	router.Run(fmt.Sprintf(":%s", env.Port))
}
