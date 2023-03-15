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

	postgresdb "github.com/kaolnwza/muniverse/seer/internal/adapters/database/pgsql"
	ginapi "github.com/kaolnwza/muniverse/seer/internal/adapters/gin"
	grpcClient "github.com/kaolnwza/muniverse/seer/internal/adapters/grpc/client"
	"github.com/kaolnwza/muniverse/seer/internal/adapters/grpc/proto/pb"
	grpcSrv "github.com/kaolnwza/muniverse/seer/internal/adapters/grpc/server"
	service "github.com/kaolnwza/muniverse/seer/internal/application/core/services"
	rpc "github.com/kaolnwza/muniverse/seer/internal/infrastructure/grpc"
	handler "github.com/kaolnwza/muniverse/seer/internal/infrastructure/handlers"
	"github.com/kaolnwza/muniverse/seer/internal/infrastructure/pgsql"
)

func main() {

	pg := postgresdb.NewPostgresDB(os.Getenv("DATABASE_URL"), "seer")
	pgTx := postgresdb.NewPostgresTransactor(pg)

	userConn := grpcClient.NewUserServiceClient()
	userCli := rpc.NewGrpcUserClient(userConn)

	storerConn := grpcClient.NewStorageServiceClient()
	storerCli := rpc.NewGrpcStorerClient(storerConn)
	// imgStorerConn := grpcClient.NewProfileServiceClient()
	// wtf := rpc.NewGrpcStorerClient(imgStorerConn)

	seerRepo := pgsql.NewSeerRepository(pgTx)
	seerSvc := service.NewSeerService(seerRepo, pgTx, userCli, storerCli)
	seerHdr := handler.NewSeerHandler(seerSvc)

	rpcLisn, rpcSrv := grpcSrv.NewGrpcServer()
	rpcSeerSrv := rpc.NewGrpcSeerServer(seerSvc)
	pb.RegisterSeerServiceServer(rpcSrv, rpcSeerSrv)

	r := ginapi.NewGinRouter()
	// r.GET("/health", func(ctx *gin.Context) {

	// 	ctx.JSON(http.StatusOK, map[string]string{"kuy": "big"})
	// })

	v1 := r.GROUP("/v1")
	{
		v1seer := v1.GROUP("/seers")
		{
			v1seerUUID := v1seer.GROUP("/:seer_uuid")
			{
				v1seerUUID.GET("/", seerHdr.GetSeerByUUID)
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

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println("Seer Rest Server Start on", os.Getenv("REST_PORT"))

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}

}
