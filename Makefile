MODULE ?= $(shell go list -m)

TARGETOS ?= $(shell go env GOOS)
TARGETARCH ?= $(shell go env GOARCH)

DIST_DIR ?= dist
BIN_SUFFIX :=
ifeq ($(TARGETOS),windows)
	BIN_SUFFIX := .exe
endif

CGO_ENABLED ?= 1 # only used to compile server

##@ Development
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "${DIST_DIR}/server${BIN_SUFFIX}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go,tpl,tmpl,html,css,scss,js,ts,sql,jpeg,jpg,gif,png,bmp,svg,webp,ico" \
		--misc.clean_on_exit "true"

##@ Build
build-server: ## Build server
	CGO_ENABLED=${CGO_ENABLED} GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ${DIST_DIR}/server${BIN_SUFFIX} .

.PHONY: build
build: build-server ## Build all binaries
