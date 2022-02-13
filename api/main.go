package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kami988/go-1.18-sample/helloworld/def"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	def.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *def.HelloRequest) (*def.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &def.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	def.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
