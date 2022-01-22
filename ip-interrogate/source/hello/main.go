package main

import (
	"github.com/SecurityPlz/golang-api/tree/master/ip-interrogate/source/hello/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := handler.Create()
	lambda.Start(handler.Run)
}
