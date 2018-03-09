// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddResourceHandlerFunc turns a function with the right signature into a add resource handler
type AddResourceHandlerFunc func(AddResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddResourceHandlerFunc) Handle(params AddResourceParams) middleware.Responder {
	return fn(params)
}

// AddResourceHandler interface for that can handle valid add resource params
type AddResourceHandler interface {
	Handle(AddResourceParams) middleware.Responder
}

// NewAddResource creates a new http.Handler for the add resource operation
func NewAddResource(ctx *middleware.Context, handler AddResourceHandler) *AddResource {
	return &AddResource{Context: ctx, Handler: handler}
}

/*AddResource swagger:route POST /resource resource addResource

Add a new resource

This increments the total count if the resource with the given name already exists.

*/
type AddResource struct {
	Context *middleware.Context
	Handler AddResourceHandler
}

func (o *AddResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddResourceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
