package router

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/yenchieh/go-binary-assets/tmpl"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	if err := setTemplates(router, "home.html", "head.html", "body.html", "common.html", "products/backlog.html", "products/cacoo.html", "products/typetalk.html"); err != nil {
		log.Fatal(err)
	}

	assetsFS := fs("assets")
	router.StaticFS("/assets", assetsFS)

	router.GET("/", homePage)
	router.GET("/product/:productName", loadHTMLPage)

	return router
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Version": "1.0.0",
	})
}

func loadHTMLPage(c *gin.Context) {
	productName := c.Param("productName")
	c.HTML(http.StatusOK, fmt.Sprintf("products/%s.html", productName), gin.H{
		"Version":  "1.0.0",
		"PageName": productName,
	})
}

func fs(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}

func setTemplates(r *gin.Engine, fileNames ...string) error {
	t := template.New("")
	err := errors.New("")
	for _, fileName := range fileNames {
		indexBytes := tmpl.MustAsset(fileName)
		t, err = t.New(fileName).Parse(string(indexBytes))
		if err != nil {
			return err
		}
	}
	r.SetHTMLTemplate(t)

	return nil
}
