package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/shinhagunn/fastfood/proto/hello"
)

type helloServer struct {
	hello.UnimplementedHelloServiceServer
}

func (s *helloServer) Hello(ctx context.Context, in *emptypb.Empty) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Msg: "Hello world!"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &helloServer{})

	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(server.Serve(listen))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	gwmux := runtime.NewServeMux()
	err = hello.RegisterHelloServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal(err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
