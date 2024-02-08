package order

import (
	"fmt"

	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/config"
	orderpb "github.com/ErwinSalas/go-grpc-order-svc/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client orderpb.OrderServiceClient
}

func InitServiceClient(c *config.Config) orderpb.OrderServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return orderpb.NewOrderServiceClient(cc)
}
