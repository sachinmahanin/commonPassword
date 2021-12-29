package web

import (
	"net/http"

	"github.com/sachinmahanin/commonpassword/handler/business"
	"github.com/sachinmahanin/commonpassword/handler/utility"
	webserver "github.com/zhongjie-cai/web-server"
)

//RegisteredStatics returns the registered static content handlers for web service hosting
func RegisteredStatics() []webserver.Static {
	return []webserver.Static{
		//add static routes
	}
}

func registeredBusinessRoutes() []webserver.Route {
	return []webserver.Route{
		webserver.Route{
			Endpoint:   "business.SearchPassword",
			Method:     http.MethodPost,
			Path:       "/CommonPassword",
			ActionFunc: business.SearchPassword,
		},
	}
}

func registeredUtilityRoutes() []webserver.Route {
	return []webserver.Route{
		webserver.Route{
			Endpoint:   "utility.Health",
			Method:     http.MethodGet,
			Path:       "/health",
			ActionFunc: utility.Health,
		},
	}
}

// RegisteredRoutes returns the registered route handlers for web service hosting
func RegisteredRoutes() []webserver.Route {
	var allRoutes = []webserver.Route{}
	allRoutes = append(
		allRoutes,
		registeredUtilityRoutesFunc()...,
	)
	allRoutes = append(
		allRoutes,
		registeredBusinessRoutesFunc()...,
	)
	return allRoutes
}
