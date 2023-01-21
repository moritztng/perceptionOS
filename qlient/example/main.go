package main

import (
	"context"
	"fmt"

	"github.com/moritztng/perceptionOS/qlient"
)

func main() {
	ctx := context.Background()
	client := qlient.NewClient("http://localhost:8080/query")
	resp, err := qlient.AddDetection(ctx, client, 12, 0.5)
	fmt.Println(resp.AddDetection.GetPerson(), err)
}
