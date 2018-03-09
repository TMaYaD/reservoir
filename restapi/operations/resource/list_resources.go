// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListResourcesHandlerFunc turns a function with the right signature into a list resources handler
type ListResourcesHandlerFunc func(ListResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListResourcesHandlerFunc) Handle(params ListResourcesParams) middleware.Responder {
	return fn(params)
}

// ListResourcesHandler interface for that can handle valid list resources params
type ListResourcesHandler interface {
	Handle(ListResourcesParams) middleware.Responder
}

// NewListResources creates a new http.Handler for the list resources operation
func NewListResources(ctx *middleware.Context, handler ListResourcesHandler) *ListResources {
	return &ListResources{Context: ctx, Handler: handler}
}

/*ListResources swagger:route GET /resource resource listResources

List all defined resources and their claims

*/
type ListResources struct {
	Context *middleware.Context
	Handler ListResourcesHandler
}

func (o *ListResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
