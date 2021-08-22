package static

import (
	"embed"
)

//go:embed index.html main.html wasm_exec.js client.wasm
var FS embed.FS
