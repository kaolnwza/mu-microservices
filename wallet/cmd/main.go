package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	postgresdb "github.com/kaolnwza/muniverse/wallet/internal/adapters/database/pgsql"
	ginapi "github.com/kaolnwza/muniverse/wallet/internal/adapters/gin"
	"github.com/kaolnwza/muniverse/wallet/internal/adapters/grpc/proto/pb"
	grpcSrv "github.com/kaolnwza/muniverse/wallet/internal/adapters/grpc/server"
	service "github.com/kaolnwza/muniverse/wallet/internal/application/core/services"
	rpc "github.com/kaolnwza/muniverse/wallet/internal/infrastructure/grpc"
	handler "github.com/kaolnwza/muniverse/wallet/internal/infrastructure/handlers"
	postgres "github.com/kaolnwza/muniverse/wallet/internal/infrastructure/pgsql"
)

func main() {
	r := ginapi.NewGinRouter()
	middleware := ginapi.NewGinMiddleware()

	pg := postgresdb.NewPostgresDB(os.Getenv("DATABASE_URL"), "wallet")
	pgTx := postgresdb.NewPostgresTransactor(pg)

	walRepo := postgres.NewWalletRepository(pgTx)
	walSvc := service.NewWalletService(pgTx, walRepo)
	walHdr := handler.NewWalletHandler(walSvc)

	rpcLisn, rpcSrv := grpcSrv.NewGrpcServer()

	rpcSeerSrv := rpc.NewRpcWalletServer(walSvc)
	pb.RegisterWalletServiceServer(rpcSrv, rpcSeerSrv)

	v1 := r.GROUP("/v1", middleware)
	{
		v1wal := v1.GROUP("/wallets")
		{
			v1wal.GET("/", walHdr.GetFundByUserUUIDHandler)
		}
	}

	//rpc serve
	defer rpcSrv.Stop()
	go func() {
		if err := rpcSrv.Serve(rpcLisn); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()

	//rest serve
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + os.Getenv("REST_PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println("Wallet Rest Server Start on", os.Getenv("REST_PORT"))

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}

}
