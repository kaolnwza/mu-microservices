package grpcSrv

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func NewGrpcUserServer() (net.Listener, *grpc.Server) {
	addr := ":" + os.Getenv("RPC_PORT")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	// defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	tls := true // change that to true if needed
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
		// 	return nil, nil
		// }

		// opts = append(opts, grpc.Creds(creds))
	}

	// opts =append(opts, grpc.ChainUnaryInterceptor(Login))
	// opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))

	s := grpc.NewServer(opts...)
	// pb.RegisterProfileServiceServer(s, &GrpcStorageServer{})

	// pb.RegisterUserServiceServer(s, rpc.NewGrpcUserServer())

	// // defer s.Stop()
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v\n", err)
	// }

	fmt.Println("GRPC User Server Started!!")

	return lis, s
}
