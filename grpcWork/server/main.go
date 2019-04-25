package main

import (
	"context"
	"grpcWork/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct{}

func (s *server) Plus(ctx context.Context, in *pb.CalcRequest) (*pb.CalcReply, error) {
	result := in.NumberA + in.NumberB
	return &pb.CalcReply{Result: result}, nil
}

func (s *server) Auth(ctx context.Context, in *pb.AuthRequest) (*pb.IsAuth, error) {
	auth := in.Username == "admin" && in.Password == "test1234"

	return &pb.IsAuth{Result: auth}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("net.Listen error: %v", err)
	}
	// register [Plus] service to grpc server
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	pb.RegisterUserAuthServer(s, &server{})
	// register reflection
	reflection.Register(s)
	// grpc server bind tcp: 5001
	if err := s.Serve(listener); err != nil {
		log.Fatalf("no service : %v", err)
	}
}
