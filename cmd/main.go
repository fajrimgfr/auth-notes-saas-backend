package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fajrimgfr/auth-notes-saas-backend/bootstrap"
)

func main() {
	env := bootstrap.NewEnv()

	router := gin.Default()
	
	router.Run("localhost:8080")
}