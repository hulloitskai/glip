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
.PHONY: default run build install get update review review-race review-bench \
		check fmt test test-v test-race bench

## Default target when no arguments are given to make (build and run program).
default: build run

## Builds and runs the program (package must be main).
run:
	@if [ -f ".env.sh" ]; then \
	   printf 'Exporting environment variables by sourcing ".env.sh"... '; \
	   . .env.sh; \
	   echo "done."; \
	 fi
	@if [ -f "$(PKG_NAME)" ]; then \
	   echo 'Running "$(PKG_NAME)"...'; \
	   ./$(PKG_NAME); \
	 else echo '[ERROR] Could not find program "$(PKG_NAME)".'; \
	 fi

## Builds the program specified by the main package.
build:
	@printf "Building... "
	@GOBUILD_OUT="$$(go build 2>&1)"; \
		if [ -n "$$GOBUILD_OUT" ]; then \
		  echo "\n[ERROR] Failed to build program:"; \
		  echo $$GOBUILD_OUT; \
		  exit 1; \
		else echo "done."; \
		fi

install:
	@printf 'Installing with "go install"... '
	@GOINSTALL_OUT="$$(go install 2>&1)"; \
		if [ -n "$$GOBUILD_OUT" ]; then \
		  echo "\n[ERROR] failed to install:"; \
		  echo "$$GOINSTALL_OUT"; \
		  exit 1; \
		else echo "done."; \
		fi


## Installs package dependencies.
get:
	@printf 'Installing package dependencies with "go get"... '
	@GOGET_OUT="$$(go get ./... 2>&1)"; \
		if [ -n "$$GOGET_OUT" ]; then \
		  echo "\n[ERROR] Failed to install package dependencies:"; \
		  echo "$$GOGET_OUT"; \
		  exit 1; \
		else echo "done."; \
		fi

## Installs and updates package dependencies.
update:
	@printf 'Installing and updating package dependencies with "go get"... '
	@GOGET_OUT="$$(go get -u 2>&1)"; \
			if [ -n "$$GOGET_OUT" ]; then \
		  echo "\n[ERROR] Failed to install package dependencies:"; \
		  echo "$$GOGET_OUT"; \
		  exit 1; \
		else echo "done."; \
		fi


## Formats, checks, and tests the code.
review: fmt check test
review-v: fmt check test-v
## Like "review", but tests for race conditions.
review-race: fmt check test-race
review-race-v: fmt check test-race-v
## Like "review-race", but includes benchmarks.
review-bench: review-race bench
review-bench-v: review-race bench-v


## Checks for formatting, linting, and suspicious code.
check:
## Check formatting...
	@printf "Check fmt...                 "
	@GOFMT_OUT="$$(gofmt -l $(SRC_FILES) 2>&1)"; \
		if [ -n "$$GOFMT_OUT" ]; then \
		  echo '\n[WARN] Fix formatting issues in the following files with \
"make fmt":'; \
		  echo "$$GOFMT_OUT\n"; \
		  exit 1; \
		else echo "ok"; \
		fi
## Lint files...
	@printf "Check lint...                "
	@GOLINT_OUT="$(for PKG in "$(SRC_PKGS)"; do golint $$PKG 2>&1; done)"; \
		if [ -n "$$GOLINT_OUT" ]; then \
		  printf "\n"; \
		  for PKG in "$$GOLINT_OUT"; do \
		    echo "$$PKG"; \
		  done; \
		  printf "\n"; \
		  exit 1; \
		else echo "ok"; \
		fi
## Check suspicious code...
	@printf "Check vet...                 "
	@GOVET_OUT="$$(go vet 2>&1)"; \
		if [ -n "$$GOVET_OUT" ]; then \
		  echo '\n[WARN] Fix suspicious code from "go vet":'; \
		  echo "$$GOVET_OUT\n"; \
		  exit 1; \
		else echo "ok"; \
		fi

## Reformats code according to "gofmt".
fmt:
	@printf "Formatting source files...   "
	@GOFMT_OUT="$$(gofmt -l -s -w $(SRC_FILES) 2>&1)"; \
	 if [ -n "$$GOFMT_OUT" ]; then \
	 	printf "\n$$GOFT_OUT\n"; \
	 	exit 1; \
	 else echo "ok"; \
     fi;

## Testing commands:
GOTEST = go test ./... -coverprofile=$(COVER_OUT) \
		               -covermode=atomic \
		               -timeout=$(TEST_TIMEOUT)
test:
	@echo "Testing:"
	@$(GOTEST)
test-v:
	@echo "Testing (verbose):"
	@$(GOTEST) -v

GOTEST_RACE = $(GOTEST) -race
test-race:
	@echo "Testing (race):"
	@$(GOTEST_RACE)
test-race-v:
	@echo "Testing (race, verbose):"
	@$(GOTEST_RACE) -v

GOBENCH = $(GOTEST) ./... -run=^$ -bench=. -benchmem
bench:
	@echo "Benchmarking:"
	@$(GOBENCH)
bench-v:
	@echo "Benchmarking (verbose):"
	@$(GOBENCH) -v
