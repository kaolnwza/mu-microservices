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

	"github.com/gin-gonic/gin"
	rpc "github.com/kaolnwza/muniverse/storage/grpc"
	"github.com/kaolnwza/muniverse/storage/middleware"
	"github.com/kaolnwza/muniverse/storage/proto/pb"
	service "github.com/kaolnwza/muniverse/storage/services"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/v1/storage")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, "ok")
		})

		upl := v1.Group("/uploads", middleware.NewGinMiddleware())
		{
			// upl.POST("/profile", service.UploadProfileImg)
			upl.POST("/", service.UploadImage)
		}
	}
	// opts =append(opts, grpc.ChainUnaryInterceptor(Login))
	// opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))

	rpcSrv, rpcLitn := newGrpcServer()

	pb.RegisterProfileServiceServer(rpcSrv, rpc.NewProfileRpcServer())
	pb.RegisterHoroSvcServiceServer(rpcSrv, rpc.NewHoroRpcService())
	pb.RegisterStorageServiceServer(rpcSrv, rpc.NewStorageRpcServer())

	defer rpcSrv.Stop()
	go func() {
		if err := rpcSrv.Serve(rpcLitn); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + os.Getenv("REST_PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Storage Rest Server Start on", os.Getenv("REST_PORT"))

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
