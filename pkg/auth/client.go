package auth

import (
	"fmt"

	"github.com/ErwinSalas/go-grpc-api-gateway/pkg/config"
	authpb "github.com/ErwinSalas/go-grpc-auth-svc/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client authpb.AuthServiceClient
}

func InitServiceClient(c *config.Config) authpb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	fmt.Println(c.AuthSvcUrl)

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return authpb.NewAuthServiceClient(cc)
}
