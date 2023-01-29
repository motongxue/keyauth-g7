package all

import (
	_ "github.com/motongxue/keyauth-g7/apps/token/api"
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	// _ "github.com/motongxue/keyauth-g7/apps/book/api"
	_ "github.com/motongxue/keyauth-g7/apps/user/api"
)
