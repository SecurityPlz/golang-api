package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/likexian/whois"
)

func HandleRequest() {
	var first string
	fmt.Scanln(&first)
	resp, err := whois.Whois(first)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}

func main() {
	lambda.Start(HandleRequest)
}
