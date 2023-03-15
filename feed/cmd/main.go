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

	postgresdb "github.com/kaolnwza/muniverse/feed/internal/adapters/database/pgsql"
	ginapi "github.com/kaolnwza/muniverse/feed/internal/adapters/gin"
	grpcClient "github.com/kaolnwza/muniverse/feed/internal/adapters/grpc/client"
	service "github.com/kaolnwza/muniverse/feed/internal/application/core/services"
	rpc "github.com/kaolnwza/muniverse/feed/internal/infrastructure/grpc"
	handler "github.com/kaolnwza/muniverse/feed/internal/infrastructure/handlers"
	"github.com/kaolnwza/muniverse/feed/internal/infrastructure/postgres"
)

func main() {

	pg := postgresdb.NewPostgresDB(os.Getenv("DATABASE_URL"), "feed")
	defer pg.Close()
	pgTx := postgresdb.NewPostgresTransactor(pg)

	// storerConn := grpcClient.NewGrpcStorageClient()
	// imgStorer := pb.NewProfileServiceClient(storerConn)
	imgStorerConn := grpcClient.NewStorageServiceClient()
	imgStorerCli := rpc.NewGrpcStorerClient(imgStorerConn)

	postRepo := postgres.NewPostRepository(pgTx)
	postSvc := service.NewPostService(pgTx, postRepo, imgStorerCli)
	postHdr := handler.NewPostHandler(postSvc)

	likeRepo := postgres.NewLikeRepository(pgTx)
	likeSvc := service.NewLikeService(pgTx, likeRepo)
	likeHdr := handler.NewLikeHandler(likeSvc)

	comntRepo := postgres.NewCommentRepository(pgTx)
	comntSvc := service.NewCommentService(pgTx, comntRepo)
	comntHdr := handler.NewCommentHandler(comntSvc)

	r := ginapi.NewGinRouter()

	v1 := r.GROUP("/v1", ginapi.NewGinMiddleware())
	{
		v1feed := v1.GROUP("/feeds")
		{
			v1feed.POST("/", postHdr.CreatePostHandler)
			v1feed.GET("/", postHdr.GetAllPostHandler)
			v1post := v1feed.GROUP("/:post_uuid")
			{
				v1post.GET("/", postHdr.GetPostByPostUUIDHandler)
				v1post.DELETE("/", postHdr.DeletePostHandler)
				v1post.PUT("/", postHdr.UpdatePostHandler)

				v1like := v1post.GROUP("/likes")
				{
					v1like.POST("/", likeHdr.PostLikeHandler)
					v1like.DELETE("/", likeHdr.PostUnlikeHandler)
				}

				v1comnt := v1post.GROUP("/comments")
				{
					v1comnt.GET("/", comntHdr.GetCommentByPostUUIDHandler)
					v1comnt.POST("/", comntHdr.CreateCommentHandler)
					v1comnt.DELETE("/:comment_uuid", comntHdr.DeleteCommentByUUIDHandler)
				}
			}
		}
	}

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
