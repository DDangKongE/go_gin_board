package Controllers

import (
	"fmt"
	"go_board/src/Models"
	"go_board/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := service.User.GetUser()
	defer c.JSON(http.StatusOK, user)

	fmt.Println(user)
}

func CreateUser(c *gin.Context) {
	user := &Models.User{}

	if err := c.BindJSON(user); nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	if err := service.User.CreateUser(user); nil != err {
		c.AbortWithStatusJSON(http.StatusConflict, err)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

func LoginUser(c *gin.Context) {
	login := &Models.Login{}

	if err := c.BindJSON(login); nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	if err := service.User.LoginUser(login); nil != err {
		c.AbortWithStatusJSON(http.StatusConflict, err)
	} else {
		c.JSON(http.StatusCreated, login)
	}
}

// func GetUserOne(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user Models.User
// 	err := Models.GetUserOne(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }
