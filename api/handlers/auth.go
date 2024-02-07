package handlers

import (
	"github.com/BeeOntime/models"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) Login(c *gin.Context) {
	var req models.Login

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.storage.Postgres().LoginCheck(c, req.Login, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Header().Set("Authorization", token)

	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
