// client_test.go

package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "looklook/app/grpc-test/pb"
	"testing"
)

const (
	address = "localhost:50051"
)

func TestServer_SayHello(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := pb.NewGreeterClient(conn)

	name := "Alice"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
