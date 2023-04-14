package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/motongxue/keyauth-g7/common/utils"
	"github.com/rs/xid"
	"net/http"
	"time"
)

const (
	AppName       = "user"
	DefaultDomain = "default"
)

var (
	validate = validator.New()
)

func NewUser(req *CreateUserRequest) *User {
	return &User{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Data:     req,
	}
}

func NewDefaultUser() *User {
	return NewUser(NewCreateUserRequest())
}

func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Data.Password)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}
// 原先是 s UserSet 改成了 -> s *UserSet，是否因为这个存在bug
func (s *UserSet) Add(item *User) {
	s.Items = append(s.Items, item)
}
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Domain: DefaultDomain,
	}
}
func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	query := r.URL.Query()
	return &QueryUserRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: query.Get("keywords"),
	}
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DescribeBy_USER_ID,
		UserId:     id,
	}
}

func NewDescribeUserRequestByName(domain, name string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DescribeBy_USER_NAME,
		Domain:     domain,
		UserName:   name,
	}
}

func NewPutUserRequest(id string) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMilli(),
		Data:       NewCreateUserRequest(),
	}
}

func NewPatchUserRequest(id string) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMilli(),
		Data:       NewCreateUserRequest(),
	}
}

func NewDeleteUserRequestWithID(id string) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: id,
	}
}
