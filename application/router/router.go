package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"POST"},
		AllowHeaders:  []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	return r
}
