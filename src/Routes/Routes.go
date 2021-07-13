package Routes

import (
	"go_board/src/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("board", Controllers.GetBoard)
		api.POST("board", Controllers.CreateBoard)
		api.GET("user", Controllers.GetUser)
		api.POST("user", Controllers.CreateUser)
		api.POST("user/login", Controllers.LoginUser)
	}
	return r
}
