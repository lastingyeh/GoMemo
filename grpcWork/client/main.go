package main

import (
	"grpcWork/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect failure: %v", err)
	}
	defer conn.Close()

	calc := pb.NewCalculatorClient(conn)
	if reply, err := calc.Plus(context.Background(), &pb.CalcRequest{NumberA: 32, NumberB: 23}); err != nil {
		log.Fatalf("exec [Plus] failure: %v", err)
	} else {
		log.Printf("Result: %d", reply.Result)
	}

	auth := pb.NewUserAuthClient(conn)
	if isAuth, err := auth.Auth(context.Background(), &pb.AuthRequest{Username: "admin",
		Password: "test1234"}); err != nil {
		log.Fatalf("exec [Plus] failure: %v", err)
	} else {
		log.Printf("Result: %v", isAuth.Result)
	}

}
