GO=go

.PHONY: server
server: bin static
	$(GO) build -v -o bin/server ./cmd/server

bin:
	mkdir bin

.PHONY: static
static: static/wasm_exec.js client

static/wasm_exec.js:
	cp $(shell $(GO) env GOROOT)/misc/wasm/wasm_exec.js static

.PHONY: client
client:
	GOOS=js GOARCH=wasm $(GO) build -v -o static/client.wasm ./cmd/client

.PHONY: dependency
dependency:
	$(GO) mod tidy && $(GO) mod vendor