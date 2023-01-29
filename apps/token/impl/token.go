package impl

import (
	"context"
	"fmt"
	"github.com/motongxue/keyauth-g7/apps/token"
	"github.com/motongxue/keyauth-g7/apps/user"
	"github.com/motongxue/keyauth-g7/common/utils"
	"time"

	"github.com/infraboard/mcube/exception"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	AUTH_ERROR           = "user or password not correct"
	DefaultTokenDuration = 10 * time.Minute
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
		tk := token.NewToken(req, DefaultTokenDuration)

		// 5. 脱敏
		tk.Data.Password = ""

		// 6. 入库持久化
		if err := i.save(ctx, tk); err != nil {
			return nil, err
		}
		return tk, nil
	default:
		return nil, fmt.Errorf("grant type %s not implemented", req.GrantedType)
	}
}

func (i *impl) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (*token.Token, error) {
	// 1. 获取AccessToken
	tk, err := i.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 2. 检查RefreshToken是否匹配
	if tk.RefreshToken != req.RefreshToken {
		return nil, exception.NewBadRequest("refresh token not conrrect")
	}

	// 3. 删除
	if err := i.delete(ctx, tk); err != nil {
		return nil, err
	}
	return tk, nil
}

func (i *impl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	// 1. 获取AccessToken
	tk, err := i.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 2. 校验Token合法性
	if tk.Validate(); err != nil {
		// 2.1 如果Aceess Token过期
		if utils.IsAccessTokenExpiredError(err) {
			if tk.IsRefreshTokenExpired() {
				return nil, exception.NewRefreshTokenExpired("refresh token expired")
			}
			// 2.2 如果Refresh没过期, 可以其过期时间
			// 类似于执行了一个Update, Update Exired 时间
			tk.Extend(DefaultTokenDuration)
			if err := i.update(ctx, tk); err != nil {
				return nil, err
			}

			// 返回续约后的Token
			return tk, nil
		}
		return nil, err
	}

	return tk, nil
}

func (i *impl) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.TokenSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryToken not implemented")
}
