package controller

import (
	"net/http"

	model "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/models"
	service "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/services"
	"github.com/gin-gonic/gin"
)

func Submit(c *gin.Context) {
	var process model.Processing

	if err := c.ShouldBindJSON(&process); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	err := service.NewProcessing(process)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Processing submitted !",
		})
	}
}

func Summary(c *gin.Context) {
	processings, err := service.ListAllProcessing()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		if len(processings) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "No processing found !",
			})
		} else {
			c.JSON(http.StatusOK, processings)
		}
	}
}
