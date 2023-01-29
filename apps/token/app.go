package token

import (
	"fmt"
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
