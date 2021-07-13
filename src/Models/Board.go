package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Board struct {
	Board_ID uint   `gorm:"primaryKey"`
	Title    string `json:"title" validate:"min=4,max=50"`
	User_ID  uint   `json:"user_id" validate:"min=1,max=20000"`
	Contents string `json:"contents" validate:"min=1"`
	User     User   `gorm:"foreignKey:user_id"`
}
