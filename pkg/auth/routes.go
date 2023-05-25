package auth

import (
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/auth"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient){
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOder)
}

func (svc *ServiceClient)CreateOder(ctx *gin.Context){
	routes.CreateOrder(ctx, svc.Client)
}