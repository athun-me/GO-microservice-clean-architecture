package api

import (
	"athun/pkg/api/handler"
	"athun/pkg/pb"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	engine := gin.New()

	go NewGrpcServer(userHandler, "8889")

	engine.Use(gin.Logger())

	return &ServerHttp{
		Engine: engine,
	}
}

func NewGrpcServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to the GRPC Port", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Could not serve the GRPC Server: ", err)
	}
}

func (ser *ServerHttp) Start() {
	ser.Engine.Run(":7777")
}
