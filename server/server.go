package main

import (
	"RedRock-web-back-end-2020-7-lv1/account"
	"RedRock-web-back-end-2020-7-lv1/database"
	"context"
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

}

func (s *server) modifyInfo(ctx context.Context, info *account.Info) (*account.StatusWithInfo, error) {

}

func (s *server) getInfo(ctx context.Context, data account.NullData) (*account.Info, error) {

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
