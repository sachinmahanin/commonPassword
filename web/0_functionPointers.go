package web

import "github.com/sachinmahanin/commonPassword/handler/miscellaneous"

// func pointers for injection / testing: web.go
var (
	miscellaneousSwaggerHandler  = miscellaneous.SwaggerHandler
	miscellaneousSwaggerRedirect = miscellaneous.SwaggerRedirect
	registeredUtilityRoutesFunc  = registeredUtilityRoutes
	registeredBusinessRoutesFunc = registeredBusinessRoutes
)
