package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fajrimgfr/auth-notes-saas-backend/bootstrap"
	"github.com/fajrimgfr/auth-notes-saas-backend/api/route"
)

func main() {
	env := bootstrap.NewEnv()
	// db := bootstrap.NewPostgresDatabase(env)

	router := gin.Default()

	publicRouter := router.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
	
	router.Run(env.Port)
}