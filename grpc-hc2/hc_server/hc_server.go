package main

import (
    "context"
    "log"
    "net"
    pb "example.com/go-hc-grpc/hc"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const (
    port = "127.0.0.1:50051"
)

type HealthServer struct {
    pb.UnimplementedHealthServer
}

func (s *HealthServer) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
    log.Printf("Received healthcheck request for service '%v'", in.Service)

    var svcStatus pb.HealthCheckResponse_ServingStatus;

    switch in.Service {
    case "foo":
        svcStatus = pb.HealthCheckResponse_SERVING;
    case "bar":
        svcStatus = pb.HealthCheckResponse_NOT_SERVING;
    default:
        svcStatus = pb.HealthCheckResponse_SERVICE_UNKNOWN;
    }

    return &pb.HealthCheckResponse{Status: svcStatus}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterHealthServer(s, &HealthServer{})
    reflection.Register(s)

    log.Printf("server listening at %v", lis.Addr())

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
