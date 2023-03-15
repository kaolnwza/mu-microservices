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

	postgresdb "github.com/kaolnwza/muniverse/user/internal/adapters/database/pgsql"
	ginapi "github.com/kaolnwza/muniverse/user/internal/adapters/gin"
	grpcClient "github.com/kaolnwza/muniverse/user/internal/adapters/grpc/client"
	"github.com/kaolnwza/muniverse/user/internal/adapters/grpc/proto/pb"
	grpcSrv "github.com/kaolnwza/muniverse/user/internal/adapters/grpc/server"
	service "github.com/kaolnwza/muniverse/user/internal/application/core/services"
	rpc "github.com/kaolnwza/muniverse/user/internal/infrastructure/grpc"
	handler "github.com/kaolnwza/muniverse/user/internal/infrastructure/handlers"
	"github.com/kaolnwza/muniverse/user/internal/infrastructure/postgres"
)

func main() {
	middleware := ginapi.NewGinMiddleware()

	pg := postgresdb.NewPostgresDB(os.Getenv("DATABASE_URL"), "user")
	pgTx := postgresdb.NewPostgresTransactor(pg)

	// storerConn := grpcClient.NewGrpcStorageClient()
	// imgStorer := pb.NewProfileServiceClient(storerConn)
	imgStorerConn := grpcClient.NewProfileServiceClient()
	imgStorerCli := rpc.NewGrpcStorerClient(imgStorerConn)

	walletConn := grpcClient.NewWalletServiceClient()
	walletCli := rpc.NewGrpcWalletClient(walletConn)

	userRepo := postgres.NewUserRepository(pgTx)
	userSvc := service.NewUserService(userRepo, pgTx, imgStorerCli, walletCli)
	userHdr := handler.NewUserHandler(userSvc)

	//GRPC
	rpcLisn, rpcSrv := grpcSrv.NewGrpcUserServer()
	rpcUserSrv := rpc.NewGrpcUserServer(userSvc)

	pb.RegisterUserServiceServer(rpcSrv, rpcUserSrv)

	r := ginapi.NewGinRouter()
	// r.GET("/health", func(ctx *gin.Context) {

	// 	ctx.JSON(http.StatusOK, map[string]string{"kuy": "big"})
	// })

	v1 := r.GROUP("/v1", middleware)
	{
		v1user := v1.GROUP("/users")
		{
			v1user.GET("/health", userHdr.Health)

			v1user.GET("/", userHdr.GetUserByTokenHandler)
			v1user.PUT("/", userHdr.UpdateUserHandler)
			v1userUUID := v1user.GROUP("/:user_uuid")
			{
				v1userUUID.GET("/", userHdr.GetUserByUUIDHandler)

			}
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

	fmt.Println("User Rest Server Start on", os.Getenv("REST_PORT"))

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}

}
