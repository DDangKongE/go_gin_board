package service

import (
	"fmt"
	"go_board/src/Config"
	"go_board/src/Models"
	"sync"

	"github.com/lestrrat-go/jwx/jwt"
)

var User = &userService{
	mutex: &sync.Mutex{},
}

type userService struct {
	mutex *sync.Mutex
}

// func (srv *boardService) CreateBoa2rd(board *Models.Board) (err error) {
// 	srv.mutex.Lock()
// 	defer srv.mutex.Unlock()

// 	tx := Config.DB.Begin()
// 	if err := tx.Create(board).Error; nil != err {
// 		tx.Rollback()

// 		return err
// 	}
// 	tx.Commit()

// 	return nil
// }

func (srv *userService) GetUser() *[]Models.User {
	user := &[]Models.User{}

	if err := Config.DB.Find(&user).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func (srv *userService) CreateUser(user *Models.User) (err error) {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()
	user.USER_PASSWORD, err = Config.HashPassword(user.USER_PASSWORD)

	tx := Config.DB.Begin()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *userService) GetUserOne(email string) *Models.User {
	user := &Models.User{}

	if err := Config.DB.Where("user_email = ?", email).First(&user).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func (srv *userService) LoginUser(login *Models.Login) jwt.Token {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()
	user := srv.GetUserOne(login.USER_EMAIL)

	if isLogin := Config.CheckPasswordHash(login.USER_PASSWORD, user.USER_PASSWORD); isLogin == false {
		return nil
	}

	result := Sign(user)

	return result
}

// func (srv *userService) UpdateUser() *Models.User {
// 	Config.DB.Save(user)
// 	return nil
// }

// func DeleteUser(user *User, id string) *Models.User {
// 	Config.DB.Where("id = ?", id).Delete(user)
// 	return nil
// }
