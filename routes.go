package main

import (
	 mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Handle	    mux.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"PostIndex",
		"GET",
		"/posts",
		PostIndex,
	},
	Route{
		"PostShow",
		"GET",
		"/posts/:postId",
		PostShow,
	},
	Route{
		"PostCreate",
		"POST",
		"/posts",
		PostCreate,
	},
	Route{
		"PostDelete",
		"POST",
		"/posts/del/:postId",
		PostDelete,
	},
}
