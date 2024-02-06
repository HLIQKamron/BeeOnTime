package handlers

import (
	"fmt"
	"strconv"

	"github.com/BeeOntime/models"
	"github.com/gin-gonic/gin"
)

// CreateStaffEntry godoc
// @Summary Create a new staff entry
// @Description Create a new staff entry
// @Tags staff
// @Accept  json
// @Produce  json
// @Param entry body models.Entry true "Entry object"
// @Success 200 {object} models.Entry
// @Router /staff/entry [post]
func (h *handlerV1) CreateStaffEntry(c *gin.Context) {
	var req models.Entry
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.storage.Postgres().CreateStaffEntry(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// GetStaffEntries godoc
// @Summary Get staff entries
// @Description Get staff entries
// @Tags staff
// @Accept  json
// @Produce  json
// @Param staff_id query string false "Staff ID"
// @Param date query string false "Date"
// @Success 200 {object} []models.Entry
// @Router /staff/entries [get]
func (h *handlerV1) GetStaffEntries(c *gin.Context) {
	staffID := c.Query("staff_id")
	date := c.Query("date")
	limit := c.Query("limit")
	page := c.Query("page")
	id := c.Query("id")
	from := c.Query("from")
	to := c.Query("to")
	var limitInt int
	pageInt := 1
	if limit == "" {
		limitInt = 10
	} else {

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid argument"})
			return
		}
		limitInt = int(limitInt)

	}
	if page == "" {
		pageInt = 1
	} else {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid argument"})
			return
		}
		pageInt = int(pageInt)
	}
	fmt.Println("limitInt", limitInt)
	fmt.Println("pageInt", pageInt)
	res, err := h.storage.Postgres().GetStaffEntries(c, models.GetStaffEntries{
		StaffID: staffID,
		Date:    date,
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
func (h *handlerV1) DeleteStaffEntry(c *gin.Context) {
	id := c.Param("id")
	err := h.storage.Postgres().DeleteStaffEntry(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
func (h *handlerV1) UpdateStaffEntry(c *gin.Context) {
	var req models.Entry
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.storage.Postgres().UpdateStaffEntry(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
