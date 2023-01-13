package main

import (
	"context"
	"fmt"

	"github.com/moritztng/perceptionOS/qlient"
)

func main() {
	ctx := context.Background()
	client := qlient.NewClient("http://localhost:8080/query")
	resp, err := qlient.AddImage(ctx, client, "file.png")
	fmt.Println(resp.AddImage.Id, err)
}
