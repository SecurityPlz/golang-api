package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := handler.Create()
	lambda.Start(handler.Run)
}
