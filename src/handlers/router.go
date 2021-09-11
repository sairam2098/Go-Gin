package handlers

import (
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (*Router) GetRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (*Router) NewRouter(router *gin.Engine) *gin.Engine {
	router.Use(CORSMiddleware())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	router.GET("/items", GetItems)
	router.GET("/items/by-category", GetItemsByCategory)
	return router
}
