package main

import (
	"time"
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	bookHotelExample(ctx)
}

func bookHotelExample(ctx context.Context) {
	select {
		case <-ctx.Done():
			fmt.Println("Timeout reached: ", ctx.Err())
			return
		case <-time.After(time.Microsecond):
			fmt.Println("Example booked")
			return
	}
}