// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/TMaYaD/reservoir/restapi/operations/claim"
	"github.com/TMaYaD/reservoir/restapi/operations/resource"
)

// NewReservoirAPI creates a new Reservoir instance
func NewReservoirAPI(spec *loads.Document) *ReservoirAPI {
	return &ReservoirAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		UrlformConsumer:     runtime.DiscardConsumer,
		JSONProducer:        runtime.JSONProducer(),
		ResourceAddResourceHandler: resource.AddResourceHandlerFunc(func(params resource.AddResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourceAddResource has not yet been implemented")
		}),
		ClaimClaimResourceHandler: claim.ClaimResourceHandlerFunc(func(params claim.ClaimResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation ClaimClaimResource has not yet been implemented")
		}),
		ResourceDeleteResourceHandler: resource.DeleteResourceHandlerFunc(func(params resource.DeleteResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourceDeleteResource has not yet been implemented")
		}),
		ClaimGetClaimHandler: claim.GetClaimHandlerFunc(func(params claim.GetClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation ClaimGetClaim has not yet been implemented")
		}),
		ResourceGetResourceHandler: resource.GetResourceHandlerFunc(func(params resource.GetResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourceGetResource has not yet been implemented")
		}),
		ResourceListResourcesHandler: resource.ListResourcesHandlerFunc(func(params resource.ListResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourceListResources has not yet been implemented")
		}),
		ClaimReleaseClaimHandler: claim.ReleaseClaimHandlerFunc(func(params claim.ReleaseClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation ClaimReleaseClaim has not yet been implemented")
		}),
		ResourceSetResourceHandler: resource.SetResourceHandlerFunc(func(params resource.SetResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourceSetResource has not yet been implemented")
		}),
		ResourceUpdateResourceWithFormHandler: resource.UpdateResourceWithFormHandlerFunc(func(params resource.UpdateResourceWithFormParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourceUpdateResourceWithForm has not yet been implemented")
		}),
	}
}

/*ReservoirAPI Reservoir is a resource manager with support for wildcards */
type ReservoirAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// UrlformConsumer registers a consumer for a "application/x-www-form-urlencoded" mime type
	UrlformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ResourceAddResourceHandler sets the operation handler for the add resource operation
	ResourceAddResourceHandler resource.AddResourceHandler
	// ClaimClaimResourceHandler sets the operation handler for the claim resource operation
	ClaimClaimResourceHandler claim.ClaimResourceHandler
	// ResourceDeleteResourceHandler sets the operation handler for the delete resource operation
	ResourceDeleteResourceHandler resource.DeleteResourceHandler
	// ClaimGetClaimHandler sets the operation handler for the get claim operation
	ClaimGetClaimHandler claim.GetClaimHandler
	// ResourceGetResourceHandler sets the operation handler for the get resource operation
	ResourceGetResourceHandler resource.GetResourceHandler
	// ResourceListResourcesHandler sets the operation handler for the list resources operation
	ResourceListResourcesHandler resource.ListResourcesHandler
	// ClaimReleaseClaimHandler sets the operation handler for the release claim operation
	ClaimReleaseClaimHandler claim.ReleaseClaimHandler
	// ResourceSetResourceHandler sets the operation handler for the set resource operation
	ResourceSetResourceHandler resource.SetResourceHandler
	// ResourceUpdateResourceWithFormHandler sets the operation handler for the update resource with form operation
	ResourceUpdateResourceWithFormHandler resource.UpdateResourceWithFormHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ReservoirAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ReservoirAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ReservoirAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ReservoirAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ReservoirAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ReservoirAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ReservoirAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ReservoirAPI
func (o *ReservoirAPI) Validate() error {
	var unregistered []string

	if o.UrlformConsumer == nil {
		unregistered = append(unregistered, "UrlformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ResourceAddResourceHandler == nil {
		unregistered = append(unregistered, "resource.AddResourceHandler")
	}

	if o.ClaimClaimResourceHandler == nil {
		unregistered = append(unregistered, "claim.ClaimResourceHandler")
	}

	if o.ResourceDeleteResourceHandler == nil {
		unregistered = append(unregistered, "resource.DeleteResourceHandler")
	}

	if o.ClaimGetClaimHandler == nil {
		unregistered = append(unregistered, "claim.GetClaimHandler")
	}

	if o.ResourceGetResourceHandler == nil {
		unregistered = append(unregistered, "resource.GetResourceHandler")
	}

	if o.ResourceListResourcesHandler == nil {
		unregistered = append(unregistered, "resource.ListResourcesHandler")
	}

	if o.ClaimReleaseClaimHandler == nil {
		unregistered = append(unregistered, "claim.ReleaseClaimHandler")
	}

	if o.ResourceSetResourceHandler == nil {
		unregistered = append(unregistered, "resource.SetResourceHandler")
	}

	if o.ResourceUpdateResourceWithFormHandler == nil {
		unregistered = append(unregistered, "resource.UpdateResourceWithFormHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ReservoirAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ReservoirAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *ReservoirAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ReservoirAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/x-www-form-urlencoded":
			result["application/x-www-form-urlencoded"] = o.UrlformConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ReservoirAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ReservoirAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the reservoir API
func (o *ReservoirAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ReservoirAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resource"] = resource.NewAddResource(o.context, o.ResourceAddResourceHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/claim"] = claim.NewClaimResource(o.context, o.ClaimClaimResourceHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/resource/{template}"] = resource.NewDeleteResource(o.context, o.ResourceDeleteResourceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/claim/{name}"] = claim.NewGetClaim(o.context, o.ClaimGetClaimHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource/{template}"] = resource.NewGetResource(o.context, o.ResourceGetResourceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource"] = resource.NewListResources(o.context, o.ResourceListResourcesHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/claim"] = claim.NewReleaseClaim(o.context, o.ClaimReleaseClaimHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/resource"] = resource.NewSetResource(o.context, o.ResourceSetResourceHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resource/{template}"] = resource.NewUpdateResourceWithForm(o.context, o.ResourceUpdateResourceWithFormHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ReservoirAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *ReservoirAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
