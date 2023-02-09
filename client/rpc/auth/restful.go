package auth

import (
	"github.com/emicklei/go-restful/v3"
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
	// 处理Request对象 *restful.Request
	// 1. 认证中间件，需要获取Toekn
	tkStr := utils.GetToken(req.Request)
	// 2. 到用户中心验证token合法性，依赖用户中心的Grpc client
	validateReq := token.NewValidateTokenRequest(tkStr)
	tk, err := a.auth.ValidateToken(req.Request.Context(), validateReq)
	if err != nil {
		response.Failed(resp.ResponseWriter, err)
		return
	}
	// 3. 处理完成后，需要直接中断返回
	req.SetAttribute("token", tk)
	// 4. chain用于将请求flow下去
	chain.ProcessFilter(req, resp)
}
