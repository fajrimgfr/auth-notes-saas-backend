package route

import (
	"github.com/fajrimgfr/auth-notes-saas-backend/bootstrap"
	"database/sql"
	"time"
	"github.com/gin-gonic/gin"

	"net/http"
)

func NewSignupRouter(env *bootstrap.Env, db *sql.DB, timeout time.Duration, group *gin.RouterGroup) {
	group.GET("/signup", func(c *gin.Context){
		c.IndentedJSON(http.StatusOK, "test")
	}) 
}