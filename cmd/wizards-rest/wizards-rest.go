package main

import (
	"fmt"
	"net/http"
	"os"

	wizardsrest "github.com/Noiidor/go-service-template/internal/app/wizards-rest"
)

func main() {
	if err := wizardsrest.Run(os.Stdout, os.Stderr); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
