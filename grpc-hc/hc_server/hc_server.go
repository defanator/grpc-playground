package main

import (
    "context"
    "log"
    "net"
    pb "example.com/go-hc-grpc/hc"
    emptypb "google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const (
    port = "127.0.0.1:50051"
)

type HealthCheckServer struct {
    pb.UnimplementedHealthCheckServer
}

func (s *HealthCheckServer) GetStatus(ctx context.Context, in *emptypb.Empty) (*pb.Status, error) {
    log.Print("Received something")
    var my_status string = "okay";
    return &pb.Status{ResponseStatus: my_status}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterHealthCheckServer(s, &HealthCheckServer{})
    reflection.Register(s)

    log.Printf("server listening at %v", lis.Addr())

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
