package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/motongxue/keyauth-g7/apps/token"
	"strings"

	"github.com/infraboard/mcube/http/response"
)

func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {
	req := token.NewIssueTokenRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (h *handler) ValidateToken(r *restful.Request, w *restful.Response) {
	// 我们Token哪里获取?
	// 1. URL Query String ?
	// 2. Custom Header ?
	// 3. Authorization Header

	accessToken := ""
	auth := r.HeaderParameter("Authorization")
	al := strings.Split(auth, " ")
	if len(al) > 1 {
		accessToken = al[1]
	} else {
		// 兼容 Authorization <token>
		accessToken = auth
	}

	req := token.NewValidateTokenRequest(accessToken)
	ins, err := h.service.ValidateToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (h *handler) RevolkToken(r *restful.Request, w *restful.Response) {
	req := token.NewRevolkTokenRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	set, err := h.service.RevolkToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
