package token

import (
	"fmt"
	"github.com/infraboard/mcube/exception"
	"github.com/motongxue/keyauth-g7/apps/user"
	"github.com/motongxue/keyauth-g7/common/utils"
	"time"
)

const (
	AppName = "token"
)

func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{
		UserDomain: user.DefaultDomain,
	}
}

func (req *IssueTokenRequest) Validate() error {
	switch req.GrantedType {
	case GrantedType_PASSWORD:
		if req.UserName == "" || req.Password == "" {
			return fmt.Errorf("password grant required username and password")
		}
	}

	return nil
}

func NewDefaultToken() *Token {
	return &Token{
		Data: &IssueTokenRequest{},
		Meta: map[string]string{},
	}
}

func NewToken(req *IssueTokenRequest, expiredDuration time.Duration) *Token {
	now := time.Now()
	// Token 10
	expired := now.Add(expiredDuration)
	refresh := now.Add(expiredDuration * 5)

	return &Token{
		AccessToken:           utils.MakeBearer(24),
		IssueAt:               now.UnixMilli(),
		Data:                  req,
		AccessTokenExpiredAt:  expired.UnixMilli(),
		RefreshToken:          utils.MakeBearer(32),
		RefreshTokenExpiredAt: refresh.UnixMilli(),
	}
}

func (t *Token) Validate() error {
	// 判断Token过期没有
	// 是一个时间戳,
	//  now   expire
	if time.Now().UnixMilli() > t.AccessTokenExpiredAt {
		return exception.NewAccessTokenExpired("access token expired")
	}

	return nil
}

func (t *Token) IsRefreshTokenExpired() bool {
	// 判断refresh Token过期没有
	// 是一个时间戳,
	//  now   expire
	if time.Now().UnixMilli() > t.RefreshTokenExpiredAt {
		return true
	}

	return false
}

// 续约Token, 延长一个生命周期
func (t *Token) Extend(expiredDuration time.Duration) {
	now := time.Now()
	// Token 10
	expired := now.Add(expiredDuration)
	refresh := now.Add(expiredDuration * 5)

	t.AccessTokenExpiredAt = expired.UnixMilli()
	t.RefreshTokenExpiredAt = refresh.UnixMilli()
}

func NewDescribeTokenRequest(at string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		AccessToken: at,
	}
}

func NewValidateTokenRequest(at string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: at,
	}
}

func NewRevolkTokenRequest() *RevolkTokenRequest {
	return &RevolkTokenRequest{}
}
