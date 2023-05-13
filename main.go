package main

import (
	"kz/mo"
	"net/http"
)

func main() {
	engine := mo.New()
	engine.GET("/hello", func(c *mo.Context) {
		c.String(200, "%s", "hello")
	})
	engine.GET("/login", func(ctx *mo.Context) {
		ctx.JSON(http.StatusOK, mo.H{
			"name":   "heyishi",
			"passwd": "12345",
		})
	})
	engine.GET("/html", func(ctx *mo.Context) {
		ctx.HTML(200, "<h>hello</h>")
	})
	engine.Run(":9090")
}
