package v1

import (
	"github.com/Javohir070/medium/server/models"
	"github.com/Javohir070/medium/storage/repo"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreateUser(ctx *gin.Context) {
	var req models.CreateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.Strg.User().Create(ctx, &repo.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(201, user)
}

func (h *handlerV1) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.Strg.User().Get(ctx, id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})
}

func (h *handlerV1) UpdateUser(ctx *gin.Context) {
	var req models.UpdateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	id := ctx.Param("id")

	err := h.Strg.User().Update(ctx, &repo.UpdateUser{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User updated successfully"})
}

func (h *handlerV1) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.Strg.User().Delete(ctx, id); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
