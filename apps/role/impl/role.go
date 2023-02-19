package impl

import (
	"context"
	"github.com/motongxue/keyauth-g7/apps/role"
)

func (s *service) QueryRole(ctx context.Context, req *role.QueryRoleRequest) (*role.RoleSet, error) {
	query := newQueryRequest(req)
	return s.query(ctx, query)
}

func (s *service) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.Role, error) {
	ins := role.NewRole(req)
	err := s.save(ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}
