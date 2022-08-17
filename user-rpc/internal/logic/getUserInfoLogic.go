package logic

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("username")
		fmt.Printf("tmp:%+v \n", tmp)
	}
	user, err := l.svcCtx.UserDataModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, errors.New("fffffuck")
	}
	return &pb.GetUserInfoResp{
		UserModel: &pb.UserModel{
			Id:       int64(user.Id),
			Nickname: "ooooooooo",
		},
	}, nil
}
