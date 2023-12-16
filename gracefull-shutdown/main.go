package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := http.Server{Addr: ":3000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)

		w.Write([]byte("Hello world"))
	})

	go func() {
		fmt.Println("Server is running at http://localhost:3000")
		err := server.ListenAndServe()
		if err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not listen on %s, %v\n", server.Addr, err)
		}
	}()

	stopChannel := make(chan os.Signal, 1)

	signal.Notify(stopChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stopChannel

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("Server stopped!")

}
