package main

import (
	"RedRock-web-back-end-2020-7-lv1/account"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	//address     = "47.98.57.152:1234"
	address = "127.0.0.1:1234"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	a := account.NewServerClient(conn)
	RegisterTest(a)
	LoginTest(a)
	ModifyInfoTest(a)
	GetInfoTest(a)
}

func RegisterTest(a account.ServerClient) {
	// 初次注册
	r, err := a.Register(context.Background(), &account.Account{
		Username: "root",
		Password: "mima",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)

	// 注册已经注册的
	r, err = a.Register(context.Background(), &account.Account{
		Username: "root",
		Password: "mima",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)
}

func LoginTest(a account.ServerClient) {
	// 用于判断没有注册的帐号
	l, err := a.Login(context.Background(), &account.Account{
		Username: "haha",
		Password: "haha",
	})
	log.Fatalln(err)
	fmt.Println(l)

	// 用于判断密码错误的帐号
	l, err = a.Login(context.Background(), &account.Account{
		Username: "root",
		Password: "haha",
	})
	log.Fatalln(err)
	fmt.Println(l)

	// 用于正确登录帐号
	l, err = a.Login(context.Background(), &account.Account{
		Username: "root",
		Password: "mima",
	})
	log.Fatalln(err)
	fmt.Println(l)

}

func GetInfoTest(a account.ServerClient) {
	g, err := a.GetInfo(context.Background(), &account.Username{
		Username: "root",
	})
	log.Fatalln(err)
	fmt.Println(g)
}

func ModifyInfoTest(a account.ServerClient) {
	m, err := a.ModifyInfo(context.Background(), &account.Info{
		Username: "root",
		Password: "mima",
		Nickname: "hah",
		Age:      20,
		Gender:   "male",
	})
	log.Fatalln(err)
	fmt.Println(m)
}
