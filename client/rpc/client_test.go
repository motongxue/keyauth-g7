package rpc_test

import (
	"context"
	"fmt"
	"github.com/motongxue/keyauth-g7/apps/policy"
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
func TestTokenQuery(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLIENT_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")
	t.Log(conf)
	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端
	keyauthClient, err := rpc.NewClient(conf)

	// 进行服务功能注册
	// keyauthClient.Endpoint().RegistryEndpoint()

	if should.NoError(err) {
		resp, err := keyauthClient.Token().ValidateToken(
			context.Background(),
			token.NewValidateTokenRequest("ICjMbuzXJi12Xel7qKxAh9WJ"),
		)
		should.NoError(err)
		fmt.Println(resp)
	}
}

// 测试鉴权
func TestValidatePermission(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLINET_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")

	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端
	keyauthClient, err := rpc.NewClient(conf)
	if should.NoError(err) {
		req := policy.NewValidatePermissionRequest()
		req.Username = "member"
		req.Service = "cmdb"
		req.Resource = "secret"
		// req.Action = "delete" Mongdb中，没有对应的delete操作，故执行应该失败
		req.Action = "list"
		// TODO 查明为什么会出现no permission？
		//  因为policy中是通过policy表中的spec.role(而非spec.username)该字符串作为标识去role表中查找对应的spec.name标识
		p, err := keyauthClient.Policy().ValidatePermission(context.TODO(), req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(p)
	}
}

func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := mcenter.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
