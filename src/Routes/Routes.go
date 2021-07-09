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
		api.POST("board", Controllers.GetBoard)
	}
	return r
}
