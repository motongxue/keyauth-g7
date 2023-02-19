package role

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"time"
)

const (
	AppName = "role"
)

var (
	validate = validator.New()
)

type Service interface {
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	RPCServer
}

func NewRoleSet() *RoleSet {
	return &RoleSet{
		Items: []*Role{},
	}
}

func (s *RoleSet) Add(item *Role) {
	s.Items = append(s.Items, item)
}

func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Permissions: []*Permission{},
		Meta:        map[string]string{},
	}
}

func NewDefaultRole() *Role {
	return &Role{
		Spec: NewCreateRoleRequest(),
	}
}

func NewRole(req *CreateRoleRequest) *Role {
	return &Role{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}
}

func NewQueryRoleRequestWithName(names []string) *QueryRoleRequest {
	return &QueryRoleRequest{
		Page:      request.NewDefaultPageRequest(),
		RoleNames: names,
	}
}

func (s *RoleSet) HasPermission(req *PermissionRequest) (bool, *Role) {
	for i := range s.Items {
		if s.Items[i].HasPermission(req) {
			return true, s.Items[i]
		}
	}
	return false, nil
}
func (r *Role) HasPermission(req *PermissionRequest) bool {
	for i := range r.Spec.Permissions {
		if r.Spec.Permissions[i].HasPermission(req) {
			return true
		}
	}
	return false
}
func (p *Permission) HasPermission(req *PermissionRequest) bool {
	if p.Service != req.Service {
		return false
	}
	// 放行所有功能
	if p.AllowAll {
		return true
	}
	// 查询所有的功能,确认是否允许
	for i := range p.Features {
		if p.Features[i].Action == req.Action &&
			p.Features[i].Resource == req.Resource {
			return true
		}
	}
	return false
}

func NewPermissionRequest(service, resource, action string) *PermissionRequest {
	return &PermissionRequest{
		Service:  service,
		Resource: resource,
		Action:   action,
	}
}
func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}
