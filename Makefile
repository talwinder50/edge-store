# Copyright SecureKey Technologies Inc.
#
# SPDX-License-Identifier: Apache-2.0

GO_CMD ?= go
EDGE_AUTH_REST_PATH=cmd/edge-auth-rest

.PHONY: all
all: checks unit-test

.PHONY: checks
checks: license lint

.PHONY: lint
lint:
	@scripts/check_lint.sh

.PHONY: license
license:
	@scripts/check_license.sh

unit-test:
	@scripts/check_unit.sh

hydra-start:
	@scripts/hydra_start.sh

hydra-configure:
	@scripts/hydra_configure.sh

hydra-stop:
	@scripts/hydra_stop.sh

hydra-test-app:
	@scripts/hydra_test_app.sh

edge-auth-rest:
	@echo "Building edge-auth-rest"
	@mkdir -p ./build/bin
	@cd ${EDGE_AUTH_REST_PATH} && go build -o ../../build/bin/edge-auth-rest main.go