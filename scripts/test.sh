#! /bin/bash

set -e

echo '-- Running tests and generating profile...'

go test -coverprofile coverage.out

echo '-- Generating HTML coverate report'

go tool cover -html coverage.out