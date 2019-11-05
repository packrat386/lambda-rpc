package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	log.Println("got a request!")
	return "Hello World!", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
