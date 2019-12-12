/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package createdemodata

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

const testURL = "http://localhost:1337"

func TestStartCmdContents(t *testing.T) {
	startCmd := GetStartCmd()

	require.Equal(t, "create-demo-data", startCmd.Use)
	require.Equal(t, "Create demo data", startCmd.Short)
	require.Equal(t, "Start populating data in strapi with default studentcards and transcripts", startCmd.Long)

	checkFlagPropertiesCorrect(t, startCmd, adminURLFlagName, adminURLFlagShorthand, adminURLFlagUsage)
}

func TestStartCmdWithBlankHostArg(t *testing.T) {
	startCmd := GetStartCmd()

	args := []string{"--" + adminURLFlagName, ""}
	startCmd.SetArgs(args)

	err := startCmd.Execute()

	require.Equal(t, errMissingAdminURL.Error(), err.Error())
}

func TestStartCmdWithMissingHostArg(t *testing.T) {
	startCmd := GetStartCmd()
	err := startCmd.Execute()

	require.Equal(t,
		"Neither admin-url (command line flag) nor STRAPI-DEMO_ADMIN_URL (environment variable) have been set.",
		err.Error())
}
func TestStartEdgeStoreWithBlankHost(t *testing.T) {
	parameters := &strapiDemoParameters{adminURL: ""}

	err := startStrapiDemo(parameters)
	require.NotNil(t, err)
	require.Equal(t, errMissingAdminURL, err)
}

func checkFlagPropertiesCorrect(t *testing.T, cmd *cobra.Command, flagName, flagShorthand, flagUsage string) {
	flag := cmd.Flag(flagName)

	require.NotNil(t, flag)
	require.Equal(t, flagName, flag.Name)
	require.Equal(t, flagShorthand, flag.Shorthand)
	require.Equal(t, flagUsage, flag.Usage)
	require.Equal(t, "", flag.Value.String())

	flagAnnotations := flag.Annotations
	require.Nil(t, flagAnnotations)
}
func TestAdminUserAndCreateRecordWithRoundTripper(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		return &http.Response{
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
	 	"jwt": "eyJhbGciOiJIU",
		"user": {
        	"id": 12,
        	"username": "strapi",
        	"email": "user@strapi.io",
        	"isAdmin": true
    	}
	}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	adminUserValues := map[string]string{"username": "strapi"}
	token, err := createAdminUser(client, testURL, adminUserValues)
	require.NotNil(t, token)
	require.Nil(t, err)
	require.Equal(t, "Bearer eyJhbGciOiJIU", token)

	studentRecord := map[string]string{
		"studentid": "1234568",
		"name":      "Tanu",
	}

	createOrFetchRecord(client, token, testURL+studentCardsEndpoint, "POST", studentRecord)

	parameters := &strapiDemoParameters{client: client, adminURL: testURL}

	err = startStrapiDemo(parameters)
	require.Nil(t, err)
}
func TestCreateAdminUserAndCreateRecordError(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		return &http.Response{
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	adminUserValues := map[string]string{"username": "strapi"}
	token, err := createAdminUser(client, testURL, adminUserValues)
	require.Equal(t, "", token)
	require.Contains(t, err.Error(), "invalid character")

	parameters := &strapiDemoParameters{client: client, adminURL: testURL}

	err = startStrapiDemo(parameters)
	require.NotNil(t, err)

	token, err = createAdminUser(client, "}}|}", make(chan int))
	require.Equal(t, "", token)
	require.Contains(t, err.Error(), "json: unsupported type: chan int")

	createOrFetchRecord(client, token, testURL+studentCardsEndpoint, "POST", make(chan int))
	require.Contains(t, err.Error(), "json: unsupported type: chan int")
}

// RoundTripFunc RoundTripper is an interface representing the ability to execute a single HTTP transaction,
// obtaining the Response for a given Request.
// https://golang.org/pkg/net/http/#RoundTripper
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip http.RoundTripper Interface has just one method RoundTrip(*Request) (*Response, error)
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client { //nolint : interfacer
	return &http.Client{
		Transport: fn,
	}
}
