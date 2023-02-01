package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Public Function 은 앞에 대문자로 기록 / 소문자는 프라이빗
func Connect() {
	_, err := gorm.Open(mysql.Open("root:test1234@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connet to the database")
	}

}
