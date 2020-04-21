package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dezkoat/dbdata/pb"
	"github.com/dezkoat/dbnet/api"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50001"
	defaultName = "world"
)

var PostServiceClient *pb.PostServiceClient
var Context *context.Context

func initGrpc() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	api.PostServiceClient = pb.NewPostServiceClient(conn)
}

func main() {
	initGrpc()

	router := mux.NewRouter()

	router.HandleFunc("/post", api.CreatePost).Methods("POST")
	router.HandleFunc("/post/{id}", api.ReadPost).Methods("GET")
	router.HandleFunc("/post/{id}", api.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", api.DeletePost).Methods("DELETE")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
