package main

import (
	"fmt"
	"net/http"
	"os"

	plainhttp "github.com/Noiidor/go-service-template/internal/app/plain-http"
)

func main() {
	if err := plainhttp.Run(os.Stdout, os.Stderr); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
