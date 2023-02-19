package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/motongxue/keyauth-g7/apps/policy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) save(ctx context.Context, ins *policy.Policy) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.Id, err)
	}
	return nil
}

func newQueryPolicyRequest(r *policy.QueryPolicyRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

// 把QueryReq --> MongoDB Options
type queryRequest struct {
	*policy.QueryPolicyRequest
}

// Find参数
func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)
	opt := &options.FindOptions{
		Sort: bson.D{
			{
				Key: "create_at", Value: -1,
			},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}
	return opt
}

// 过滤条件
// 由于Mongodb支持嵌套, JSON, 如何过滤嵌套里面的条件, 使用.访问嵌套对象属性
func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.Namespace != "" {
		filter["spec.namespace"] = r.Namespace
	}
	if r.Username != "" {
		filter["spec.username"] = r.Username
	}
	if r.Role != "" {
		filter["spec.role"] = r.Role
	}
	return filter
}

// LIST, Query, 会很多条件(分页, 关键字, 条件过滤, 排序)
// 需要单独为其 做过滤参数构建
func (s *service) query(ctx context.Context, req *queryRequest) (*policy.PolicySet, error) {
	resp, err := s.col.Find(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}
	policySet := policy.NewPolicySet()
	for resp.Next(ctx) {
		ins := policy.NewDefaultPolicy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}
		policySet.Add(ins)
	}
	count, err := s.col.CountDocuments(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get policy count error, error is %s", err)
	}
	policySet.Total = count
	return policySet, nil
}
