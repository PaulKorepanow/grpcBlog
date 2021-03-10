package blogserver

import (
	"bloggrpc/pkg/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

type BlogServer struct {
	api.UnimplementedBloggerServer
	collection *mongo.Collection
}

func NewBlogServer(c *mongo.Collection) *BlogServer {
	return &BlogServer{
		collection: c,
	}
}

var example = bson.D{
	{"$set", bson.D{
		{"author_id", 1},
		{"title", 1},
		{"content", 1},
	}},
}
