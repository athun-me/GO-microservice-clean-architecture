package product

import (
	"fmt"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/config"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/product/pb"
	"google.golang.org/grpc"
)


type ServcieClient struct{
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config)pb.ProductServiceClient{
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
    if err != nil {
        fmt.Println("Could not connect:", err)
    }
	return pb.NewProductServiceClient(cc)
}