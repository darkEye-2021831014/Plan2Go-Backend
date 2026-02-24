package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	GlobalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		GlobalMiddlewares: make([]Middleware, 0),
	}
}
func (mngr *Manager) Use(middlewares ...Middleware)  {
	mngr.GlobalMiddlewares = append(mngr.GlobalMiddlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, middlewares ...Middleware)http.Handler{
	n:=next
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	
	//for global middlewares
	for _, middleware := range m.GlobalMiddlewares {
		n = middleware(n)
	}
	return n
}