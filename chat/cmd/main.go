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

	ginAdapter "github.com/kaolnwza/muniverse/chat/internal/adapters/gin"
	"github.com/kaolnwza/muniverse/chat/internal/adapters/grpc/proto/pb"
	wsAdapter "github.com/kaolnwza/muniverse/chat/internal/adapters/ws"
	service "github.com/kaolnwza/muniverse/chat/internal/application/core/services"
	rpc "github.com/kaolnwza/muniverse/chat/internal/infrastructure/grpc"
	"github.com/kaolnwza/muniverse/chat/internal/infrastructure/handler"
	pg "github.com/kaolnwza/muniverse/chat/internal/infrastructure/pg"
	"github.com/kaolnwza/muniverse/chat/internal/infrastructure/ws"

	pgAdapter "github.com/kaolnwza/muniverse/chat/internal/adapters/database/pgsql"
	rpcAdapter "github.com/kaolnwza/muniverse/chat/internal/adapters/grpc/server"
)

func main() {
	upgrader := wsAdapter.NewWebsocketUpgrader()

	pgDB := pgAdapter.NewPostgresDB(os.Getenv("DATABASE_URL"), "chat")
	pgTx := pgAdapter.NewPostgresTransactor(pgDB)
	// wsHub := ws.NewWSHub()

	// cli := ws.NewClient()

	wsHub := ws.NewWSHub()
	wsCli := ws.NewWSChatClient(wsHub.Hub)
	wsSrv := ws.NewWSChatServer(wsHub.Hub, *upgrader, wsCli)

	roomRepo := pg.NewRoomRepository(pgTx)
	roomSvc := service.NewRoomService(pgTx, roomRepo)
	roomHdr := handler.NewRoomHandler(wsSrv, roomSvc)

	rpcLisn, rpcSrv := rpcAdapter.NewGrpcServer()
	rpcChatSrv := rpc.NewGrpcChatServer(roomSvc)
	pb.RegisterChatServiceServer(rpcSrv, rpcChatSrv)

	r := ginAdapter.NewGinRouter()

	go wsCli.Run()

	// v1 := r.GROUP("/v1", ginAdapter.NewGinMiddleware())
	v1 := r.GROUP("/v1", ginAdapter.NewGinMiddleware())
	{
		v1room := v1.GROUP("/rooms")
		{
			v1room.GET("/", roomHdr.GetRoomByUserUUIDHandler)
			v1roomId := v1room.GROUP("/:room_uuid")
			{
				v1roomId.GET("/join", roomHdr.JoinRoomHandler)
				v1roomMsg := v1roomId.GROUP("/messages")
				{
					_ = v1roomMsg
				}
			}

		}

	}

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

	fmt.Println("Feed Rest Server Start on", os.Getenv("REST_PORT"))

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
