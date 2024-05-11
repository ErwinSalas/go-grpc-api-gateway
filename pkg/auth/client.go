package auth

import (
	"fmt"
	"log"

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

	tlsCredentials, err := config.LoadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(tlsCredentials))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return authpb.NewAuthServiceClient(cc)
}
