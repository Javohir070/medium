package v1

import (
	"net/http"

	"github.com/Javohir070/medium/server/models"
	"github.com/Javohir070/medium/storage/repo"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreatePost(ctx *gin.Context) {
	var req models.CreatePost

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	post, err := h.Strg.Post().Create(ctx.Request.Context(), &repo.Post{
		Title:     req.Title,
		Body:      req.Body,
		UserID:    req.UserID,
		Published: req.Published,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, post)
}

func (h *handlerV1) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := h.Strg.Post().Get(ctx, id)

	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, models.Post{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		UserID:    post.UserID,
		Published: post.Published,
	})
}

func (h *handlerV1) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Strg.Post().Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete success",
	})
}
