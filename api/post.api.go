package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dezkoat/dbdata/pb"
	"github.com/gorilla/mux"
)

var PostServiceClient pb.PostServiceClient
var Context context.Context

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &pb.Post{
		Title:   "Hello world",
		Content: "Lorem ipsum",
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := PostServiceClient.CreatePost(
		context,
		&pb.CreatePostRequest{Post: post},
	)

	if err != nil {
		log.Fatalf("[CreatePost] Failed: %v", err)
	}

	if res.GetPost() != nil {
		log.Printf("CreatePost success")
	} else {
		log.Printf("CreatePost fail")
	}
}

func ReadPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := PostServiceClient.ReadPost(
		context,
		&pb.ReadPostRequest{Id: vars["id"]},
	)

	if err != nil {
		log.Fatalf("[ReadPost] Failed: %v", err)
	}

	if res.GetPost() != nil {
		log.Printf("Read success")
	} else {
		log.Printf("Read fail")
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	post := &pb.Post{
		Id:      vars["id"],
		Title:   "Hello world",
		Content: "Lorem ipsum",
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := PostServiceClient.UpdatePost(
		context,
		&pb.UpdatePostRequest{Post: post},
	)

	if err != nil {
		log.Fatalf("[UpdatePost] Failed: %v", err)
	}

	if res.GetPost() != nil {
		log.Printf("Update success")
	} else {
		log.Printf("Update fail")
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := PostServiceClient.DeletePost(
		context,
		&pb.DeletePostRequest{Id: vars["id"]},
	)

	if err != nil {
		log.Fatalf("[DeletePost] Failed: %v", err)
	}

	if res.GetSuccess() {
		log.Printf("Delete success")
	} else {
		log.Printf("Delete fail")
	}
}
