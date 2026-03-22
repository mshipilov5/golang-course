package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "golang-course/task2/api/proto/collectorv1"
	gh "golang-course/task2/internal/collector/adapter/github"
	grpcadapter "golang-course/task2/internal/collector/adapter/grpc"
	"golang-course/task2/internal/collector/usecase"
)

func main() {
	addr := getenv("COLLECTOR_LISTEN", ":50051")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	uc := &usecase.GetRepository{
		Fetcher: &gh.Client{},
	}
	srv := grpc.NewServer()
	pb.RegisterCollectorServiceServer(srv, &grpcadapter.Server{UC: uc})

	log.Printf("collector listening on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
