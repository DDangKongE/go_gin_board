package Models

import "go_board/src/Config"

type User struct {
	USER_ID       uint   `gorm:"primaryKey"`
	USER_EMAIL    string `json:"user_email"`
	USER_PASSWORD string `json:"user_password"`
	USER_NICKNAME string `json:"user_nickname"`
}

func GetUser(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserOne(user *User, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
