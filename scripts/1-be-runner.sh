#!/bin/bash

set -e

echo "Cleaning old build..."
rm -f weather-server

echo "Building Go server..."
if go build -o weather-server ./cmd/server; then
  echo "Build succeeded."
else
  echo "Build failed! Exiting."
  exit 1
fi

echo "Server running at: http://localhost:9090/"
./weather-server