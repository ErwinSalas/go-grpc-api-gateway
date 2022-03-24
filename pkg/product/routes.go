package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	c := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/product")
	routes.POST("/", c.AddProduct)
	routes.GET("/:id", c.FindOne)
}

func (c *ServiceClient) FindOne(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (c *ServiceClient) AddProduct(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}