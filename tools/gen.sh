#!/bin/bash

set -e

# Generate new types
qt \
--in=$(dirname $PWD)/google-cloudevents/jsonschema \
--out=$PWD \
--l=golang

# Move generated library into correct directory
cp -a google/events/. .
rm -r google/events