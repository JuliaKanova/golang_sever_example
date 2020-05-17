package router

import (
	"net/http"

	"../storages"
)

// Router ...
type Router struct {
	rootHandler rootHandler
}

// New ...
func New(store storages.Store) *Router {
	return &Router{
		rootHandler: newRootHandler(store),
	}
}

// RootHandler ...
func (r *Router) RootHandler() http.Handler {
	return r.rootHandler
}
