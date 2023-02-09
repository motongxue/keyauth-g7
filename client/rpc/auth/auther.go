package auth

import (
	"github.com/infraboard/mcube/logger"
	"github.com/motongxue/keyauth-g7/apps/token"
)

func NewKeyauthAuther(auth token.ServiceClient) *KeyauthAuther {
	return &KeyauthAuther{
		auth: auth,
	}
}

type KeyauthAuther struct {
	log  logger.Logger
	auth token.ServiceClient
}
