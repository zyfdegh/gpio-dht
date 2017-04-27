#!/bin/sh

go get github.com/stianeikeland/go-rpio

export GOARCH=arm
export GOARM=7

go build -o bin/dht

# docker build -t zyfdedh/gpio-dht:dev .
# docker push zyfdedh/gpio-dht:dev
