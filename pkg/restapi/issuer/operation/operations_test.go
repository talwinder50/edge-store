/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"

	"github.com/stretchr/testify/require"
)

const (
	testCreateCredentialRequest = `{
"context":"https://www.w3.org/2018/credentials/examples/v1",
"type": [
    "VerifiableCredential",
    "UniversityDegreeCredential"
  ],
  "credentialSubject": {
    "id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
    "degree": {
      "type": "BachelorDegree",
      "university": "MIT"
    },
    "name": "Jayden Doe",
    "spouse": "did:example:c276e12ec21ebfeb1f712ebc6f1"
  },

  "issuer": {
    "id": "did:example:76e12ec712ebc6f1c221ebfeb1f",
    "name": "Example University"
  }
}`
)

func TestCreateCredentialHandler_InvalidJSON(t *testing.T) {
	op := New()

	createCredentialHandler := getHandler(t, op, createCredentialEndpoint)

	req, err := http.NewRequest(http.MethodPost, "/credentials", bytes.NewBuffer([]byte("")))
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	createCredentialHandler.Handle().ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Contains(t, rr.Body.String(), "Credential creation failed: EOF")
}
func TestCreateCredential(t *testing.T) {
	op := New()
	createCredentialSuccess(t, op)
}

func createCredentialSuccess(t *testing.T, op *Operation) {
	req, err := http.NewRequest(http.MethodPost, createCredentialEndpoint,
		bytes.NewBuffer([]byte(testCreateCredentialRequest)))
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	createEndpointHandler := getHandler(t, op, createCredentialEndpoint)
	createEndpointHandler.Handle().ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)

	vc := verifiable.Credential{}
	err = json.Unmarshal(rr.Body.Bytes(), &vc)
	require.NoError(t, err)

	require.Equal(t, "did:example:76e12ec712ebc6f1c221ebfeb1f", vc.Issuer.ID)
	require.Equal(t, "Example University", vc.Issuer.Name)
	require.Equal(t, ID, vc.ID)
}

func getHandler(t *testing.T, op *Operation, lookup string) Handler {
	return getHandlerWithError(t, op, lookup)
}

func getHandlerWithError(t *testing.T, op *Operation, lookup string) Handler {
	return handlerLookup(t, op, lookup)
}

func handlerLookup(t *testing.T, op *Operation, lookup string) Handler {
	handlers := op.GetRESTHandlers()
	require.NotEmpty(t, handlers)

	for _, h := range handlers {
		if h.Path() == lookup {
			return h
		}
	}

	require.Fail(t, "unable to find handler")

	return nil
}
