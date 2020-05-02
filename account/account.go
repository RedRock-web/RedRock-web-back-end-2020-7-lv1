package account

import (
	"RedRock-web-back-end-2020-7-lv1/database"
	"errors"
	"fmt"
)

func Isregistered(username string) bool {
	var a database.Account

	if database.G_db == nil {
		fmt.Println("G_db is nil!")
	}

	database.G_db.Where("username = ?", username).Find(&a)

	return a.Password != ""
}

func PasswdIsOk(passwd string) bool {
	var a database.Account

	if err := database.G_db.Where("password = ?", passwd).Find(&a).Error; err != nil {
		fmt.Println(err)
		errors.New("judge password if failed!")
	}

	return a.Password == passwd
}

func GetInfo(username string) *Info {
	var account database.Account

	if err := database.G_db.Where("username = ?", username).Find(&account).Error; err != nil {
		fmt.Println(err)
	}

	info := &Info{
		Username: account.Username,
		Password: account.Password,
		Nickname: account.Nickname,
		Age:      account.Age,
		Gender:   account.Gender,
	}

	return info
}
