package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/motongxue/keyauth-g7/apps/token"
)

var (
	h = &handler{}
)

type handler struct {
	service token.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(token.AppName)
	h.service = app.GetGrpcApp(token.AppName).(token.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/issue").To(h.IssueToken).
		Doc("issue token").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(response.NewData(token.Token{})))

	// ws.Route(ws.GET("/").To(h.QueryUser).
	// 	Doc("get all users").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Metadata("action", "list").
	// 	Reads(user.CreateUserRequest{}).
	// 	Writes(response.NewData(user.UserSet{})).
	// 	Returns(200, "OK", user.UserSet{}))

	// ws.Route(ws.GET("/{id}").To(h.DescribeUser).
	// 	Doc("get a user").
	// 	Param(ws.PathParameter("id", "identifier of the user").DataType("integer").DefaultValue("1")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Writes(response.NewData(user.User{})).
	// 	Returns(200, "OK", response.NewData(user.User{})).
	// 	Returns(404, "Not Found", nil))

	// ws.Route(ws.PUT("/{id}").To(h.UpdateUser).
	// 	Doc("update a user").
	// 	Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(user.CreateUserRequest{}))

	// ws.Route(ws.PATCH("/{id}").To(h.PatchUser).
	// 	Doc("patch a user").
	// 	Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(user.CreateUserRequest{}))

	// ws.Route(ws.DELETE("/{id}").To(h.DeleteUser).
	// 	Doc("delete a user").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Param(ws.PathParameter("id", "identifier of the user").DataType("string")))
	ws.Route(ws.GET("/validate").To(h.ValidateToken).
		Doc("validate token").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "get"))

	ws.Route(ws.POST("/revolk").To(h.RevolkToken).
		Doc("revolk token").
		Param(ws.PathParameter("id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "delete"))
}

func init() {
	app.RegistryRESTfulApp(h)
}
