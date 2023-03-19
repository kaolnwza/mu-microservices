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

	postgresdb "github.com/kaolnwza/muniverse/order/internal/adapters/database/pgsql"
	ginapi "github.com/kaolnwza/muniverse/order/internal/adapters/gin"
	grpcClient "github.com/kaolnwza/muniverse/order/internal/adapters/grpc/client"
	service "github.com/kaolnwza/muniverse/order/internal/application/core/services"
	rpc "github.com/kaolnwza/muniverse/order/internal/infrastructure/grpc"
	handler "github.com/kaolnwza/muniverse/order/internal/infrastructure/handlers"
	"github.com/kaolnwza/muniverse/order/internal/infrastructure/postgres"
)

func main() {
	middleware := ginapi.NewGinMiddleware()

	pg := postgresdb.NewPostgresDB(os.Getenv("DATABASE_URL"), "order")
	pgTx := postgresdb.NewPostgresTransactor(pg)

	//rpc client
	seerConn := grpcClient.NewSeerServiceClient()
	seerCli := rpc.NewGrpcSeerClient(seerConn)

	voucherConn := grpcClient.NewVoucherServiceClient()
	voucherCli := rpc.NewGrpcVoucherClient(voucherConn)

	walletConn := grpcClient.NewWalletServiceClient()
	walletCli := rpc.NewGrpcWalletClient(walletConn)

	imgConn := grpcClient.NewImageStorageServiceClient()
	imgCli := rpc.NewImageStorageServiceClient(imgConn)

	chatConn := grpcClient.NewChatServiceClient()
	chatCli := rpc.NewChatServiceClient(chatConn)

	// storerConn := grpcClient.NewStorageServiceClient()
	// storerCli := rpc.NewGrpcStorerClient(storerConn)
	// // imgStorerConn := grpcClient.NewProfileServiceClient()
	// // wtf := rpc.NewGrpcStorerClient(imgStorerConn)

	horoRepo := postgres.NewHoroRepository(pgTx)
	horoSvc := service.NewHoroService(pgTx, horoRepo, seerCli, imgCli)
	horoHdr := handler.NewHoroHandler(horoSvc)

	orderRepo := postgres.NewHoroOrderRepository(pgTx)
	orderSvc := service.NewHoroOrderService(pgTx, orderRepo, voucherCli, walletCli, horoSvc, seerCli, chatCli)
	orderHdr := handler.NewHoroOrderHandler(orderSvc)

	r := ginapi.NewGinRouter()

	v1 := r.GROUP("/v1")
	{
		v1odr := v1.GROUP("/orders", middleware)
		{
			v1odr_id := v1odr.GROUP("/:order_uuid")
			{
				v1odr_id.POST("/", orderHdr.CreateNewHoroOrderHandler)
				v1odr_id.PATCH("/confirmed", orderHdr.UpdateOrderStatusConfirmedByUUIDHandler)
				v1odr_id.PATCH("/success", orderHdr.UpdateOrderStatusSuccessByUUIDHandler)
			}

			v1odr_user := v1odr.GROUP("/users")
			{
				v1odr_user.GET("/", orderHdr.GetOrderByUserUUIDHandler)
				v1odr_user.GET("/history", orderHdr.GetOrderHistoryByUserUUIDHandler)
			}

			v1odr_seer := v1odr.GROUP("/seers", middleware)
			{
				v1odr_seer.GET("/", orderHdr.GetUpcomingCustomerOrderHandler)
				v1odr_seer.GET("/history", orderHdr.GetCustomerOrderHistoryHandler)

			}
		}

		v1horo := v1.GROUP("/horoes", middleware)
		{
			v1horo.POST("/", horoHdr.CreateHoroServiceHandler)
			v1horo.GET("/", horoHdr.GetAllHoroServiceHandler)
			v1horo_id := v1horo.GROUP("/:horo_uuid")
			{
				v1horo_id.PUT("/", horoHdr.UpdateHoroServiceHandler)
				v1horo_id.GET("/", horoHdr.GetHoroByHoroUUIDHandler)
				v1horo_id.PUT("/event", horoHdr.UpdateHoroOnEventHandler)
				v1horo_id.PUT("/status", horoHdr.UpdateHoroStatusHandler)

				v1horo_id.GET("/available", horoHdr.GetHoroAvailableEventByDateHandler)
				v1horo_id.GET("/schedule", horoHdr.GetHoroScheduleEventByDate)

				v1horo_id.POST("/order", orderHdr.CreateNewHoroOrderHandler)
			}

			v1horo_seer := v1horo.GROUP("/seers")
			{
				v1horo_seer_id := v1horo_seer.GROUP("/:seer_uuid")
				{
					v1horo_seer_id.GET("/", horoHdr.GetAllHoroServiceBySeerUUIDHandler)
				}
			}
		}
	}

	//rpc serve
	// defer rpcSrv.Stop()
	// if err := rpcSrv.Serve(rpcLisn); err != nil {
	// 	log.Fatalf("Failed to serve: %v\n", err)
	// }

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

	fmt.Println("Order Rest Server Start on", os.Getenv("REST_PORT"))

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}

}
