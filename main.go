package main

import (
	"context"
	"fmt"
	"github.com/jinivus/go-xivapi"
)

func main() {
	client := xivapi.NewClient(nil)

	// list all items containing the term "aiming"
	items, _, err := client.Search.Items(context.Background(), "aiming")
	if err != nil {
		println(err)
	}
	for _, item := range items.Items {
		fmt.Printf("%s - %d\n", item.Name, item.ID)
	}
}
