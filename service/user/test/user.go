package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	v1 "user/api/user/v1"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func main() {
	Init()
	TestGetUserList() // 获取用户列表
	//TestCreateUser() // 创建用户
	TestUpdateUser()      // 更新用户
	TestGetUserByMobile() // 根据手机获取用户
	//TestGetUserById() // 根据ID 获取用户
	conn.Close()
}

// Init 初始化 grpc 链接
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

func TestGetUserById() {
	rsp, err := userClient.GetUserById(context.Background(), &v1.IdRequest{
		Id: 3,
	})
	if err != nil {
		panic("grpc get user by ID err" + err.Error())
	}
	fmt.Println(rsp)
}

func TestGetUserByMobile() {
	rsp, err := userClient.GetUserByMobile(context.Background(), &v1.MobileRequest{
		Mobile: "13501167232",
	})
	if err != nil {
		panic("grpc get user by mobile err" + err.Error())
	}
	fmt.Println(rsp)
}

func TestUpdateUser() {
	rsp, err := userClient.UpdateUser(context.Background(), &v1.UpdateUserInfo{
		Id:       9,
		Gender:   "female",
		NickName: fmt.Sprintf("YWW%d", 233),
	})
	if err != nil {
		panic("grpc update user err" + err.Error())
	}
	fmt.Println(rsp)
}

// TestCreateUser 测试创建 10 个用户数据
func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &v1.CreateUserInfo{
			Mobile:   fmt.Sprintf("1350116723%d", i),
			Password: "admin",
			NickName: fmt.Sprintf("YW%d", i),
		})
		if err != nil {
			panic("grpc 创建用户失败" + err.Error())
		}
		fmt.Println(rsp.Id)
	}
}

func TestGetUserList() {
	r, err := userClient.GetUserList(context.Background(), &v1.PageInfo{
		Pn:    1,
		PSize: 6,
	})

	if err != nil {
		panic("grpc get err" + err.Error())
	}

	for _, user := range r.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		//checkRsp, err := userClient.CheckPassword(context.Background(), &v1.PasswordCheckInfo{
		//	Password:          "admin",
		//	EncryptedPassword: user.Password,
		//})
		//if err != nil {
		//	panic(" get check user  psw err" + err.Error())
		//}
		//fmt.Println(checkRsp.Success)
	}
	fmt.Println(r.Total)
}
