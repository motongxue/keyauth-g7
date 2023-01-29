package impl

import (
	"context"
	"fmt"
	"github.com/motongxue/keyauth-g7/apps/token"
	"github.com/motongxue/keyauth-g7/apps/user"
	"time"

	"github.com/infraboard/mcube/exception"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	AUTH_ERROR = "user or password not correct"
)

func (i *impl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate issue token error, %s", err)
	}

	// 根据不同授权模型来做不同的验证
	switch req.GrantedType {
	case token.GrantedType_PASSWORD:
		// 1. 获取用户对象(User Object)
		descReq := user.NewDescribeUserRequestByName(req.UserDomain, req.UserName)
		u, err := i.user.DescribeUser(ctx, descReq)
		if err != nil {
			i.log.Debug("describe user error, %s", err)
			if exception.IsNotFoundError(err) {
				// 401
				return nil, exception.NewUnauthorized(AUTH_ERROR)
			}
			return nil, err
		}

		// 2. 校验用户密码是否正确
		i.log.Debug(u)
		if ok := u.CheckPassword(req.Password); !ok {
			// 401
			return nil, exception.NewUnauthorized(AUTH_ERROR)
		}

		// 3. 颁发一个Token, 颁发<JWT> xxx  Sign(url+ body) Sing-->Heander -->  Hash
		// 4. rfc: Bearer  字符串:   Header: Authorization  Header Value: bearer <access_token>
		tk := token.NewToken(req, 10*time.Minute)
		return tk, nil
	default:
		return nil, fmt.Errorf("grant type %s not implemented", req.GrantedType)
	}
}

func (i *impl) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (*token.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevolkToken not implemented")
}

func (i *impl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func (i *impl) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.TokenSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryToken not implemented")
}
