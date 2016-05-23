package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"image/gif"
	"io"
	"os"
)

var cache map[string]*gif.GIF

func serveGif(c *gin.Context, g *gif.GIF) {
	b := new(bytes.Buffer)
	w := io.Writer(b)
	gif.EncodeAll(w, g)
	c.Data(200, "image/gif", b.Bytes())
}

func previewHandler(c *gin.Context) {
	key := c.Param("key")
	g, ok := cache[key]
	if ok == true {
		fmt.Printf("Found cache for: %v\n", key)
		serveGif(c, g)
	}

	f, err := os.Open(key)
	if err != nil {
		fmt.Errorf("Cannot find gif %v\n", err)
	}

	g, err = gif.DecodeAll(f)
	if err != nil {
		fmt.Errorf("Cannot read gif %v\n", err)
	}

	serveGif(c, g)
	fmt.Printf("Set cache for: %v\n", key)
	cache[key] = g
}

func main() {
	cache = make(map[string]*gif.GIF)
	r := gin.Default()
	r.GET("/preview/:key", previewHandler)
	r.Run()
}
