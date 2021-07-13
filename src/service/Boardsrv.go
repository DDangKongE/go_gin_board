package service

import (
	"go_board/src/Config"
	"go_board/src/Models"
	"sync"

	"gopkg.in/validator.v2"
)

var Board = &boardService{
	mutex: &sync.Mutex{},
}

type boardService struct {
	mutex *sync.Mutex
}

func (srv *boardService) GetAllBoard() *[]Models.Board {
	board := &[]Models.Board{}

	if err := Config.DB.Preload("User").Find(&board).Error; err != nil {
		return nil
	}

	return board
}

func (srv *boardService) CreateBoard(board *Models.Board) (err error) {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	if err := validator.Validate(board); err != nil {
		return err
	}

	tx := Config.DB.Begin()
	if err := tx.Create(board).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}
