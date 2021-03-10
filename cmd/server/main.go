package main

import (
	"bloggrpc/pkg/api"
	"bloggrpc/pkg/blogserver"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	mongodbURL   = "mongodb://localhost:27017"
	listenerPort = ":8080"
	amqpPort     = ":5672"
)

func main() {

	client, err := getMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	log.Println("Connection to db is done")

	s := grpc.NewServer()
	srv := blogserver.NewBlogServer(client.Database("blog").Collection("blogs"))
	api.RegisterBloggerServer(s, srv)

	l, err := net.Listen("tcp", amqpPort)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func getMongoClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	return client, err
}
