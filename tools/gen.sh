#!/bin/bash
set -e

echo "Generating the Google Events Library..."

ROOT=$(dirname $PWD)

find cloud -type f ! -name '*_test.go' -delete
find firebase -type f ! -name '*_test.go' -delete

# Generate new types
qt \
--in=$ROOT/google-cloudevents/jsonschema \
--out=$PWD \
--enum-as-string=true \
--l=golang

# Move generated library into correct directory
cp -a google/events/. .
rm -r google

# Run postgen
echo '- Running postgen step'
(cd tools && npm start)

# Re-format the golang files
echo "- Re-format golang files"
gofmt -s -w .
