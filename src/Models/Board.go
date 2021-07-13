package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Board struct {
	Board_ID uint   `gorm:"primaryKey"`
	Title    string `json:"title"`
	User_ID  uint   `json:"user_id"`
	Contents string `json:"contents"`
	User     User   `gorm:"foreignKey:user_id"`
}
