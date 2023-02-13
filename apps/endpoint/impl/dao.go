package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/motongxue/keyauth-g7/apps/endpoint"
)

// Save Object
func (s *service) save(ctx context.Context, set *endpoint.EndpointSet) error {
	// s.col.InsertMany()
	if _, err := s.col.InsertMany(ctx, set.ToDocs()); err != nil {
		return exception.NewInternalServerError("inserted service %s endpoint document error, %s",
			set.Service, err)
	}
	return nil
}
