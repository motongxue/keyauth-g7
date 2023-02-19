package role_test

import (
	"github.com/motongxue/keyauth-g7/apps/role"
	"testing"
)

func TestHasPermission(t *testing.T) {
	set := role.NewRoleSet()
	r := &role.Role{
		Spec: &role.CreateRoleRequest{
			Permissions: []*role.Permission{
				{
					// 设置指定服务权限
					// Service: "cmdb",
					// Features: []*role.Feature{
					// 	{
					// 		Resource: "secret",
					// 		Action:   "list",
					// 	},
					// 	{
					// 		Resource: "secret",
					// 		Action:   "get",
					// 	},
					// 	{
					// 		Resource: "secret",
					// 		Action:   "create",
					// 	},
					// },

					// 设置AllowAll 服务权限
					Service:  "cmdb",
					AllowAll: true,
				},
			},
		},
	}
	set.Add(r)
	perm, testRole := set.HasPermission(&role.PermissionRequest{
		Service:  "cmdb",
		Resource: "secret",
		// Action:   "delete", // 无delete方法，测试为失败
		Action: "create", // 有create、list、get方法，故测试成功
	})
	t.Log(testRole)
	if perm != true {
		t.Fatal("has perm error")
	}
}
