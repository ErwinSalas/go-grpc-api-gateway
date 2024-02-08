package routes

import (
	"context"
	"net/http"

	authpb "github.com/ErwinSalas/go-grpc-auth-svc/proto"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c authpb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &authpb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
