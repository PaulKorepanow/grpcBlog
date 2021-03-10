package blogserver

import (
	"bloggrpc/pkg/api"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlogServer_CreateBlog(t *testing.T) {
	srv, dropAndClose := CreateTestBlogServer()
	defer dropAndClose()
	req := &api.CreateBlogRequest{Blog: NewTestBlog()}

	res, err := srv.CreateBlog(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	assert.NotEqual(t, "", res.Blog.Id)
	assert.Equal(t, req.GetBlog().AuthorId, res.GetBlog().AuthorId)
	assert.Equal(t, req.GetBlog().Title, res.GetBlog().Title)
	assert.Equal(t, req.GetBlog().Content, res.GetBlog().Content)
}

func TestBlogServer_ReadBlog(t *testing.T) {
	srv, dropAndClose := CreateTestBlogServer()
	defer dropAndClose()

	req := &api.CreateBlogRequest{Blog: NewTestBlog()}

	createBlogResponse, err := srv.CreateBlog(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, createBlogResponse)

	readBlogResponse, err := srv.ReadBlog(
		context.Background(),
		&api.ReadBlogRequest{Id: createBlogResponse.Blog.Id},
	)

	assert.NoError(t, err)
	assert.NotNil(t, readBlogResponse)

	assert.Equal(t, createBlogResponse.Blog, readBlogResponse.Blog)
}

func TestBlogServer_UpdateBlog(t *testing.T) {
	srv, dropAndClose := CreateTestBlogServer()
	defer dropAndClose()

	req := &api.CreateBlogRequest{Blog: NewTestBlog()}

	createBlogResponse, err := srv.CreateBlog(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, createBlogResponse)

	updatedBlog := &api.Blog{
		Id:       createBlogResponse.Blog.Id,
		AuthorId: "123",
		Title:    "Статья о жизни",
		Content:  "Ничего интересного...",
	}
	updateBlogResponse, err := srv.UpdateBlog(
		context.Background(),
		&api.UpdateBlogRequest{Blog: updatedBlog},
	)
	assert.NoError(t, err)
	assert.NotEqual(t, NewTestBlog().GetAuthorId(), updateBlogResponse.Blog.GetAuthorId())
	assert.NotEqual(t, NewTestBlog().GetTitle(), updateBlogResponse.Blog.GetTitle())
	assert.NotEqual(t, NewTestBlog().GetContent(), updateBlogResponse.Blog.GetContent())
}

func TestBlogServer_DeleteBlog(t *testing.T) {
	srv, dropAndClose := CreateTestBlogServer()
	defer dropAndClose()

	req := &api.CreateBlogRequest{Blog: NewTestBlog()}

	createBlogResponse, err := srv.CreateBlog(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, createBlogResponse)

	deleteBlogResponse, err := srv.DeleteBlog(
		context.Background(),
		&api.DeleteBlogRequest{Id: req.Blog.Id},
	)
	assert.NoError(t, err)
	assert.True(t, deleteBlogResponse.GetSuccess())
}

func TestBlogServer_ListBlogs(t *testing.T) {
	t.Skip()
	srv, dropAndClose := CreateTestBlogServer()
	defer dropAndClose()

	blogs := []*api.Blog{
		NewTestBlog(),
		{
			Id:       "2",
			AuthorId: "1",
			Title:    "Статья по информатике",
			Content:  "Бинарный код",
		},
	}

	for _, blog := range blogs {
		req := &api.CreateBlogRequest{Blog: blog}
		createBlogResponse, err := srv.CreateBlog(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createBlogResponse)
	}

	err := srv.ListBlogs(&api.ListBlogRequest{}, nil)
	assert.NoError(t, err)
}
