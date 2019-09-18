//go:generate go-bindata -o tmpl/bindata.go -ignore=\\.DS_Store -ignore=\\.gitkeep -pkg tmpl -prefix view/template view/template/...
//go:generate go-bindata -o router/bindata.go -ignore=\\.DS_Store -ignore=\\.gitkeep -pkg router -prefix view/ view/assets/...

package main

import (
	"log"

	"github.com/yenchieh/go-binary-assets/router"
)

func main() {

	r := router.New()
	port := ":8081"
	log.Printf("Running on port %s", port)

	if err := r.Run(port); err != nil {
		log.Fatal(err)

	}
}
