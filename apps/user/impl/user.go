package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"github.com/motongxue/keyauth-g7/apps/user"
	"github.com/motongxue/keyauth-g7/common/utils"
)

func (i *impl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate create user error, %s", err)
	}
	// ins是由创建用户请求而产生的对应的user对象，已经赋值了Name、Password、Domain、CreateBy
	ins := user.NewUser(req)
	// 对密码进行加密处理
	ins.Data.Password = utils.HashPassword(ins.Data.Password)
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted user(%s) document error, %s",
			ins.Data.Name, err)
	}
	return ins, nil
}
func (i *impl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	query := newQueryRequest(req)
	return i.query(ctx, query)
}
func (i *impl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	return i.get(ctx, req)
}
func (i *impl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	ins, err := i.DescribeUser(ctx, user.NewDescribeUserRequestById(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err := ins.Data.Validate(); err != nil {
		return nil, err
	}

	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil

}
func (i *impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.User, error) {
	ins, err := i.DescribeUser(ctx, user.NewDescribeUserRequestById(req.Id))

	if err != nil {
		return nil, err
	}

	if err := i.deleteUser(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
