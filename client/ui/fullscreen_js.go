package ui

import "syscall/js"

func RequestFullscreen() {
	requestFullscreen := js.Global().Get("document").Get("documentElement").Get("requestFullscreen")
	js.Global().Get("document").Get("body").Set("onclick", requestFullscreen)
}

func ExitFullscreen() {
	exitFullscreen := js.Global().Get("document").Get("exitFullscreen")
	if exitFullscreen.Truthy() {
		exitFullscreen.Invoke()
	}
}
