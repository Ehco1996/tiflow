LDFLAGS += -X "github.com/pingcap/tidb-enterprise-tools/pkg/utils.ReleaseVersion=$(shell git describe --tags --dirty="-dev")"
LDFLAGS += -X "github.com/pingcap/tidb-enterprise-tools/pkg/utils.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "github.com/pingcap/tidb-enterprise-tools/pkg/utils.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "github.com/pingcap/tidb-enterprise-tools/pkg/utils.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "github.com/pingcap/tidb-enterprise-tools/pkg/utils.GoVersion=$(shell go version)"


CURDIR   := $(shell pwd)
GO       := GO15VENDOREXPERIMENT="1" go
GOBUILD  := CGO_ENABLED=0 $(GO) build
GOTEST   := GOPATH=$(CURDIR)/_vendor:$(GOPATH) CGO_ENABLED=1 $(GO) test
PACKAGES := $$(go list ./... | grep -vE 'vendor')
FILES    := $$(find . -name "*.go" | grep -vE "vendor")
TOPDIRS  := $$(ls -d */ | grep -vE "vendor")
TIDBDIR  := vendor/github.com/pingcap/tidb
SHELL    := /usr/bin/env bash

RACE_FLAG = 
ifeq ("$(WITH_RACE)", "1")
	RACE_FLAG = -race
	GOBUILD   = GOPATH=$(GOPATH) CGO_ENABLED=1 $(GO) build
endif


ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

.PHONY: build syncer loader test check deps dm-worker dm-master dmctl 

build: syncer loader check test dm-worker dm-master dmctl

dm-worker:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o bin/dm-worker ./cmd/dm-worker

dm-master:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o bin/dm-master ./cmd/dm-master

dmctl:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o bin/dmctl ./cmd/dm-ctl

syncer:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o bin/syncer ./cmd/syncer

loader:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o bin/loader ./cmd/loader

test:
	bash -x ./wait_for_mysql.sh
	@export log_level=error; \
	$(GOTEST) -cover $(PACKAGES)

check: fmt lint vet

fmt:
	@echo "gofmt (simplify)"
	@ gofmt -s -l -w $(FILES) 2>&1 | awk '{print} END{if(NR>0) {exit 1}}'

errcheck:
	go get github.com/kisielk/errcheck
	@echo "errcheck"
	@ GOPATH=$(GOPATH) errcheck -blank $(PACKAGES) | grep -v "_test\.go" | awk '{print} END{if(NR>0) {exit 1}}'

lint:
	go get github.com/golang/lint/golint
	@echo "golint"
	@ golint -set_exit_status $(PACKAGES)

vet:
	@echo "vet"
	@ go tool vet -all -shadow $(TOPDIRS) 2>&1 | tee /dev/stderr | awk '/shadows declaration/{next}{count+=1} END{if(count>0) {exit 1}}'
	

update: update_vendor parser clean_vendor
update_vendor:
	@which glide >/dev/null || curl https://glide.sh/get | sh
	@which glide-vc || go get -v -u github.com/sgotti/glide-vc
ifdef PKG
	@glide get -s -v --skip-test ${PKG}
else
	@glide update -s -v -u --skip-test
endif

goyacc:
	$(GOBUILD) -o $(TIDBDIR)/bin/goyacc $(TIDBDIR)/parser/goyacc/main.go

parser: goyacc
	$(TIDBDIR)/bin/goyacc -o /dev/null $(TIDBDIR)/parser/parser.y
	$(TIDBDIR)/bin/goyacc -o $(TIDBDIR)/parser/parser.go $(TIDBDIR)/parser/parser.y 2>&1 | egrep "(shift|reduce)/reduce" | awk '{print} END {if (NR > 0) {print "Find conflict in parser.y. Please check y.output for more information."; exit 1;}}'
	rm -f y.output

	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		sed -i -e 's|//line.*||' -e 's/yyEofCode/yyEOFCode/' $(TIDBDIR)/parser/parser.go; \
	elif [ $(ARCH) = $(MAC) ]; \
	then \
		/usr/bin/sed -i "" 's|//line.*||' $(TIDBDIR)/parser/parser.go; \
		/usr/bin/sed -i "" 's/yyEofCode/yyEOFCode/' $(TIDBDIR)/parser/parser.go; \
	fi

	@awk 'BEGIN{print "// Code generated by goyacc"} {print $0}' $(TIDBDIR)/parser/parser.go > tmp_parser.go && mv tmp_parser.go $(TIDBDIR)/parser/parser.go;

clean_vendor:
	@echo "removing test files"
	@glide vc --use-lock-file --only-code --no-tests
