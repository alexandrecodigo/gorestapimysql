package webapi

import (
	"github.com/alexandrecodigo/gorestapimysql/internal/app/grocery/webhandler/shoppinglisthandler"
	"github.com/gin-gonic/gin"
)

// Route routes
func Route(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1ShoppingList := r.Group("/v1/shopping-list")
	{
		v1ShoppingList.POST("/", shoppinglisthandler.CreateHandler)
		v1ShoppingList.GET("/:id", shoppinglisthandler.ShowHandler)
	}

	return r
}
