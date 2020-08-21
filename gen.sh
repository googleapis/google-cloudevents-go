#!/bin/bash
set -e

EVENTS_SPEC_REPO='https://github.com/michaelawyu/google-cloudevents'

echo "Cleaning things up..."
for f in *; do
    if [ -d "$f" ]; then
        rm -rf $f 
    fi
done

echo "Cloning the event specification repository from GitHub..."
git clone $EVENTS_SPEC_REPO workplace/

echo "Checking if quicktype is installed..."
command -v quicktype
if [ $? -ne 0 ]; then
    echo "quicktype is not installed on the system."
    echo "To install the quicktype package, see https://quicktype.io/."
    exit 1
fi

echo "Generating the Google Events Library for Go..."
go run gen.go

echo "Cleaning things up..."
rm -rf workplace/ || true

echo "All done."