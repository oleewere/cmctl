GIT_REV_SHORT = $(shell git rev-parse --short HEAD)
GITHUB_TOKEN := $(shell git config --global --get github.token || echo $$GITHUB_TOKEN)

ifeq ("$(shell git rev-list --tags --max-count=1)", "")
  LAST_RELEASE="v0.0.0"
else
  LAST_RELEASE=$(shell git describe --tags $(shell git rev-list --tags --max-count=1))
endif

SNAPSHOT_VERSION=$(shell echo $(LAST_RELEASE) | awk '{split($$0,a,"."); print "v"a[1]+1"."0"."0}')

ifeq ("$(shell git name-rev --tags --name-only $(shell git rev-parse HEAD))", "undefined")
	VERSION_FOR_BUILD="$(SNAPSHOT_VERSION)-SNAPSHOT"
else
	VERSION_FOR_BUILD=$(shell git name-rev --tags --name-only $(shell git rev-parse HEAD) | sed 's/\^.*$///')
endif

install:
	go install -ldflags "-X main.GitRevString=$(GIT_REV_SHORT) -X main.Version=$(VERSION_FOR_BUILD)" .

build:
	go build -ldflags "-X main.GitRevString=$(GIT_REV_SHORT) -X main.Version=$(VERSION_FOR_BUILD)" -o cmctl .

test:
	go test

all: build test

clean:
	rm -rf dist

last-release:
	@echo "Last release version: $(LAST_RELEASE)"

version:
	@echo "Release/Snapshot version: $(VERSION_FOR_BUILD)"