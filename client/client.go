package main

import (
	"RedRock-web-back-end-2020-7-lv1/account"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	//remoteAddress     = "47.98.57.152:1234"
	localAddress = "127.0.0.1:1234"
)

func main() {
	conn, err := grpc.Dial(localAddress, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		errors.New("1")
		log.Fatalln(err)
	}
	defer conn.Close()

	client := account.NewServerClient(conn)

	RegisterTest(client)
	LoginTest(client)
	GetInfoTest(client)
	ModifyInfoTest(client)
}

func RegisterTest(client account.ServerClient) {
	// 初次注册
	r, err := client.Register(context.Background(), &account.Account{
		Username: "root",
		Password: "mima",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("尝试第一次初次注册：")
	fmt.Println(r)

	// 注册已经注册的
	r, err = client.Register(context.Background(), &account.Account{
		Username: "root",
		Password: "mima",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("尝试第二次注册：")
	fmt.Println(r)
}

func LoginTest(client account.ServerClient) {
	// 用于判断没有注册的帐号
	l, _ := client.Login(context.Background(), &account.Account{
		Username: "haha",
		Password: "haha",
	})
	fmt.Println("用于判断没有注册的帐号")
	fmt.Println(l)

	// 用于判断密码错误的帐号
	l, _ = client.Login(context.Background(), &account.Account{
		Username: "root",
		Password: "haha",
	})
	fmt.Println("用于判断密码错误的帐号")
	fmt.Println(l)

	// 用于正确登录帐号
	l, _ = client.Login(context.Background(), &account.Account{
		Username: "root",
		Password: "mima",
	})
	fmt.Println(" 用于正确登录帐号")
	fmt.Println(l)

}

func GetInfoTest(client account.ServerClient) {
	g, err := client.GetInfo(context.Background(), &account.Username{
		Username: "root",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("获取信息:")
	fmt.Println(g)
}

func ModifyInfoTest(client account.ServerClient) {
	m, _ := client.ModifyInfo(context.Background(), &account.Info{
		Username: "root",
		Password: "mima",
		Nickname: "hah",
		Age:      20,
		Gender:   "male",
	})
	fmt.Println("修改信息：")
	fmt.Println(m)
}
