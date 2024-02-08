package routes

import (
	"context"
	"fmt"
	"net/http"

	authpb "github.com/ErwinSalas/go-grpc-auth-svc/proto"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c authpb.AuthServiceClient) {
	b := RegisterRequestBody{}
	fmt.Println(b)
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &authpb.RegisterRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
