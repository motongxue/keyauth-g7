package policy

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"time"
)

const (
	AppName = "policy"
)

var (
	validate = validator.New()
)

type Service interface {
	RPCServer
	CreatePolicy(ctx context.Context, request *CreatePolicyRequest) (*Policy, error)
}

func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}

func NewPolicySet() *PolicySet {
	return &PolicySet{
		Items: []*Policy{},
	}
}

func (s *PolicySet) Add(item *Policy) {
	s.Items = append(s.Items, item)
}

func NewDefaultPolicy() *Policy {
	return &Policy{
		Spec: NewCreatePolicyRequest(),
	}
}

func NewPolicy(req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &Policy{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}, nil
}

func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func NewQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

// role 名称的列表
func (s *PolicySet) Roles() (roles []string) {
	for i := range s.Items {
		roles = append(roles, s.Items[i].Spec.Role)
	}
	return
}

// 根据Role名称查询集合里面的策略
func (s *PolicySet) GetPolicyByRole(role string) *Policy {
	for i := range s.Items {
		if s.Items[i].Spec.Role == role {
			return s.Items[i]
		}
	}
	return nil
}

func NewValidatePermissionRequest() *ValidatePermissionRequest {
	return &ValidatePermissionRequest{}
}
