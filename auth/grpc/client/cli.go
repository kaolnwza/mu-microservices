package grpcClient

import (
	"fmt"
	"os"

	pb "github.com/kaolnwza/muniverse/auth/grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient() pb.UserServiceClient {
	return pb.NewUserServiceClient(grpcClientConn(os.Getenv("USER_RPC_PORT"), os.Getenv("USER_DNS")))
}

func NewSeerServiceClient() pb.SeerServiceClient {
	return pb.NewSeerServiceClient(grpcClientConn(os.Getenv("SEER_RPC_PORT"), os.Getenv("SEER_DNS")))
}

func grpcClientConn(port string, dns string) *grpc.ClientConn {
	opts := []grpc.DialOption{}
	addr := dns + ":" + port

	tls := false

	if tls {
		// 	certFile := os.Getenv("DEPLOY_SSL_CLIENT_CERT_FILE")
		// 	if certFile == "" {
		// 		certFile = os.Getenv("SSL_CLIENT_CERT_FILE")
		// 	}
		// 	creds, err := credentials.NewClientTLSFromFile(certFile, "")

		// 	if err != nil {
		// 		fmt.Println("Error while loading CA trust certificate: ", err)
		// 	}
		// 	opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		fmt.Println("Did not connect: ", err)
	}

	fmt.Println("GRPC Cient Connected Server!")

	// defer conn.Close()

	return conn
}
