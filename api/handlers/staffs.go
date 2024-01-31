package handlers

import (
	"github.com/BeeOntime/models"
	"github.com/gin-gonic/gin"
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
