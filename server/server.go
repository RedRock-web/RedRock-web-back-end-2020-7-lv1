package main

import (
	"RedRock-web-back-end-2020-7-lv1/account"
	"RedRock-web-back-end-2020-7-lv1/database"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	port = ":1234"
)

type server struct {
	account.UnimplementedServerServer
}

func (s *server) Register(ctx context.Context, a *account.Account) (*account.StatusWithData, error) {
	if account.Isregistered(a.Username) {
		return &account.StatusWithData{
			IsRegistered: "yes",
			Data:         "account had registered!",
		}, nil
	} else {
		err := database.G_db.Create(&database.Account{
			Username: a.Username,
			Password: a.Password,
			Nickname: "user" + strconv.FormatInt(time.Now().Unix(), 10),
			Age:      18,
			Gender:   "male",
		}).Error
		if err != nil {
			log.Fatalln(err)
			return &account.StatusWithData{
				IsRegistered: "yes",
				Data:         "field!!",
			}, nil
		} else {
			return &account.StatusWithData{
				IsRegistered: "no",
				Data:         "successful!",
			}, nil
		}
	}
}

func (s *server) Login(ctx context.Context, a *account.Account) (*account.StatusWithData, error) {
	if account.Isregistered(a.Username) {
		if account.PasswdIsOk(a.Password) {
			return &account.StatusWithData{
				IsRegistered: "yes",
				Data:         "login successful!",
			}, nil
		} else {
			return &account.StatusWithData{
				IsRegistered: "yes",
				Data:         "password is error!",
			}, nil
		}
	} else {
		return &account.StatusWithData{
			IsRegistered: "no",
			Data:         "account not registered!",
		}, nil
	}
}

func (s *server) ModifyInfo(ctx context.Context, info *account.Info) (*account.StatusWithInfo, error) {
	a := database.Account{
		Username: info.Username,
		Password: info.Password,
		Nickname: info.Nickname,
		Age:      info.Age,
		Gender:   info.Gender,
	}

	if err := database.G_db.Model(&a).Where("username = ?", info.Username).Updates(&a).Error; err != nil {
		fmt.Println(err)

		return &account.StatusWithInfo{
			Status: "failed",
			Info:   nil,
		}, nil
	} else {
		return &account.StatusWithInfo{
			Status: "successful!",
			Info:   account.GetInfo(info.Username),
		}, nil
	}
}

func (s *server) GetInfo(ctx context.Context, username *account.Username) (info *account.Info, err error) {
	return account.GetInfo(username.Username), nil
}

func main() {
	database.Start()

	//监听
	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println(err)
	}

	//创建服务端
	s := grpc.NewServer()

	//注册实现服务方法
	account.RegisterServerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
		errors.New("5")
	}
}
