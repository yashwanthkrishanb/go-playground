package main

import (
	"context"
	"fmt"
	ls "grocerylist/protos/listserver"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could  not connect : %v", err)
	}
	// defer conn.Close()
	c := ls.NewListServiceClient(conn)
	message := ls.Item{
		Item: "milk",
	}
	response, err := c.InsertListItem(context.Background(), &message)
	fmt.Println(response)
}
