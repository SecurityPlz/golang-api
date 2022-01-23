package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/likexian/whois"
)

type first struct {
	IP string 'json:"IP"'
}

func HandleRequest(ctx context.Context, IP first) (string, error) {
	return whois.Whois(IP), nil 
}

func handler() {
	lambda.Start(HandleRequest)
}