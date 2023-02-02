package models

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	// 비밀번호 공백 처리
	Password []byte `json:"-"`
}
