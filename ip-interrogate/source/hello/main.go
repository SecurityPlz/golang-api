package main

import (
	"github.com/SecurityPlz/golang-api/tree/functional/ip-interrogate/source/hello/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := handler.Create()
	lambda.Start(handler.Run)
}
