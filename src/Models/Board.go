package Models

import (
	"go_board/src/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Board struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	User     User
	Contents string `json:"contents"`
}

func GetAllBoard(board *[]Board) (err error) {
	if err = Config.DB.Find(board).Error; err != nil {
		return err
	}
	return nil
}

func CreateBoard(board *Board) (err error) {
	if err := Config.DB.Create(&board).Error; err != nil {
		return err
	}
	return nil
}
