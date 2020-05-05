package main

import (
	"context"
	"fmt"
	"log"

	hellopb "github.com/mel-github/grpc-golang-app/hello_proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Print("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "Berry"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
}
