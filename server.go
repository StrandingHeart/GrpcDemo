package main

import (
	"context"
	"flag"
	service "github.com/StrandingHeart/GrpcDemo/service"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	user "github.com/StrandingHeart/GrpcDemo/grpc/user"
)

const (
	port = ":8008"
)

type UserServiceImpl struct {
	// 实现 User 服务的业务对象，可以理解为接口的实现类，serviceImpl
}

//往MongoDB的test库的user集合加一条数据
func (userService *UserServiceImpl) UserInsert(ctx context.Context, in *user.UserInsertRequest) (*user.UserInsertResponse, error) {
	err := service.InsertUser(ctx, in)
	if err != nil {
		return &user.UserInsertResponse{
			Err: 500,
			Msg: err.Error(),
		}, err
	}
	return &user.UserInsertResponse{
		Err: 0,
		Msg: `success`,
	}, err
}

func (userService *UserServiceImpl) UserId(ctx context.Context, in *user.UserIdRequest) (*user.UserIdResponse, error) {
	glog.Infoln(`receive UserId API : `, in.Id, in.String())
	return &user.UserIdResponse{
		Err: 0,
		Msg: "this is UserId response",
		Data: &user.UserEntity{
			Name: "zy", Age: 23, Sex: 0, Hobby: []string{`LOL`, `ball`},
		},
	}, nil
}

// UserService 实现了 User 服务接口中声明的所有方法,在user.RegisterUserServer(grpcServer, &UserService{})注册时必须实现API
func (userService *UserServiceImpl) UserIndex(ctx context.Context, in *user.UserIndexRequest) (*user.UserIndexResponse, error) {
	glog.Infoln(`receive user index API : `, in.Page, in.PageSize, in.String())

	return &user.UserIndexResponse{
		Err: 0,
		Msg: "success",
		Data: []*user.UserEntity{
			{Name: "zy", Age: 23, Sex: 0, Hobby: []string{`LOL`, `ball`}},
			{Name: "cjf", Age: 18, Sex: 1, Hobby: []string{`LOL`, `buy`}},
		},
	}, nil
}

func (userService *UserServiceImpl) UserDelete(ctx context.Context, in *user.UserDeleteRequest) (*user.UserDeleteResponse, error) {
	glog.Infoln(`receive user delete request: `, in.Id)

	return &user.UserDeleteResponse{
		Err: 0,
		Msg: "delete success",
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalln("failed to listen: ", err)
	}

	// 创建RPC服务server
	grpcServer := grpc.NewServer()

	// 注册user服务到grpc
	user.RegisterUserServer(grpcServer, &UserServiceImpl{})
	// 注册反射服务 给调用方CLI(例如grpcurl)准备的 跟服务本身没有关系
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		glog.Fatalln("failed to serve: ", err)
	}
}
