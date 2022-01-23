package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/likexian/whois"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

// Handler - interface
type Handler interface {
	Run(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (Response, error)
}

func hello() {
	var first string 
	fmt.Scanln(&first)
	resp, err := whois.Whois(first)
	iff err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}