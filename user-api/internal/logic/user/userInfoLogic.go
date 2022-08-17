package user

import (
	"context"
	"github.com/pkg/errors"
	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		UserId:   userResp.UserModel.Id,
		NickName: userResp.UserModel.Nickname,
	}, nil
}

func (l *UserInfoLogic) testOne() error {
	return l.testTwo()
}
func (l *UserInfoLogic) testTwo() error {
	return l.testThree()
}
func (l *UserInfoLogic) testThree() error {
	return errors.Wrap(errors.New("这是故意的"), "哈哈哈")
}
