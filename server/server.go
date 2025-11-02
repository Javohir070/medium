package server

import (
	v1 "github.com/Javohir070/medium/server/v1"
	"github.com/Javohir070/medium/storage"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Strg storage.StorageI
}

func NewServer(opts Options) *gin.Engine {
	router := gin.New()

	handler := v1.New(&v1.HandlerV1{
		Strg: opts.Strg,
	})
	router.POST("/v1/users", handler.CreateUser)
	router.GET("/v1/user/:id", handler.GetUser)
	router.PUT("/v1/user/:id", handler.UpdateUser)
	router.DELETE("/v1/user/:id", handler.DeleteUser)

	router.POST("/v1/posts", handler.CreatePost)
	router.GET("/v1/post/:id", handler.GetPost)
	// router.PUT("/v1/post/:id", handler.UpdatePost)
	router.DELETE("/v1/post/:id", handler.DeletePost)

	return router
}
