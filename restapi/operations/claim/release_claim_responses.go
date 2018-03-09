// Code generated by go-swagger; DO NOT EDIT.

package claim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ReleaseClaimBadRequestCode is the HTTP code returned for type ReleaseClaimBadRequest
const ReleaseClaimBadRequestCode int = 400

/*ReleaseClaimBadRequest Invalid count supplied

swagger:response releaseClaimBadRequest
*/
type ReleaseClaimBadRequest struct {
}

// NewReleaseClaimBadRequest creates ReleaseClaimBadRequest with default headers values
func NewReleaseClaimBadRequest() *ReleaseClaimBadRequest {
	return &ReleaseClaimBadRequest{}
}

// WriteResponse to the client
func (o *ReleaseClaimBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ReleaseClaimNotFoundCode is the HTTP code returned for type ReleaseClaimNotFound
const ReleaseClaimNotFoundCode int = 404

/*ReleaseClaimNotFound resource not found

swagger:response releaseClaimNotFound
*/
type ReleaseClaimNotFound struct {
}

// NewReleaseClaimNotFound creates ReleaseClaimNotFound with default headers values
func NewReleaseClaimNotFound() *ReleaseClaimNotFound {
	return &ReleaseClaimNotFound{}
}

// WriteResponse to the client
func (o *ReleaseClaimNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}