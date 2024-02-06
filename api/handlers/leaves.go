package handlers

import (
	"github.com/BeeOntime/models"
	"github.com/gin-gonic/gin"
)

// CreateLeaveRequest godoc
// @Summary Create a new leave request
// @Description Create a new leave request
// @Tags staff
// @Accept  json
// @Produce  json
// @Param leave body models.LeaveRequest true "Leave object"
// @Success 200 {object} models.LeaveRequest
// @Router /staff/leave [post]
func (h *handlerV1) CreateStaffLeave(c *gin.Context) {
	var req models.LeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.storage.Postgres().CreateLeaveRequest(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
