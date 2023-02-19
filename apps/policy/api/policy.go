package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/motongxue/keyauth-g7/apps/policy"

	"github.com/infraboard/mcube/http/response"
)

func (h *handler) CreatePolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewCreatePolicyRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.CreatePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}
