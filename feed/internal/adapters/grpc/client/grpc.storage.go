package grpcClient

import (
	"fmt"
	"os"

	"github.com/kaolnwza/muniverse/feed/internal/adapters/grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// func NewProfileServiceClient() pb.ProfileServiceClient {
// 	return pb.NewProfileServiceClient(grpcImageStoreClient())
// }
func NewStorageServiceClient() pb.StorageServiceClient {
	return pb.NewStorageServiceClient(grpcClientConn(os.Getenv("STORER_RPC_PORT"), "stoage", os.Getenv("STORAGE_DNS")))
}

func grpcClientConn(port string, srvName string, dns string) *grpc.ClientConn {
	// conn, err := grpc.Dial("localhost:50051")
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }

	addr := dns + ":" + port

	opts := []grpc.DialOption{}

	tls := false
	if tls {
		// certFile := os.Getenv("DEPLOY_SSL_CLIENT_CERT_FILE")
		// if certFile == "" {
		// 	certFile = os.Getenv("SSL_CLIENT_CERT_FILE")
		// }
		// creds, err := credentials.NewClientTLSFromFile(certFile, "")

		// if err != nil {
		// 	fmt.Println("Error while loading CA trust certificate: ", err)
		// }
		// opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		fmt.Println("Did not connect: ", err)
	}

	fmt.Println("GRPC Cient Connected " + srvName + " Server!")

	// defer conn.Close()

	return conn
}
