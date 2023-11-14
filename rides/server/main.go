package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"learning-grpc/rides/pb"
)

func main() {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	srv := createServer()

	log.Printf("info: sever ready on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

func createServer() *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(timingInterceptor))

	var u Rides
	pb.RegisterRidesServer(srv, &u)
	reflection.Register(srv)

	return srv
}

func timingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Printf("info: %s took %v", info.FullMethod, duration)
	}()

	return handler(ctx, req)
}

type Rides struct {
	pb.UnimplementedRidesServer
}

func (r *Rides) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no metadata")
	}

	log.Printf("info: api_key %s", md["api_key"])

	// TODO: Validate req

	resp := pb.StartResponse{
		Id: req.Id,
	}

	return &resp, nil
}

func (r *Rides) End(ctx context.Context, req *pb.EndRequest) (*pb.EndResponse, error) {
	// TODO: Validate req

	resp := pb.EndResponse{
		Id: req.Id,
	}

	return &resp, nil
}

func (r *Rides) Location(stream pb.Rides_LocationServer) error {
	count := int64(0)
	driverId := ""

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			status.Error(codes.Internal, "can't read")
		}

		// TODO: update db
		driverId = req.DriverId
		count++

	}

	resp := pb.LocationResponse{
		DriverId: driverId,
		Count:    count,
	}

	return stream.SendAndClose(&resp)
}
