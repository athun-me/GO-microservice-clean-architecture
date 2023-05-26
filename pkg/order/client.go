package order

import (
	"fmt"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/config"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)

}
