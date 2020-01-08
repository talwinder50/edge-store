/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import "github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"

// VCData input data for VC services
type VCData struct {
	Subject verifiable.Subject `json:"credentialSubject"`
	Issuer  verifiable.Issuer  `json:"issuer"`
	Type    []string           `json:"type,omitempty"`
	Context string             `json:"context,omitempty"`
}
