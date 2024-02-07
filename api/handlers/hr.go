package handlers

import (
	"github.com/BeeOntime/models"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreateHr(c *gin.Context) {
	var req models.Hr
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.storage.Postgres().CreateHr(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// GetHrs godoc
// @Summary Get hrs
// @Description Get hrs
// @Tags hr
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Hr
// @Router /hr [get]
func (h *handlerV1) GetHrs(c *gin.Context) {

	var resp models.GetHrs
	resp.Id = c.Query("id")

	res, err := h.storage.Postgres().GetHrs(c.Request.Context(), resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
//DeleteHr godoc
// @Summary Delete hr
// @Description Delete hr
// @Tags hr
// @Accept  json
// @Produce  json
// @Param id param string true "id"
// @Success 200 {object} models.Hr
// @Router /hr/{id} [delete]
func (h *handlerV1) DeleteHr(c *gin.Context) {
	id := c.Param("id")
	err := h.storage.Postgres().DeleteHr(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}