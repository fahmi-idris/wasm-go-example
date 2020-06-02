package main

import (
	"github.com/fahmi-idris/wasm-go-example/counter"
)

func main() {
	counter.VueCounter("app")
	select {}
}
