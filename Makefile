.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o vue.wasm ./component
