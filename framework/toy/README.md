### 使用示例
```go
package main

import (
	"toy"
	"net/http"
)

func main() {
	r := toy.Default()
	r.GET("/", func(c *toy.Context) {
		c.HTML(http.StatusOK, "<h1>Hello toy</h1>")
	})

	r.GET("/hello", func(c *toy.Context) {
		// expect /hello?name=Stephen Curry
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *toy.Context) {
		// expect /hello/Stephen Curry
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *toy.Context) {
		c.JSON(http.StatusOK, toy.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

```