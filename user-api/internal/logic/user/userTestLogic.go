package user

import (
	"context"
	"fmt"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTestLogic {
	return &UserTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTestLogic) UserTest(req *types.UserTestReq) (resp *types.UserTestResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("user test ing")
	return
}
