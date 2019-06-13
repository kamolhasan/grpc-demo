package main

import (
	"context"
	"errors"
	"log"
	"net"

	demo "github.com/kamolhasan/grpc-demo/src"
	"google.golang.org/grpc"
)

type DemoServer struct{}

func (d *DemoServer) Fnc(ctx context.Context, request *demo.FncRequest) (*demo.FncResponse, error) {
	log.Println("message received at server: ", request.Msg)
	if len(request.Msg) > 0 {
		return &demo.FncResponse{Msg: "I got your message! Thanks! :D", Code: 200}, nil
	} else {
		return &demo.FncResponse{Msg: "Bad Request!", Code: 400}, errors.New("bad request")
	}
}

func main() {
	srv := grpc.NewServer()
	var demoSrv DemoServer
	demo.RegisterDemoServiceServer(srv, &demoSrv)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Cloud not listen to :8080: %v", err)
	}

	log.Fatal(srv.Serve(lis))
}
