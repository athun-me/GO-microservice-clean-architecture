package main

import (
	"log"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/auth"

)


func main(){
	c, err := config.LoadConfig()

	if err != nil{
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.
	
}