package auth

import "github.com/emicklei/go-restful/v3"

// Go-Restful auth Middleware
func (a *KeyauthAuther) RestfulAuthHandlerFunc(
	req *restful.Request,
	resp *restful.Response,
	chain *restful.FilterChain,
) {

}
