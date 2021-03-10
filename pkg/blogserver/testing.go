package blogserver

import (
	"bloggrpc/pkg/api"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	mongodbURL = "mongodb://localhost:27017"
)

func NewTestBlog() *api.Blog {
	return &api.Blog{
		AuthorId: "0",
		Title:    "Статья по алгоритмам",
		Content:  "Много полезной информации...",
	}
}

func GetTestMongoClient() (*mongo.Client, error) {
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

func CreateTestBlogServer() (*BlogServer, func()) {
	mongoClient, err := GetTestMongoClient()
	if err != nil {
		panic(err)
	}
	srv := NewBlogServer(mongoClient.Database("test").Collection("blogs"))
	return srv, func() {
		if err := mongoClient.Database("test").Drop(context.Background()); err != nil {
			panic(err)
		}
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
		log.Println("drop and disconnect is done")
	}
}
