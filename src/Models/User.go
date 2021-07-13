package Models

type User struct {
	USER_ID       uint   `gorm:"primaryKey"`
	USER_EMAIL    string `json:"user_email"`
	USER_PASSWORD string `json:"user_password"`
	USER_NICKNAME string `json:"user_nickname"`
	Boards        []Board
}

type Login struct {
	USER_EMAIL    string `json:"user_email"`
	USER_PASSWORD string `json:"user_password"`
}

type Token struct {
	Access_token string
	User         User
}
