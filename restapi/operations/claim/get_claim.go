// Code generated by go-swagger; DO NOT EDIT.

package claim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetClaimHandlerFunc turns a function with the right signature into a get claim handler
type GetClaimHandlerFunc func(GetClaimParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClaimHandlerFunc) Handle(params GetClaimParams) middleware.Responder {
	return fn(params)
}

// GetClaimHandler interface for that can handle valid get claim params
type GetClaimHandler interface {
	Handle(GetClaimParams) middleware.Responder
}

// NewGetClaim creates a new http.Handler for the get claim operation
func NewGetClaim(ctx *middleware.Context, handler GetClaimHandler) *GetClaim {
	return &GetClaim{Context: ctx, Handler: handler}
}

/*GetClaim swagger:route GET /claim/{name} claim getClaim

Find claim by name

name should be an fqdn and match one of the predefined resource templates.

*/
type GetClaim struct {
	Context *middleware.Context
	Handler GetClaimHandler
}

func (o *GetClaim) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetClaimParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
