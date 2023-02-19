package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/motongxue/keyauth-g7/apps/policy"
	"github.com/motongxue/keyauth-g7/apps/role"
	"github.com/motongxue/keyauth-g7/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svr = &service{}
)

type service struct {
	log  logger.Logger
	col  *mongo.Collection
	role role.Service
	policy.UnimplementedRPCServer
}

func (s *service) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	s.role = app.GetInternalApp(role.AppName).(role.Service)
	return nil

}
func (s *service) Name() string {
	return policy.AppName
}

func (s *service) Registry(server *grpc.Server) {
	policy.RegisterRPCServer(server, svr)
}
func init() {
	app.RegistryGrpcApp(svr)
	app.RegistryInternalApp(svr)
}
