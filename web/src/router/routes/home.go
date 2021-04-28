package routes

import (
	"net/http"
	"web/src/views"
)

var homeRoute = Route{
	URI:          "/home",
	Method:       http.MethodGet,
	Function:     views.HomePage,
	AuthRequired: true,
}
