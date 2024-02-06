package handlers

import (
	"strconv"

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

// GetStaffLeaves godoc
// @Summary Get staff leaves
// @Description Get staff leaves
// @Tags staff
// @Accept  json
// @Produce  json
// @Param staff_id query string false "Staff ID"
// @Param date query string false "Date"
// @Success 200 {object} []models.LeaveRequest
// @Router /staff/leaves [get]
func (h *handlerV1) GetStaffLeaves(c *gin.Context) {
	staffID := c.Query("staff_id")
	// date := c.Query("date")
	limit := c.Query("limit")
	page := c.Query("page")
	id := c.Query("id")
	from := c.Query("from")
	to := c.Query("to")
	var limitInt, pageInt int

	if limit == "" {
		limitInt = 10
	} else {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid limit",
			})
			return
		}
		limitInt = int(limitInt)
	}
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid page",
			})
			return
		}
		pageInt = int(pageInt)
	}else{
		pageInt = 1
	}
	res, err := h.storage.Postgres().GetStaffLeaves(c, models.GetStaffLeavesRequest{
		StaffID: staffID,
		Limit:   limitInt,
		Page:    pageInt,
		Id:      id,
		From:    from,
		To:      to,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
