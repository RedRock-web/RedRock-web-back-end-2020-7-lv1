package main

import (
	"RedRock-web-back-end-2020-7-lv1/account"
	"RedRock-web-back-end-2020-7-lv1/database"
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":1234"
)

type server struct {
	account.UnimplementedServerServer
}

func (s *server) register(ctx context.Context, a *account.Account) (*account.StatusWithData, error) {
	if account.Isregistered(a.Username) {
		return &account.StatusWithData{
			IsRegistered: 1,
			Data:         "account had registered!",
		}, nil
	} else {
		err := database.G_db.Create(&database.Account{
			Username: a.Username,
			Password: a.Password,
			Nickname: "user" + string(time.Now().Unix()),
			Age:      18,
			Gender:   "male",
		}).Error
		if err != nil {
			log.Fatalln(err)
			return &account.StatusWithData{
				IsRegistered: 1,
				Data:         "field!!",
			}, nil
		} else {
			return &account.StatusWithData{
				IsRegistered: 0,
				Data:         "successful!",
			}, nil
		}
	}
}

func (s *server) login(ctx context.Context, a *account.Account) (*account.StatusWithData, error) {
	if account.Isregistered(a.Username) {
		if account.PasswdIsOk(a.Password) {
			return &account.StatusWithData{
				IsRegistered: 0,
				Data:         "login successful!",
			}, nil
		} else {
			return &account.StatusWithData{
				IsRegistered: 1,
				Data:         "password is error!",
			}, nil
		}
	} else {
		return &account.StatusWithData{
			IsRegistered: 1,
			Data:         "account not registered!",
		}, nil
	}
}

func (s *server) modifyInfo(ctx context.Context, info *account.Info) (*account.StatusWithInfo, error) {
	var a account.Account

	modifiedInfo := map[string]interface{}{
		"username": info.Username,
		"passowrd": info.Password,
		"nickname": info.Nickname,
		"gender":   info.Gender,
		"age":      info.Age,
	}
	if err := database.G_db.Model(&a).Where("username = ?", info.Username).Updates(modifiedInfo).Error; err != nil {
		log.Fatalln(err)
		errors.New("failed modify info!")
		return &account.StatusWithInfo{
			Status: "failed",
			Info:   account.GetInfo(info.Username),
		}, nil
	} else {
		return &account.StatusWithInfo{
			Status: "successful!",
			Info:   nil,
		}, nil
	}

}

func (s *server) getInfo(ctx context.Context, username account.Username) (info *account.Info, err error) {
	return account.GetInfo(username.Username), nil
}

func main() {
	//监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	//创建服务端
	s := grpc.NewServer()
	//注册实现服务方法
	account.RegisterServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
