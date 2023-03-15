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
	handler "github.com/kaolnwza/muniverse/voucher/handlers"
	"github.com/kaolnwza/muniverse/voucher/proto/pb"
	service "github.com/kaolnwza/muniverse/voucher/services"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1voc := v1.Group("/vouchers")
		{
			v1voc_code := v1voc.Group("/codes")
			{
				v1voc_codeid := v1voc_code.Group("/:code")
				{
					v1voc_codeid.GET("/", handler.GetVoucherByCodeHandler)
				}
			}
		}
	}
	s, lis := newGrpcServer()

	pb.RegisterVoucherServiceServer(s, &service.RpcVoucherServer{})

	defer s.Stop()
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:           ":" + os.Getenv("REST_PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("User Rest Server Start on", os.Getenv("REST_PORT"))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}
