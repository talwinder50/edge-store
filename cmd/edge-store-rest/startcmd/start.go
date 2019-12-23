/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package startcmd

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/trustbloc/edge-store/pkg/restapi/edv"
	"github.com/trustbloc/edge-store/pkg/storage/memstore"
	cmdutils "github.com/trustbloc/edge-store/pkg/utils/cmd"
)

const (
	hostURLFlagName      = "host-url"
	hostURLFlagShorthand = "u"
	hostURLFlagUsage     = "URL to run the edge-store instance on. Format: HostName:Port."
	hostURLEnvKey        = "EDGE-STORE_HOST_URL"
)

var errMissingHostURL = errors.New("host URL not provided")

type edgeStoreParameters struct {
	srv     server
	hostURL string
}

type server interface {
	ListenAndServe(host string, router http.Handler) error
}

// HTTPServer represents an actual HTTP server implementation.
type HTTPServer struct{}

// ListenAndServe starts the server using the standard Go HTTP server implementation.
func (s *HTTPServer) ListenAndServe(host string, router http.Handler) error {
	return http.ListenAndServe(host, router)
}

// GetStartCmd returns the Cobra start command.
func GetStartCmd(srv server) *cobra.Command {
	startCmd := createStartCmd(srv)

	createFlags(startCmd)

	return startCmd
}

func createStartCmd(srv server) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start edge-store",
		Long:  "Start edge-store",
		RunE: func(cmd *cobra.Command, args []string) error {
			hostURL, err := cmdutils.GetUserSetVar(cmd, hostURLFlagName, hostURLEnvKey)
			if err != nil {
				return err
			}
			parameters := &edgeStoreParameters{
				srv:     srv,
				hostURL: hostURL,
			}
			return startEdgeStore(parameters)
		},
	}
}

func createFlags(startCmd *cobra.Command) {
	startCmd.Flags().StringP(hostURLFlagName, hostURLFlagShorthand, "", hostURLFlagUsage)
}

func startEdgeStore(parameters *edgeStoreParameters) error {
	if parameters.hostURL == "" {
		return errMissingHostURL
	}

	edvService, err := edv.New(memstore.NewProvider())
	if err != nil {
		return err
	}

	handlers := edvService.GetOperations()
	router := mux.NewRouter()

	for _, handler := range handlers {
		router.HandleFunc(handler.Path(), handler.Handle()).Methods(handler.Method())
	}

	err = parameters.srv.ListenAndServe(parameters.hostURL, router)

	return err
}
