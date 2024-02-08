package routes

import (
	"context"
	"net/http"
	"strconv"

	productpb "github.com/ErwinSalas/go-grpc-product-svc/proto"
	"github.com/gin-gonic/gin"
)

func FineOne(ctx *gin.Context, c productpb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOne(context.Background(), &productpb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
