package auth

import (
	"log"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/auth/pb"
	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect :", err)
	}
	return pb.NewAuthServiceClient(cc)
}
