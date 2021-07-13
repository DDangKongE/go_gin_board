package Controllers

import (
	"fmt"
	"go_board/src/Models"
	"go_board/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBoard(c *gin.Context) {
	board := service.Board.GetAllBoard()
	defer c.JSON(http.StatusOK, board)

	fmt.Println(board)
}

func CreateBoard(c *gin.Context) {
	board := &Models.Board{}

	if err := c.BindJSON(board); nil != err {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	}

	if err := service.Board.CreateBoard(board); nil != err {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusCreated, board)
	}
}
