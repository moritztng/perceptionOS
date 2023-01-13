package main

import (
	"context"

	"github.com/moritztng/perceptionOS/storage"
)

func main() {
	ctx := context.Background()
	storage := storage.NewStorage("localhost:9000", "77Ecj8b8avVP6WCG", "Cw1kJUY9z46U33VUwajoeh2nY5tLQWGf", false)
	storage.Store(ctx, "images", "image.jpg", "image.jpg", "image/jpeg")
}
