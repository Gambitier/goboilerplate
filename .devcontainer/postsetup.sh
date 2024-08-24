#!/bin/bash

# Export necessary environment variables
export PKG_CONFIG_PATH=/usr/lib/pkgconfig:/usr/local/lib/pkgconfig
export GOFLAGS=-buildvcs=false

# Print Go version
go version

# Setup dev environment
make dev-init

# Print the HOST_IP for verification
make host-ip
