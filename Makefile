GOTEST_FLAGS=-cpu=1,2,4 -benchmem -benchtime=5s

TEXT_PKGS=Gokit Logrus Log15 Gologging Seelog Zerolog Fortiolog
JSON_PKGS=Gokit Logrus Log15 Zerolog

TEXT_PKG_TARGETS=$(addprefix test-text-,$(TEXT_PKGS))
JSON_PKG_TARGETS=$(addprefix test-json-,$(JSON_PKGS))

.PHONY: all deps test test-text test-json $(TEXT_PKG_TARGETS) $(JSON_PKG_TARGETS)

all: deps test

deps:
	go get -u github.com/Sirupsen/logrus
	go get -u gopkg.in/inconshreveable/log15.v2
	go get -u github.com/op/go-logging
	go get -u github.com/cihub/seelog
	go get -u github.com/go-kit/kit/log
	go get -u github.com/rs/zerolog
	go get -u fortio.org/fortio

test: test-text test-json

test-text: $(TEXT_PKG_TARGETS)

$(TEXT_PKG_TARGETS): test-text-%:
	go test $(GOTEST_FLAGS) -bench "$*.*Text"

test-json: $(JSON_PKG_TARGETS)

$(JSON_PKG_TARGETS): test-json-%:
	go test $(GOTEST_FLAGS) -bench "$*.*JSON"
