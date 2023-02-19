package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/motongxue/keyauth-g7/apps/policy"
	"github.com/motongxue/keyauth-g7/apps/role"
)

func (s *service) ValidatePermission(ctx context.Context, req *policy.ValidatePermissionRequest) (
	*policy.Policy, error) {
	// 根据用户和命名空间找到该用户的授权策略
	// 由于使用分页, 只查询100条数据
	query := policy.NewQueryPolicyRequest()
	query.Namespace = req.Namespace
	query.Username = req.Username
	query.Page.PageSize = 100

	// 从policy表中通过上述查询条件查到PolicySet集合
	set, err := s.QueryPolicy(ctx, query)
	if err != nil {
		return nil, err
	}

	// 从集合中获取用户的角色名
	roleNames := set.Roles()

	s.log.Debugf("found roles: %s", roleNames)
	// 通过角色名获取查询角色列表的请求
	queryRoleReq := role.NewQueryRoleRequestWithName(roleNames)
	queryRoleReq.Page.PageSize = 100
	// 通过角色名查询角色列表
	roles, err := s.role.QueryRole(ctx, queryRoleReq)
	if err != nil {
		return nil, err
	}

	// 判断是否有权限，权限对应的角色
	hasPerm, permissionRole := roles.HasPermission(role.NewPermissionRequest(req.Service, req.Resource, req.Action))
	if !hasPerm {
		return nil, exception.NewPermissionDeny("not permission access service %s resource %s action %s",
			req.Service,
			req.Resource,
			req.Action,
		)
	}

	byRole := set.GetPolicyByRole(permissionRole.Spec.Name)
	return byRole, nil
}

func (s *service) QueryPolicy(ctx context.Context, req *policy.QueryPolicyRequest) (*policy.PolicySet, error) {
	query := newQueryPolicyRequest(req)
	return s.query(ctx, query)
}
func (s *service) CreatePolicy(ctx context.Context, req *policy.CreatePolicyRequest) (
	*policy.Policy, error) {
	ins, err := policy.NewPolicy(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create book error, %s", err)
	}
	if err := s.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, err
}
