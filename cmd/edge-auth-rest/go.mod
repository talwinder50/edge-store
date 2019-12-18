// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/edge-store/cmd/edge-auth-rest

replace github.com/trustbloc/edge-store => ../..

require (
	github.com/gorilla/mux v1.7.3
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	github.com/trustbloc/edge-store v0.0.0
)

go 1.13
