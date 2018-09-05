## ----- Variables -----
PKG_NAME = $(shell basename "$$(pwd)")

## Source config
SRC_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
SRC_PKGS = $(shell go list ./... | grep -v /vendor/)

## Testing config
TEST_TIMEOUT = 20s
COVER_OUT = coverage.out


## ------ Commands (targets) -----
## Prevent targeting filenames...
.PHONY: default init setup verify run build install get update fix \
		_review_base review review-race review-bench check fmt test test-v \
		test-race bench

## Default target when no arguments are given to make (build and run program).
default: build run

## Setup sets up a Go module by initializing the module and then verifying
## its dependencies.
setup: init verify

## Initializes a Go module in the current directory.
init:
	@printf "Initializing Go module:\n"
	@go mod init

## Builds and runs the program (package must be main).
run:
	@if [ -f ".env.sh" ]; then \
	   printf 'Exporting environment variables by sourcing ".env.sh"... '; \
	   . .env.sh; \
	   printf "done.\n"; \
	 fi
	@if [ -f "$(PKG_NAME)" ]; then \
	   printf 'Running "$(PKG_NAME)"...\n'; \
	   ./$(PKG_NAME); \
	 else printf '[ERROR] Could not find program "$(PKG_NAME)".\n'; \
	 fi

## Builds the program specified by the main package.
build:
	@printf "Building... "
	@GOBUILD_OUT="$$(go build ./... 2>&1)"; \
		if [ -n "$$GOBUILD_OUT" ]; then \
		  printf "\n[ERROR] Failed to build program:\n"; \
		  printf "$$GOBUILD_OUT\n"; \
		  exit 1; \
		else printf "done.\n"; \
		fi

install:
	@printf 'Installing... '
	@GOINSTALL_OUT="$$(go install 2>&1)"; \
		if [ -n "$$GOBUILD_OUT" ]; then \
		  printf "\n[ERROR] failed to install:\n"; \
		  printf "$$GOINSTALL_OUT\n"; \
		  exit 1; \
		else printf "done.\n"; \
		fi

## Installs dependencies.
get:
	@printf "Installing dependencies... "
	@GOGET_OUT="$$(go get ./... 2>&1)"; \
		if [ -n "$$GOGET_OUT" ]; then printf "\n$$GOGET_OUT\n"; \
		else printf "done.\n"; \
		fi

## Installs and updates package dependencies.
update:
	@printf 'Installing and updating package dependencies with "go get"... '
	@GOGET_OUT="$$(go get -u 2>&1)"; \
		if [ -n "$$GOGET_OUT" ]; then \
		  printf "\n[ERROR] Failed to install package dependencies:\n"; \
		  printf "$$GOGET_OUT\n"; \
		  exit 1; \
		else printf "done.\n"; \
		fi

## Fixes Go code using "go fix"
fix:
	@printf 'Fixing Go code with "go fix"... '
	@GOFIX_OUT="$$(go fix 2>&1)"; \
		if [ -n "$$GOFIX_OUT" ]; then \
		  printf "\n$$GOFIX_OUT\n"; \
		else printf "done.\n"; \
		fi

## Verifies that Go module dependencies are satisfied.
verify:
	@printf "Verifying Go module dependencies:\n"
	@go mod verify


## Formats, checks, and tests the code.
_review_base: verify fmt check
review: _review_base test
review-v: _review_base test-v
## Like "review", but tests for race conditions.
review-race: _review_base test-race
review-race-v: _review_base test-race-v
## Like "review-race", but includes benchmarks.
review-bench: review-race bench
review-bench-v: review-race bench-v


## Checks for formatting, linting, and suspicious code.
check:
## Check formatting...
	@printf "Check fmt...                 "
	@GOFMT_OUT="$$(gofmt -l $(SRC_FILES) 2>&1)"; \
		if [ -n "$$GOFMT_OUT" ]; then \
		  printf '\n[WARN] Fix formatting issues in the following files with \
"make fmt":\n'; \
		  printf "$$GOFMT_OUT\n"; \
		  exit 1; \
		else printf "ok\n"; \
		fi
## Lint files...
	@printf "Check lint...                "
	@GOLINT_OUT="$(for PKG in "$(SRC_PKGS)"; do golint $$PKG 2>&1; done)"; \
		if [ -n "$$GOLINT_OUT" ]; then \
		  printf "\n"; \
		  for PKG in "$$GOLINT_OUT"; do \
		    printf "$$PKG\n"; \
		  done; \
		  printf "\n"; \
		  exit 1; \
		else printf "ok\n"; \
		fi
## Check suspicious code...
	@printf "Check vet...                 "
	@GOVET_OUT="$$(go vet 2>&1)"; \
		if [ -n "$$GOVET_OUT" ]; then \
		  printf '\n[WARN] Fix suspicious code from "go vet":\n'; \
		  printf "$$GOVET_OUT\n"; \
		  exit 1; \
		else printf "ok\n"; \
		fi

## Reformats code according to "gofmt".
fmt:
	@printf "Formatting source files...   "
	@GOFMT_OUT="$$(gofmt -l -s -w $(SRC_FILES) 2>&1)"; \
	 if [ -n "$$GOFMT_OUT" ]; then \
	 	printf "\n$$GOFT_OUT\n"; \
	 	exit 1; \
	 else printf "ok\n"; \
     fi;

## Testing commands:
GOTEST = go test ./... -coverprofile=$(COVER_OUT) \
		               -covermode=atomic \
		               -timeout=$(TEST_TIMEOUT)
test:
	@printf "Testing:\n"
	@$(GOTEST)
test-v:
	@printf "Testing (verbose):\n"
	@$(GOTEST) -v

GOTEST_RACE = $(GOTEST) -race
test-race:
	@printf "Testing (race):\n"
	@$(GOTEST_RACE)
test-race-v:
	@printf "Testing (race, verbose):\n"
	@$(GOTEST_RACE) -v

GOBENCH = $(GOTEST) ./... -run=^$ -bench=. -benchmem
bench:
	@printf "Benchmarking:\n"
	@$(GOBENCH)
bench-v:
	@printf "Benchmarking (verbose):\n"
	@$(GOBENCH) -v
