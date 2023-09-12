package main

import (
	"os"
	"io"
	"net/http"
	"time"
	"context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/context", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}