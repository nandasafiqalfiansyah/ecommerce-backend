package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)


func init() {
	app = gin.New()
	app.Use(cors.Default())
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "ok",
		})
	})

	// // Run server
	// app.Run(":8080")
}

// Handler is the entrypoint for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
