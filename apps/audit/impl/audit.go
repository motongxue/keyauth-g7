package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/motongxue/keyauth-g7/apps/audit"
)

func (s *service) AuditOperate(ctx context.Context, req *audit.OperateLog) (*audit.AuditOperateLogResponse, error) {
	s.log.Debug("operate log %s", req)
	if _, err := s.col.InsertOne(ctx, req); err != nil {
		return nil, exception.NewInternalServerError("inserted audit log(%s) document error, %s",
			req, err)
	}
	return &audit.AuditOperateLogResponse{}, nil
}
