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
					Service: "cmdb",
					Features: []*role.Feature{
						{
							Resource: "secret",
							Action:   "list",
						},
						{
							Resource: "secret",
							Action:   "get",
						},
						{
							Resource: "secret",
							Action:   "create",
						},
					},
				},
			},
		},
	}
	set.Add(r)
	perm, testRole := set.HasPermission(&role.PermissionRequest{
		Service:  "cmdb",
		Resource: "secret",
		// Action:   "delete",  无delete方法，测试为失败
		Action: "create",
	})
	t.Log(testRole)
	if perm != true {
		t.Fatal("has perm error")
	}
}
