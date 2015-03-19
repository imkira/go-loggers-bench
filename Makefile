.PHONY: all deps test

all: deps test test-text test-json

deps:
	go get -u github.com/Sirupsen/logrus
	go get -u gopkg.in/inconshreveable/log15.v2
	go get -u github.com/op/go-logging
	go get -u github.com/cihub/seelog

test: test-text test-json

test-text:
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench Logrus.*Text
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench Log15.*Text
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench Gologging.*Text
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench Seelog.*Text

test-json:
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench Logrus.*JSON
	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench Log15.*JSON
