package handlers

import (
	"strconv"

	"github.com/BeeOntime/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Staffs
// @ID create-staff
// @Router /staffs [post]
// @Summary Create staff
// @Description Create staff
// @Accept json
// @Produce json
// @Param staff body models.Staff true "Staff"
// @Success 200 {object} http.Response{data=models.Staff} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateStaff(c *gin.Context) {

	var req models.Staff

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid argument"})
		return
	}

	isTheStaffExist, err := h.storage.Postgres().GetByLogin(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error()})
		return
	}
	if isTheStaffExist.Email != "" {
		c.JSON(400, gin.H{
			"message": "staff already exist"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "server error"})
		return
	}
	req.Password = string(hashedPassword)

	resp, err := h.storage.Postgres().CreateStaff(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "server error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    resp,
	})
}
func (h *handlerV1) GetStaffs(c *gin.Context) {

	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "10"
	}

	page, ok := c.GetQuery("page")
	if !ok {
		page = "1"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid argument"})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid argument"})
		return
	}
	req := models.GetStaffs{
		Limit: limitInt,
		Page:  pageInt,
	}

	resp, err := h.storage.Postgres().GetStaffs(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "server error : " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    resp,
	})
}
func (h *handlerV1) DeleteStaff(c *gin.Context) {

	id := c.Param("id")

	err := h.storage.Postgres().DeleteStaff(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "server error : " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
func (h *handlerV1) UpdateStaff(c *gin.Context) {

	var req models.Staff

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid argument"})
		return
	}
	if req.Id == "" {
		c.JSON(400, gin.H{
			"message": "invalid argument"})
		return
	}

	resp, err := h.storage.Postgres().UpdateStaff(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "server error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    resp,
	})
}
