#!/bin/bash
set -e

echo "Generating the Google Events Library..."

ROOT=$(dirname $PWD)

# Generate new types
qt \
--in=$ROOT/google-cloudevents/jsonschema \
--out=$PWD \
--l=golang

# Move generated library into correct directory
cp -a google/events/. .
rm -r google

echo $ROOT

# Run postgen
cd tools
npm start
