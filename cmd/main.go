package main

import (
	"fmt"
	"log"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/auth"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/config"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/order"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	fmt.Println()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
