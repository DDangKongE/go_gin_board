package Controllers

import (
	"go_board/src/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBoard(c *gin.Context) {
	var board []Models.Board
	err := Models.GetAllBoard(&board)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, board)
	}
}

func CreateBoard(c *gin.Context) {
	var board Models.Board
	c.BindJSON(&board)
	err := Models.CreateBoard(&board)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, board)
	}
}
