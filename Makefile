SHELL := /bin/bash -eu -o pipefail

GO_TEST_FLAGS  := -v

PACKAGE_DIRS := $(shell go list ./... 2> /dev/null | grep -v /vendor/)
SRC_FILES    := $(shell find . -name '*.go' -not -path './vendor/*')


.PHONY: test
test:
	@go test $(GO_TEST_FLAGS) $(PACKAGE_DIRS)

.PHONY: ci-test
ci-test:
	@echo > coverage.txt
	@for d in $(PACKAGE_DIRS); do \
		go test -coverprofile=profile.out -covermode=atomic -race -v $$d; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done
	rm coverage.txt