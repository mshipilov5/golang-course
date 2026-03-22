package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "golang-course/task2/docs"
	grpcadapter "golang-course/task2/internal/gateway/adapter/grpc"
	ghandler "golang-course/task2/internal/gateway/handler"
	"golang-course/task2/internal/gateway/usecase"
)

func main() {
	collectorAddr := getenv("COLLECTOR_ADDR", "localhost:50051")
	conn, err := grpc.NewClient(collectorAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc dial: %v", err)
	}
	defer conn.Close()

	uc := &usecase.GetRepository{
		Reader: grpcadapter.NewCollectorClient(conn),
	}
	h := &ghandler.HTTP{UC: uc}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	h.Register(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := getenv("GATEWAY_LISTEN", ":8080")
	log.Printf("gateway listening on %s (collector=%s)", addr, collectorAddr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("http: %v", err)
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
