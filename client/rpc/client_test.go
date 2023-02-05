package rpc_test

import (
	"context"
	"fmt"
	"github.com/motongxue/keyauth-g7/apps/token"
	"github.com/motongxue/keyauth-g7/client/rpc"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	mcenter "github.com/infraboard/mcenter/client/rpc"
)

// keyauth 客户端
// 需要配置注册中心的地址
// 获取注册中心的客户端，使用注册中心的客户端 查询 keyauth的地址
func TestBookQuery(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLIENT_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")
	t.Log(conf)
	// 传递Mcenter配置, 客户端通过Mcenter进行搜索
	c, err := rpc.NewClient(conf)
	if should.NoError(err) {
		resp, err := c.Token().ValidateToken(
			context.Background(),
			token.NewValidateTokenRequest("ICjMbuzXJi12Xel7qKxAh9WJ"),
		)
		should.NoError(err)
		fmt.Println(resp)
	}
}
func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := mcenter.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
