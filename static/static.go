package static

import (
	"embed"
)

//go:embed main.html wasm_exec.js client.wasm
var Static embed.FS

//go:embed index.html
var Index []byte
