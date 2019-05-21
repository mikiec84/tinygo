package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	updateCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		update()
		return nil
	})
	document.Call("getElementById", "a").Set("oninput", updateCallback)
	document.Call("getElementById", "b").Set("oninput", updateCallback)
	update()
}

//go:export add
func add(a, b int) int {
	return a + b
}

func update() {
	document := js.Global().Get("document")
	aStr := document.Call("getElementById", "a").Get("value").String()
	bStr := document.Call("getElementById", "b").Get("value").String()
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)
	result := add(a, b)
	document.Call("getElementById", "result").Set("value", result)
}
