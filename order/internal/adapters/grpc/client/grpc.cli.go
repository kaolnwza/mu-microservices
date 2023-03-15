package grpcClient

import (
	"fmt"
	"os"

	"github.com/kaolnwza/muniverse/order/internal/adapters/grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewSeerServiceClient() pb.SeerServiceClient {
	return pb.NewSeerServiceClient(grpcClientConn(os.Getenv("SEER_RPC_PORT"), "seer", os.Getenv("SEER_DNS")))
}

func NewVoucherServiceClient() pb.VoucherServiceClient {
	return pb.NewVoucherServiceClient(grpcClientConn(os.Getenv("VOUCHER_RPC_PORT"), "voucher", os.Getenv("VOUCHER_DNS")))
}

func NewWalletServiceClient() pb.WalletServiceClient {
	return pb.NewWalletServiceClient(grpcClientConn(os.Getenv("WALLET_RPC_PORT"), "wallet", os.Getenv("WALLET_DNS")))
}

func NewImageStorageServiceClient() pb.HoroSvcServiceClient {
	return pb.NewHoroSvcServiceClient(grpcClientConn(os.Getenv("STORER_RPC_PORT"), "stoage", os.Getenv("STORAGE_DNS")))
}

func grpcClientConn(port string, srvName string, dns string) *grpc.ClientConn {
	opts := []grpc.DialOption{}
	addr := dns + ":" + port
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
