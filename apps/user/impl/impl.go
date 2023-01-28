package impl

import (
	"context"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/motongxue/keyauth-g7/apps/user"
	"github.com/motongxue/keyauth-g7/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	user.UnimplementedServiceServer
}

func (i *impl) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	// 创建索引
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{
					Key: "data.domain", Value: bsonx.Int32(-1),
				},
				{
					Key: "data.name", Value: bsonx.Int32(-1),
				},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	if _, err := i.col.Indexes().CreateMany(context.Background(), indexs); err != nil {
		return err
	}
	return nil
}

func (i *impl) Name() string {
	return user.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	user.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
