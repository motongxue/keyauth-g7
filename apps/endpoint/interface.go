package endpoint

const (
	AppName = "endpoint"
)

type Service interface {
	// RPCServer 若只想将一部分接口通过RPC暴露出去，则仅将需暴露接口写在proto中的rpc服务中，
	// 新定义interface，并将其他不想暴露的接口则继承RPCServer接口
	// 将不需要暴露的接口添加在这里，并自己实现和注册
	// app.RegistryInternalApp(svr)
	RPCServer
}

func (s *EndpointSet) ToDocs() (docs []interface{}) {
	for i := range s.Endpoints {
		docs = append(docs, s.Endpoints[i])
	}
	return
}

func NewRegistryResponse() *RegistryResponse {
	return &RegistryResponse{}
}
