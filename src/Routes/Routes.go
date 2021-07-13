package Routes

import (
	"go_board/src/Controllers"
	"go_board/src/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("board", Controllers.GetBoard)
		api.POST("board", service.TokenAuthMiddleware(), Controllers.CreateBoard)
		api.GET("user", Controllers.GetUser)
		api.POST("user", Controllers.CreateUser)
		api.POST("user/login", Controllers.LoginUser)
	}
	return r
}
