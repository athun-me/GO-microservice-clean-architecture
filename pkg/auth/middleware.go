package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/athunlal/mini-ecommerce-microservice-clean-architecture/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfige struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfige {
	return AuthMiddlewareConfige{svc}
}

func (c *AuthMiddlewareConfige) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("autherization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
