package blogserver

import (
	"bloggrpc/pkg/api"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (b *BlogServer) CreateBlog(ctx context.Context, request *api.CreateBlogRequest) (*api.CreateBlogResponse, error) {
	blog := request.GetBlog()

	data := Blog{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}
	//log.Println(data)

	res, err := b.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	resID := res.InsertedID.(primitive.ObjectID)
	blog.Id = resID.Hex()
	return &api.CreateBlogResponse{Blog: blog}, nil
}

func (b *BlogServer) ReadBlog(ctx context.Context, request *api.ReadBlogRequest) (*api.ReadBlogResponse, error) {
	id := request.GetId()
	log.Println(id)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot convert string to objectID")
	}
	res := b.collection.FindOne(
		ctx,
		bson.M{"_id": objID},
	)

	data := &Blog{}
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find blog with Object Id %s: %v", request.GetId(), err),
		)
	}

	response := &api.ReadBlogResponse{
		Blog: &api.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}
	return response, nil
}

func (b *BlogServer) UpdateBlog(ctx context.Context, request *api.UpdateBlogRequest) (*api.UpdateBlogResponse, error) {
	blog := request.GetBlog()
	oid, err := primitive.ObjectIDFromHex(blog.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot convert string to objectID")
	}
	_, err = b.collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.D{
			{"$set", bson.D{
				{"author_id", blog.GetAuthorId()},
				{"title", blog.GetTitle()},
				{"content", blog.GetContent()},
			}},
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot update document: %v", err))
	}
	return &api.UpdateBlogResponse{Blog: blog}, nil
}

func (b *BlogServer) DeleteBlog(ctx context.Context, request *api.DeleteBlogRequest) (*api.DeleteBlogResponse, error) {
	id := request.GetId()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot convert string to objectID")
	}
	_, err = b.collection.DeleteOne(
		ctx,
		bson.M{"_id": objID},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot delete document: %v", err))
	}
	return &api.DeleteBlogResponse{Success: true}, nil
}

func (b *BlogServer) ListBlogs(request *api.ListBlogRequest, server api.Blogger_ListBlogsServer) error {
	data := &Blog{}
	cursor, err := b.collection.Find(
		server.Context(),
		bson.D{},
	)
	if err != nil {
		return status.Errorf(codes.NotFound, fmt.Sprintf("cannot read document: %v", err))
	}
	defer cursor.Close(server.Context())

	for cursor.Next(server.Context()) {
		if err := cursor.Decode(data); err != nil {
			return status.Errorf(codes.Internal, "cannot decode data from cursor: %v", err)
		}
		response := &api.ListBlogResponse{
			Blog: &api.Blog{
				Id:       data.ID.Hex(),
				AuthorId: data.AuthorID,
				Title:    data.Title,
				Content:  data.Content,
			}}
		if err := server.Send(response); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("cannot send response: %v", err))
		}
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("unkown cursor error: %v", err))
	}
	return nil
}
