package account

import (
	"RedRock-web-back-end-2020-7-lv1/database"
	"errors"
	"log"
)

func Isregistered(username string) bool {
	var a database.Account

	if err := database.G_db.Where("username = ?", username).Find(&a).Error; err != nil {
		log.Fatalln(err)
		errors.New("judge account if registered is failed!")
	}

	return a.Password == ""
}

func PasswdIsOk(passwd string) bool {
	var a database.Account

	if err := database.G_db.Where("password = ?", passwd).Find(&a).Error; err != nil {
		log.Fatalln(err)
		errors.New("judge password if failed!")
	}

	return a.Password == passwd
}

func GetInfo(username string) (info *Info) {
	if err := database.G_db.Where("username = ?", username).Find(&info).Error; err != nil {
		log.Fatalln(err)
		errors.New("failed get info!")
	}
	return info
}
