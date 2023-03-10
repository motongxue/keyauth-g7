package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/motongxue/keyauth-g7/apps/role"
)

func (h *handler) CreateRole(r *restful.Request, w *restful.Response) {
	req := role.NewCreateRoleRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	set, err := h.service.CreateRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
