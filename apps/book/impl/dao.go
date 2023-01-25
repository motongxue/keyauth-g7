package impl

import (
	"context"
	"fmt"

	"github.com/motongxue/keyauth-g7/apps/book"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Save Object
func (s *service) save(ctx context.Context, ins *book.Book) error {
	// s.col.InsertMany()
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted book(%s) document error, %s",
			ins.Data.Name, err)
	}
	return nil
}

// GET, Describe
// filter 过滤器(Collection),类似于MYSQL Where条件
// 调用Decode方法来进行 反序列化  bytes ---> Object (通过BSON Tag)
func (s *service) get(ctx context.Context, id string) (*book.Book, error) {
	filter := bson.M{"_id": id}

	ins := book.NewDefaultBook()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("book %s not found", id)
		}

		return nil, exception.NewInternalServerError("find book %s error, %s", id, err)
	}

	return ins, nil
}

func newQueryBookRequest(r *book.QueryBookRequest) *queryBookRequest {
	return &queryBookRequest{
		r,
	}
}

// 把QueryReq --> MongoDB Options
type queryBookRequest struct {
	*book.QueryBookRequest
}

// Find参数
func (r *queryBookRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		// 排序: Order By create_at Desc
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
		// 分页: limit 0,10  skip:0, limit:10
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

// 过滤条件
// 由于Mongodb支持嵌套, JSON, 如何过滤嵌套里面的条件, 使用.访问嵌套对象属性
func (r *queryBookRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.Keywords != "" {
		filter["$or"] = bson.A{
			bson.M{"data.name": bson.M{"$regex": r.Keywords, "$options": "im"}},
			bson.M{"data.author": bson.M{"$regex": r.Keywords, "$options": "im"}},
		}
	}
	return filter
}

// LIST, Query, 会很多条件(分页, 关键字, 条件过滤, 排序)
// 需要单独为其 做过滤参数构建
func (s *service) query(ctx context.Context, req *queryBookRequest) (*book.BookSet, error) {
	resp, err := s.col.Find(ctx, req.FindFilter(), req.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find book error, error is %s", err)
	}

	set := book.NewBookSet()
	// 循环
	for resp.Next(ctx) {
		ins := book.NewDefaultBook()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode book error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get book count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

// UpdateByID, 通过主键来更新对象
func (s *service) update(ctx context.Context, ins *book.Book) error {
	// SQL update obj(SET f=v,f=v) where id=?
	// s.col.UpdateOne(ctx, filter(), ins)
	if _, err := s.col.UpdateByID(ctx, ins.Id, ins); err != nil {
		return exception.NewInternalServerError("inserted book(%s) document error, %s",
			ins.Data.Name, err)
	}

	return nil
}

func (s *service) deleteBook(ctx context.Context, ins *book.Book) error {
	if ins == nil || ins.Id == "" {
		return fmt.Errorf("book is nil")
	}

	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.Id})
	if err != nil {
		return exception.NewInternalServerError("delete book(%s) error, %s", ins.Id, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("book %s not found", ins.Id)
	}

	return nil
}
