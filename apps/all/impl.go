package all

import (
	_ "github.com/motongxue/keyauth-g7/apps/endpoint/impl"
	_ "github.com/motongxue/keyauth-g7/apps/policy/impl"
	_ "github.com/motongxue/keyauth-g7/apps/role/impl"
	_ "github.com/motongxue/keyauth-g7/apps/token/impl"
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	// _ "github.com/motongxue/keyauth-g7/apps/book/impl"
	_ "github.com/motongxue/keyauth-g7/apps/user/impl"
)
