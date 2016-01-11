#!/bin/sh

set -x

# test
goapp get golang.org/x/tools/cmd/cover
goapp get github.com/stretchr/testify

# AppEngine

goapp get google.golang.org/appengine
goapp get google.golang.org/appengine/datastore
goapp get google.golang.org/cloud/storage
goapp get golang.org/x/oauth2/google

# Other library

