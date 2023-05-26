package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}

	
    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(http.StatusCreated, &res)
}
