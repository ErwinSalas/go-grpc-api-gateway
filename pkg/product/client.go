package product

import (
	"fmt"

	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/config"
	productpb "github.com/ErwinSalas/go-grpc-product-svc/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client productpb.ProductServiceClient
}

func InitServiceClient(c *config.Config) productpb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return productpb.NewProductServiceClient(cc)
}
