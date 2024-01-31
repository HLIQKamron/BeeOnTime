package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @ID ping
// @Router /ping [GET]
// @Summary returns "pong" message
// @Description this returns "pong" messsage to show service is working
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
func (h *handlerV1) Ping(c *gin.Context) {
	fmt.Println("ping")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
