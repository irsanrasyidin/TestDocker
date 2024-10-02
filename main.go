package main

import (
	"context"
	"fmt"
)

func HandleRequest(ctx context.Context) (string, error) {
	return fmt.Sprintf("Hello from Lambda on GitHub Codespaces!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
