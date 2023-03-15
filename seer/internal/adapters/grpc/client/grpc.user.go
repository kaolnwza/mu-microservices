package grpcClient

import (
	"fmt"
	"os"

	"github.com/kaolnwza/muniverse/seer/internal/adapters/grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient() pb.UserServiceClient {
	return pb.NewUserServiceClient(grpcClientConn(os.Getenv("USER_RPC_PORT"), os.Getenv("USER_DNS")))
}
func NewStorageServiceClient() pb.ProfileServiceClient {
	return pb.NewProfileServiceClient(grpcClientConn(os.Getenv("STORER_RPC_PORT"), os.Getenv("STORAGE_DNS")))
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

	fmt.Println("GRPC Cient Connected User Server!")

	// defer conn.Close()

	return conn
}
