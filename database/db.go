package database

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var G_db *gorm.DB

type Account struct {
	gorm.Model
	Username string
	Password string
	Nickname string
	Age      int
	Gender   string
}

type Db struct {
}

func Start() {
	d := Db{}
	d.Connet()
	d.CreateTable()
}

func (d *Db) Connet() {
	db, err := gorm.Open("mysql", "root:mima@/rpc_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
		errors.New("open database failed!")
	}
	G_db = db
}

func (d *Db) CreateTable() {
	if G_db == nil {
		errors.New("G_db is nil!")
		return
	}
	if G_db.HasTable(&Account{}) {
		G_db.AutoMigrate()
	} else {
		G_db.CreateTable(&Account{})
	}
}
