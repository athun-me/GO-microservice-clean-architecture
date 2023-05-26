package product

import (
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/auth"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/config"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSev *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSev)

	svc := &ServcieClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/:id", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServcieClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}
func (svc *ServcieClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
