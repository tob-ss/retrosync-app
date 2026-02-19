package main

import (
	"context"
	"time"
	"log"
	"fmt"
	"net/http"
	"github.com/Khan/genqlient/graphql"

)

func quickScan(device string, userID int) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

		resp, err := getPaths(ctx, client, device, userID)

		if err != nil {
			log.Println("json.Compact:", err)
		}

		fmt.Println(resp)

	}
}