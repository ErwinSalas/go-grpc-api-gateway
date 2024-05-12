package main

import (
	"log"
	"net/http"

	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/auth"
	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/config"
	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/order"
	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/product"
	"github.com/ErwinSalas/go-grpc-tls/pkg/gogrpctls"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	c.CertM = gogrpctls.NewCertManager()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
