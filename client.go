package main

import (
	"context"
	"github.com/golang/glog"
	"time"

	"google.golang.org/grpc"

	user "github.com/StrandingHeart/GrpcDemo/grpc/user"
)

const (
	address = "127.0.0.1:8008"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		glog.Fatalln("did not connect: ", err)
	}

	defer conn.Close()

	//创建RPC user服务的客户端
	userClient := user.NewUserClient(conn)

	// 设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// UserIndex 请求，普通的请求，没有数据库连接
	userIndexResponse, err := userClient.UserIndex(ctx, &user.UserIndexRequest{Page: 1, PageSize: 2})
	if err != nil {
		glog.Errorln(`user index API error :`, err)
		return
	}
	for _, datum := range userIndexResponse.Data {
		glog.Infoln(`success: `, datum.String())
	}

	// UserDelete 请求
	userDeleteReponse, err := userClient.UserDelete(ctx, &user.UserDeleteRequest{Id: 1})
	if err != nil {
		glog.Errorln(`user delete API error :`, err)
		return
	} else {
		glog.Infoln(`success: `, userDeleteReponse.Msg)
	}

	//userInsert这个demo的API连了MongoDB，需要在server上配置MongoDB的连接URL
	userInsertRes, err := userClient.UserInsert(ctx, &user.UserInsertRequest{Data: &user.UserEntity{Name: "zy", Age: 23, Sex: 0, Hobby: []string{`LOL`, `ball`}}})
	if err != nil {
		glog.Errorln(`user insert API error : `, err)
		return
	} else {
		glog.Infoln(`success: `, userInsertRes.Msg)
	}

	userIdResponse, err := userClient.UserId(ctx, &user.UserIdRequest{Id: 1})
	if err != nil {
		glog.Errorln(`user id API error :`, err)
		return
	} else {
		glog.Infoln(`success: `, userIdResponse.Data.String())
	}

}
