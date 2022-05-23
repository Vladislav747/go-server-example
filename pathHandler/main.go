package main

import "net/http"

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye", goodbye)
	http.ListenAndServe("8080", pr)
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ni
}
