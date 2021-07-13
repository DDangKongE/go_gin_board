package main

import (
	"fmt"
	"go_board/src/Config"
	"go_board/src/Models"
	"go_board/src/Routes"

	"github.com/gin-contrib/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	// Config.DB, err = gorm.Open(mysql.Open("gorm.db"), &gorm.Config{Config.DbURL(Config.BuildDBConfig())})

	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("status : ", err)
	}

	sqlDB, err := Config.DB.DB()

	if err != nil {
		fmt.Println("status : ", err)
	}

	defer sqlDB.Close()

	Config.DB.AutoMigrate(&Models.Board{}, &Models.User{})
	Config.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.Board{}, &Models.User{})

	r := Routes.SetupRouter()

	r.Use(cors.Default())

	r.Run()
}
