package fb_token_cache

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	return router
}

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("/logs/token_cache.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := setupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
