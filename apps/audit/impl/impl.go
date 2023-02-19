package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/motongxue/keyauth-g7/apps/audit"
	"github.com/motongxue/keyauth-g7/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	audit.Service
}

func (s *service) Name() string {
	return audit.AppName
}

func (s *service) Registry(server *grpc.Server) {
	audit.RegisterRPCServer(server, svr)
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	return nil
}

func init() {
	app.RegistryGrpcApp(svr)

}
