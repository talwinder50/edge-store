/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/trustbloc/edge-store/pkg/internal/common/support"
)

const (
	token = "/token"
)

// Handler http handler for each controller API endpoint
type Handler interface {
	Path() string
	Method() string
	Handle() http.HandlerFunc
}

// New returns authorization instance
func New() *Operation {
	svc := &Operation{}
	svc.registerHandler()

	return svc
}

// Operation defines handlers for authorization service
type Operation struct {
	handlers []Handler
}

// GetToken swagger:route GET /token edge-auth
//
// get authorization token.
//
// Responses:
//    default: genericError
//        200: getTokenResponse
func (c *Operation) GetToken(rw http.ResponseWriter, req *http.Request) {
	// Implement
	c.writeResponse(rw, "token")
}

// registerHandler register handlers to be exposed from this service as REST API endpoints
func (c *Operation) registerHandler() {
	// Add more protocol endpoints here to expose them as controller API endpoints
	c.handlers = []Handler{
		support.NewHTTPHandler(token, http.MethodGet, c.GetToken),
	}
}

// writeResponse writes interface value to response
func (c *Operation) writeResponse(rw io.Writer, v interface{}) {
	err := json.NewEncoder(rw).Encode(v)
	// as of now, just log errors for writing response
	if err != nil {
		log.Errorf("Unable to send error response, %s", err)
	}
}

// GetRESTHandlers get all controller API handler available for this service
func (c *Operation) GetRESTHandlers() []Handler {
	return c.handlers
}
