.PHONY: run

dev: dev-bindata
	go run main.go

dev-bindata:
	go-bindata -o tmpl/bindata.go -ignore=\\.DS_Store -ignore=\\.gitkeep -debug=1 -pkg tmpl -prefix view/template view/template/...
	go-bindata -o router/bindata.go -ignore=\\.DS_Store -ignore=\\.gitkeep -debug=1 -pkg router -prefix view/ view/assets/...

generate:
	go generate

build: generate
	GOOS=linux GOARCH=amd64 \
	go build -o build/binary *.go

build-image: build
	docker build -t yenchieh/bowery-golang-bindata-demo:latest .

push-image: build-image
	docker push yenchieh/bowery-golang-bindata-demo:latest
