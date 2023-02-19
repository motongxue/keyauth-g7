package auth

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/motongxue/keyauth-g7/apps/audit"
	"github.com/motongxue/keyauth-g7/apps/policy"
	"github.com/motongxue/keyauth-g7/apps/token"
	"github.com/motongxue/keyauth-g7/common/utils"
	"time"
)

// Go-Restful auth Middleware
// 通过 r.Filter(<middleware>) 添加中间件
func (a *KeyauthAuther) RestfulAuthHandlerFunc(
	req *restful.Request,
	resp *restful.Response,
	chain *restful.FilterChain,
) {
	// 认证逻辑
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
			response.Failed(resp.ResponseWriter, exception.NewUnauthorized(err.Error()))
			return
		}
		// 处理完成后 需要直接中断返回
		// *restful.Response
		// resp.Header()
		// resp.Write()
		// 3. 处理完成后，需要直接中断返回
		// req.SetAttribute("token", tk)
		req.SetAttribute("token", tk)
	}

	// 鉴权逻辑
	// 获取meta信息, get, 判断是否开启鉴权
	var isPermEnable bool
	if permV, ok := meta[label.Permission]; ok {
		switch v := permV.(type) {
		case bool:
			isPermEnable = v
		case string:
			isPermEnable = v == "true"
		}
	}
	// 认证后，才能鉴权
	if isAuthEnable && isPermEnable {
		tk := req.Attribute("token").(*token.Token)
		permReq := policy.NewValidatePermissionRequest()
		permReq.Service = a.serviceName
		permReq.Username = tk.Data.UserName

		if meta != nil {
			if v, ok := meta[label.Resource]; ok {
				permReq.Resource, _ = v.(string)
			}
			if v, ok := meta[label.Action]; ok {
				permReq.Action, _ = v.(string)
			}
		}
		_, err := a.perm.ValidatePermission(req.Request.Context(), permReq)
		if err != nil {
			response.Failed(resp.ResponseWriter, exception.NewPermissionDeny(err.Error()))
			return
		}
	}

	// 获取meta信息, get, 判断是否开启鉴权
	var isAuditEnable bool
	if authV, ok := meta[label.Audit]; ok {
		switch v := authV.(type) {
		case bool:
			isAuditEnable = v
		case string:
			isAuditEnable = v == "true"
		}
	}
	start := time.Now()

	// chain 用于将请求flow下去
	chain.ProcessFilter(req, resp)

	cost := time.Now().Sub(start).Milliseconds()

	// 如果有记录Respone需求的话，需要在Process才有进行的
	// 认证后，才能审计
	if isAuthEnable && isAuditEnable {
		tk := req.Attribute("token").(*token.Token)
		auditReq := audit.NewOperateLog(tk.Data.UserName, "", "")
		auditReq.Service = a.serviceName
		auditReq.Url = req.Request.URL.String()
		auditReq.Cost = cost
		auditReq.StatusCode = int64(resp.StatusCode())
		auditReq.UserAgent = req.Request.UserAgent()
		// 采用封装好的库实现用户ip的获取，先判断是否使用代理，没有代理则直接RemoteAddr
		auditReq.RemoteIp = request.GetRemoteIP(req.Request)

		if meta != nil {
			if v, ok := meta[label.Resource]; ok {
				auditReq.Resource, _ = v.(string)
			}
			if v, ok := meta[label.Action]; ok {
				auditReq.Action, _ = v.(string)
			}
		}
		a.log.Debug("audit request %s", auditReq)
		_, err := a.audit.AuditOperate(req.Request.Context(), auditReq)
		if err != nil {
			a.log.Warnf("audit operate failed, %s", err)
			return
		}
	}
}
