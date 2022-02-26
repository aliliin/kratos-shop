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
	//TestDeleteAddress() // 删除用户地址
	//TestDefaultAddress() // 设置用户默认地址
	//TestUpdateAddress() // 修改用户地址列表
	TestGetUserAddressList() // 获取用户地址列表
	//TestCreateAddress() // 创建用户地址
	//TestGetUserList() // 获取用户列表
	//TestCreateUser() // 创建用户
	//TestUpdateUser()      // 更新用户
	//TestGetUserByMobile() // 根据手机获取用户
	//TestGetUserById() // 根据ID 获取用户
	conn.Close()
}
func TestDeleteAddress() {
	rsp, err := userClient.DeleteAddress(context.Background(), &v1.AddressReq{
		Id:  4,
		Uid: 2,
	})
	if err != nil {
		panic("grpc 删除用户地址失败" + err.Error())
	}
	fmt.Println(rsp)
}
func TestDefaultAddress() {
	rsp, err := userClient.DefaultAddress(context.Background(), &v1.AddressReq{
		Id:  2,
		Uid: 1,
	})
	if err != nil {
		panic("grpc 设置默认地址失败" + err.Error())
	}
	fmt.Println(rsp)
}
func TestUpdateAddress() {
	rsp, err := userClient.UpdateAddress(context.Background(), &v1.UpdateAddressReq{
		Id:        1,
		Uid:       2,
		Name:      "test1111",
		Mobile:    "13161006666",
		Province:  "北京市",
		City:      "北京",
		Districts: "朝阳",
		Address:   "十八里",
		PostCode:  "00001",
		IsDefault: 0,
	})
	if err != nil {
		panic("grpc 修改地址失败" + err.Error())
	}
	fmt.Println(rsp)
}
func TestCreateAddress() {
	rsp, err := userClient.CreateAddress(context.Background(), &v1.CreateAddressReq{
		Uid:       2,
		Name:      "test666",
		Mobile:    "13161006666",
		Province:  "北京市",
		City:      "北京",
		Districts: "朝阳",
		Address:   "十八里",
		PostCode:  "00001",
		IsDefault: 0,
	})
	if err != nil {
		panic("grpc 创建地址失败" + err.Error())
	}
	fmt.Println(rsp.Id)
}

func TestGetUserAddressList() {
	rsp, err := userClient.ListAddress(context.Background(), &v1.ListAddressReq{
		Uid: 2,
	})
	if err != nil {
		panic("grpc get user by ID err" + err.Error())
	}
	fmt.Println(rsp)
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
			Password: "Gaofei123456",
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
		PSize: 60,
	})

	if err != nil {
		panic("grpc get err" + err.Error())
	}

	for _, user := range r.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)

		if user.Mobile == "13501167242" {
			checkRsp, err := userClient.CheckPassword(context.Background(), &v1.PasswordCheckInfo{
				Password:          "1234567890",
				EncryptedPassword: user.Password,
			})
			if err != nil {
				panic(" get check user  psw err" + err.Error())
			}
			fmt.Println(checkRsp.Success)
		} else {
			checkRsp, err := userClient.CheckPassword(context.Background(), &v1.PasswordCheckInfo{
				Password:          "admin",
				EncryptedPassword: user.Password,
			})
			if err != nil {
				panic(" get check user  psw err" + err.Error())
			}
			fmt.Println(checkRsp.Success)
		}

	}
	fmt.Println(r.Total)
}
