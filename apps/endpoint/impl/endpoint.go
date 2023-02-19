package impl

import (
	"context"
	"github.com/motongxue/keyauth-g7/apps/endpoint"
)

// RegistryEndpoint 用户和角色都注册到系统中，
// 需要将服务的功能注册到权限中心
func (s *service) RegistryEndpoint(ctx context.Context, req *endpoint.EndpointSet) (
	*endpoint.RegistryResponse, error) {
	if err := s.save(ctx, req); err != nil {
		return nil, err
	}
	resp := endpoint.NewRegistryResponse()
	return resp, nil
}
