
[![Release](https://img.shields.io/github/release/trustbloc/edge-store.svg?style=flat-square)](https://github.com/trustbloc/edge-store/releases/latest)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/trustbloc/edge-store/master/LICENSE)
[![Godocs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/trustbloc/edge-store)

[![Build Status](https://dev.azure.com/trustbloc/edge/_apis/build/status/trustbloc.edge-store?branchName=master)](https://dev.azure.com/trustbloc/edge/_build/latest?definitionId=27&branchName=master)
[![codecov](https://codecov.io/gh/trustbloc/edge-store/branch/master/graph/badge.svg)](https://codecov.io/gh/trustbloc/edge-store)
[![Go Report Card](https://goreportcard.com/badge/github.com/trustbloc/edge-store)](https://goreportcard.com/report/github.com/trustbloc/edge-store)



Create Demo Data using [Strapi](https://strapi.io/)
### make strapi-start 
Strapi start bring up the strapi on default 1337 port using mysql. 
Once the strapi is up and running you can verify by browsing http://localhost:1337/admin/

### make strapi-setup
Strapi setup 
1. Create the schemas APIs studentcards and transcripts 
2. Create an admin user (username/password: strapi/strapi)
3. Populate the schemas APIs with the data

### Verify demo data 
1. Login info with the admin user credentials http://localhost:1337/admin/
2. Browse the Content Type section on the left side to verify the data. 
