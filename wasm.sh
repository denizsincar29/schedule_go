#!/bin/bash

# This script builds the go code and copies the goroot wasm_exec.js file to the output directory

# Define output directory
OUTDIR="dist"

# Create output directory
echo "Creating output directory"
mkdir -p "$OUTDIR"

# Build Go code for WASM
echo "Building go code"
GOOS="js"
GOARCH="wasm"
go build -o "$OUTDIR/main.wasm" ./main_wasm
unset GOOS
unset GOARCH

# Get GOROOT from go env
GOROOT=$(go env GOROOT)
echo "GOROOT: $GOROOT"

# Copy wasm_exec.js
echo "Copying wasm_exec.js and index.html"
cp "$GOROOT/lib/wasm/wasm_exec.js" "$OUTDIR/wasm_exec.js"
cp index.html "$OUTDIR/index.html"
echo "Build process complete.  Files located in '$OUTDIR'"