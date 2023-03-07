package main

import (
	"context"
	"encoding/json"
	"fmt"
	ls "grocerylist/protos/listserver"
	"io"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InsertItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could  not connect : %v", err)
	}
	defer conn.Close()
	c := ls.NewListServiceClient(conn)
	message := ls.Item{
		Item: item,
	}
	response, err := c.InsertListItem(context.Background(), &message)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response)
	log.Println(item)
	io.WriteString(w, response.String())
}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could  not connect : %v", err)
	}
	defer conn.Close()
	c := ls.NewListServiceClient(conn)
	message := ls.ItemReq{
		ItemName: item,
	}
	response, err := c.DeleteListItem(context.Background(), &message)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response)
	log.Println(item)
	io.WriteString(w, response.String())
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could  not connect : %v", err)
	}
	defer conn.Close()
	c := ls.NewListServiceClient(conn)
	message := ls.ItemReq{
		ItemName: item,
	}
	response, err := c.FindListItem(context.Background(), &message)
	log.Println(response, err)
	if response == nil {
		fmt.Println("Not found")
		io.WriteString(w, "Not Found")
		return
	}
	fmt.Println(response)
	log.Println(item)
	io.WriteString(w, response.String())
}
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	newName := r.FormValue("newName")
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could  not connect : %v", err)
	}
	defer conn.Close()
	c := ls.NewListServiceClient(conn)
	message := ls.UpdateReq{
		ItemName: name,
		NewName:  newName,
	}
	response, err := c.UpdateListItem(context.Background(), &message)
	log.Println(response, err)
	if response == nil {
		fmt.Println("Not found")
		io.WriteString(w, "Not Found")
		return
	}
	fmt.Println(response)
	io.WriteString(w, response.String())
}
func GetList(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	client := ls.NewListServiceClient(conn)
	in := &ls.Empty{}
	stream, err := client.GetList(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)
	response := []string{}
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp)
			response = append(response, resp.Item)
		}
	}()
	<-done
	log.Println("finished")

	type JsonResp struct {
		Count int
		Data  []string
	}
	out := JsonResp{Count: len(response), Data: response}

	resp, err := json.Marshal(out)
	log.Println(out)
	log.Println(string(resp))

	if err != nil {
		io.WriteString(w, fmt.Sprintf("Error in Marshalling : %v", err))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
func main() {

	http.HandleFunc("/find", GetItem)
	http.HandleFunc("/delete", DeleteItem)
	http.HandleFunc("/new", InsertItem)
	http.HandleFunc("/update", UpdateItem)
	http.HandleFunc("/", GetList)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Failed to Listen at port 3000 : %v", err)
	}
}
