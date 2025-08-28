package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/fajrimgfr/auth-notes-saas-backend/bootstrap"
	"github.com/fajrimgfr/auth-notes-saas-backend/api/route"
)

func main() {
	env := bootstrap.NewEnv()

	db := bootstrap.NewPostgresDatabase(env)
	defer db.Close()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	route.Setup(env, db, timeout, router)
	
	router.Run(fmt.Sprintf(":%s", env.Port))
}