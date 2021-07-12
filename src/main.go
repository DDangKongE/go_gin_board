package main

import (
	"fmt"
	"go_board/src/Config"
	"go_board/src/Models"
	"go_board/src/Routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status : ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Board{}, &Models.User{})

	r := Routes.SetupRouter()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Run()
}
