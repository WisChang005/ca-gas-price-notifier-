#! /bin/bash

env GOOS=linux GOARCH=arm GOARM=5 go build -o output/gas-notifier src/cmd/main.go