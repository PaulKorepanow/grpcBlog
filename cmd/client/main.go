package main

import (
	"github.com/PaulKorepanow/grpcBlog/cmd"
	"log"
)

func main() {
	log.Println("Starting client creation ...")
	cmd.Execute()
	//conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client = api.NewBloggerClient(conn)
}
