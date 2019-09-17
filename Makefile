.PHONY: run

dev: dev-bindata
	go run main.go

dev-bindata:
	go-bindata -o template/bindata.go -ignore=\\.DS_Store -ignore=\\.gitkeep -debug=1 -pkg template -prefix view/template view/template/...
	go-bindata -o router/bindata.go -ignore=\\.DS_Store -ignore=\\.gitkeep -debug=1 -pkg router -prefix view/ view/assets/...

generate:
	go generate

build: generate
	go build -o build/build *.go
