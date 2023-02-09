package auth

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/motongxue/keyauth-g7/apps/token"
	"github.com/motongxue/keyauth-g7/common/utils"
)

// Go-Restful auth Middleware
// 通过 r.Filter(<middleware>) 添加中间件
func (a *KeyauthAuther) RestfulAuthHandlerFunc(
	req *restful.Request,
	resp *restful.Response,
	chain *restful.FilterChain,
) {
	// 1. 能不能获取路由装饰信息
	meta := req.SelectedRoute().Metadata()
	a.log.Debug(meta)

	// 获取meta信息, get, 判断是否开启认证
	var isAuthEnable bool
	if authV, ok := meta[label.Auth]; ok {
		switch v := authV.(type) {
		case bool:
			isAuthEnable = v
		case string:
			isAuthEnable = v == "true"
		}
	}
	if isAuthEnable {
		// 处理Request对象
		// *restful.Request

		// 1. 认证中间件, 需要获取Token
		tkStr := utils.GetToken(req.Request)

		// 2. 到用户中心验证token合法性, 依赖用户中心的Grpc Client
		validateReq := token.NewValidateTokenRequest(tkStr)
		tk, err := a.auth.ValidateToken(req.Request.Context(), validateReq)
		if err != nil {
			response.Failed(resp.ResponseWriter, err)
			return
		}
		// 处理完成后 需要直接中断返回
		// *restful.Response
		// resp.Header()
		// resp.Write()
		// 3. 处理完成后，需要直接中断返回
		// req.SetAttribute("token", tk)
		req.SetAttribute("token", tk)
		// 4. chain用于将请求flow下去
		chain.ProcessFilter(req, resp)
	}
}
