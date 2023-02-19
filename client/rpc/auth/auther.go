package auth

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/motongxue/keyauth-g7/apps/policy"
	"github.com/motongxue/keyauth-g7/apps/token"
	"github.com/motongxue/keyauth-g7/client/rpc"
)

func NewKeyauthAuther(client *rpc.ClientSet, serviceName string) *KeyauthAuther {
	return &KeyauthAuther{
		log:         zap.L().Named("http.auther"),
		auth:        client.Token(),
		perm:        client.Policy(),
		serviceName: serviceName,
	}
}

// 有Keyauth提供的 HTTP认证中间件
type KeyauthAuther struct {
	log         logger.Logger
	auth        token.ServiceClient
	perm        policy.RPCClient
	serviceName string
}
