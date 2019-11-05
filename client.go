package main

import (
	"log"
	"net/rpc"

	"github.com/aws/aws-lambda-go/lambda/messages"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	req := messages.InvokeRequest{}
	res := messages.InvokeResponse{}

	err = client.Call("Function.Invoke", req, &res)
	if err != nil {
		log.Fatal("invoke error:", err)
	}

	log.Println("Got response: ", string(res.Payload))
}
