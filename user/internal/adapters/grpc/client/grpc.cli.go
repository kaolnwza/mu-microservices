package grpcClient

import (
	"fmt"
	"os"

	"github.com/kaolnwza/muniverse/user/internal/adapters/grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProfileServiceClient() pb.ProfileServiceClient {
	return pb.NewProfileServiceClient(grpcClientConn(os.Getenv("STORER_RPC_PORT"), os.Getenv("STORAGE_DNS")))
}

func NewWalletServiceClient() pb.WalletServiceClient {
	return pb.NewWalletServiceClient(grpcClientConn(os.Getenv("WALLET_RPC_PORT"), os.Getenv("WALLET_DNS")))
}

func grpcClientConn(port string, dns string) *grpc.ClientConn {
	// conn, err := grpc.Dial("localhost:50051")
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }

	opts := []grpc.DialOption{}

	isTls := false
	if isTls {
		// certFile := os.Getenv("DEPLOY_SSL_CLIENT_CERT_FILE")
		// if certFile == "" {
		// 	certFile = os.Getenv("SSL_CLIENT_CERT_FILE")
		// }

		// creds, err := credentials.NewClientTLSFromFile(certFile, "")
		// if err != nil {
		// 	fmt.Println("Error while loading CA trust certificate: ", err)
		// 	return nil
		// }

		// cp := x509.NewCertPool()
		// if !cp.AppendCertsFromPEM([]byte(os.Getenv("SSL_CLIENT_CERT"))) {
		// 	// return nil, fmt.Errorf("credentials: failed to append certificates")
		// 	fmt.Println("credentials: failed to append certificates")
		// }
		// creds := credentials.NewTLS(&tls.Config{ServerName: "", RootCAs: cp})

		// opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	conn, err := grpc.Dial(dns+":"+port, opts...)
	if err != nil {
		fmt.Println("Did not connect: ", err)
		return nil
	}

	fmt.Println("GRPC Cient Connected Storage Server!")

	// defer conn.Close()

	return conn
}
