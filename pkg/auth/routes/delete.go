package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context, c pb.AuthServiceClient){
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil{
		log.Println(err)
	}
	fmt.Println("id", id)
	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
