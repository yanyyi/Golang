package zin

import (
	"net/http"
	"strings"
)

type router struct {
	//使用 roots 来存储每种请求方式的Trie 树根节点, 使用 handlers 存储每种请求方式的 HandlerFunc
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// path: /p/:lang/doc 或者 /lang/*, 因此需要写一个函数获得具体路径

func (router *router) addRoute(method string, pattern string, handler HandlerFunc) {
	// 对router.handlers的处理
	key := method + "-" + pattern
	router.handlers[key] = handler
	//解析路径
	parts := parsePattern(pattern)

	// 对router.roots的处理
	//查看是否已注册方法
	if _, ok := router.roots[method]; !ok {
		router.roots[method] = &node{}
	}
	//将该方法注册到nodes里去
	router.roots[method].insert(pattern, parts, 0)
}

/*
解析了:和*两种匹配符的参数，返回一个 map 。
例如

	/p/go/doc         匹配到   /p/:lang/doc，     解析结果为：{lang: "go"}，

/static/css/geektutu.css  匹配到   /static/*filepath，解析结果为{filepath: "css/geektutu.css"}
*/

func (router *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	n := router.roots[method].search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '!' && len(parts) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

//	func (router *router) handle(c *Context) {
//		n, params := router.getRoute(c.Method, c.Path)
//		if n != nil {
//			key := c.Method + "-" + n.pattern
//			c.Params = params
//			router.handlers[key](c)
//		} else {
//			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
//		}
//	}
func (router *router) handle(c *Context) {
	n, params := router.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, router.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
