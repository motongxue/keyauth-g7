package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/motongxue/keyauth-g7/apps/endpoint"
	"github.com/motongxue/keyauth-g7/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	endpoint.UnimplementedRPCServer
}

func (s *service) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	// 获取一个Collection对象, 通过Collection对象 来进行CRUD
	s.col = db.Collection(s.Name())

	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return endpoint.AppName
}

func (s *service) Registry(server *grpc.Server) {
	endpoint.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
	// 将扩展接口后的实例注册到内部类中，不暴露到RPC服务接口中
	app.RegistryInternalApp(svr)
}
