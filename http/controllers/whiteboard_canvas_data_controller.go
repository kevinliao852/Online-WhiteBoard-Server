package controllers

import (
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WhiteboardCanvasDataController struct {
}

func (wcdc *WhiteboardCanvasDataController) GetWhiteboardCanvasDataById(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.String(http.StatusUnprocessableEntity, "")
		return
	}

	whiteboardId, err := strconv.Atoi(id)

	if err != nil {
		c.String(http.StatusUnprocessableEntity, "")
		return
	}

	var whiteboardCanvasDataList []models.WhiteboardCanvasData

	err = models.GetWhiteboardCanvasDataById(&whiteboardCanvasDataList, uint(whiteboardId))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, whiteboardCanvasDataList)
	}

}
