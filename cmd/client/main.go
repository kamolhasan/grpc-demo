package main

import (
	"context"
	"fmt"
	"log"
	"os"

	demo "github.com/kamolhasan/grpc-demo/src"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client connection failed with %v", err)
	}
	defer connection.Close()

	client := demo.NewDemoServiceClient(connection)

	msg := ""
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	resp, err := client.Fnc(context.Background(), &demo.FncRequest{Msg: msg})
	if err != nil {
		log.Fatalf("Client request failed with: %v", err)
	}

	fmt.Println("Succeed! Code: ", resp.Code)
	fmt.Println(resp.Msg)

}
