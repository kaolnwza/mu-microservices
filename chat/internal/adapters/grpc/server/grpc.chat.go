package grpcSrv

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func NewGrpcServer() (net.Listener, *grpc.Server) {
	addr := ":" + os.Getenv("RPC_PORT")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	// defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	tls := false // change that to true if needed
	if tls {
		// certFile := "ssl/server.crt"
		// keyFile := "ssl/server.pem"
		// certFile := os.Getenv("DEPLOY_SSL_SERVER_CERT_FILE")
		// if certFile == "" {
		// 	certFile = os.Getenv("SSL_SERVER_CERT_FILE")
		// }

		// keyFile := os.Getenv("DEPLOY_SSL_SERVER_KEY_FILE")
		// if keyFile == "" {
		// 	keyFile = os.Getenv("SSL_SERVER_KEY_FILE")
		// }
		// // creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		// creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		// if err != nil {
		// 	log.Fatalf("Failed loading certificates: %v\n", err)
		// }

		// opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	fmt.Println("GRPC Seer Server Started!!")

	return lis, s
}
