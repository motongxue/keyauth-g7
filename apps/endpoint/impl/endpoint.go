package impl

import (
	"context"
	"github.com/motongxue/keyauth-g7/apps/endpoint"
)

func (s *service) RegistryEndpoint(ctx context.Context, req *endpoint.EndpointSet) (
	*endpoint.RegistryResponse, error) {
	if err := s.save(ctx, req); err != nil {
		return nil, err
	}
	resp := endpoint.NewRegistryResponse()
	return resp, nil
}
