package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) (string, error) {
	data, err := os.ReadFile("message/message.txt")
	if err != nil {
		return "Failed to read message", err
	}
	return string(data), nil
}

func main() {
	lambda.Start(Handler)
}
