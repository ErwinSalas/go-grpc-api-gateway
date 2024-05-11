package product

import (
	"fmt"
	"log"

	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/config"
	productpb "github.com/ErwinSalas/go-grpc-product-svc/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client productpb.ProductServiceClient
}

func InitServiceClient(c *config.Config) productpb.ProductServiceClient {
	tlsCredentials, err := config.LoadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithTransportCredentials(tlsCredentials))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return productpb.NewProductServiceClient(cc)
}
