package gee

import (
    "fmt"
    "net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
    // 存放路由和对应的处理函数
    router map[string]HandleFunc
}

// 构造方法
function New() *Engine {
    // 返回一个engine的指针，里面路由信息是一个切片
    return &Engine{router: make(map[string]HandleFunc)}
}

// 添加路由信息
func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
    key := method + "-" + pattern
    engine.router[key] = handler
}

// GET处理方法
func (engine *Engine) GET(pattern string, handler HandleFunc) {
    engine.addRoute("GET", pattern, handler)
}

// POST处理方法
func (engine *Engine) POST(pattern string, handler HandleFunc) {
    engine.addRoute("POST", pattern, handler)
}

// run server
func (engine *Engine) Run(addr string) (err error) {
    return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHttp(w http.ResponseWriter, req *http.Request) {
    key := req.Method + "-" + req.URL.Path
    if handler, ok := engine.router[key]; ok {
        handler(w, req)
    } else {
        fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
    }
}
