// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetResourceHandlerFunc turns a function with the right signature into a get resource handler
type GetResourceHandlerFunc func(GetResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResourceHandlerFunc) Handle(params GetResourceParams) middleware.Responder {
	return fn(params)
}

// GetResourceHandler interface for that can handle valid get resource params
type GetResourceHandler interface {
	Handle(GetResourceParams) middleware.Responder
}

// NewGetResource creates a new http.Handler for the get resource operation
func NewGetResource(ctx *middleware.Context, handler GetResourceHandler) *GetResource {
	return &GetResource{Context: ctx, Handler: handler}
}

/*GetResource swagger:route GET /resource/{template} resource getResource

Find resource by template

Returns a single resource

*/
type GetResource struct {
	Context *middleware.Context
	Handler GetResourceHandler
}

func (o *GetResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetResourceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
