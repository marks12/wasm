#!/usr/bin/env bash

GOOS=js GOARCH=wasm go build -o ./public/test.wasm test.go

