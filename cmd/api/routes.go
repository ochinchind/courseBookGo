package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler) 


	router.HandlerFunc(http.MethodGet, "/v1/courses", app.requirePermission("courses:read", app.listCoursesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/courses", app.requirePermission("courses:write", app.createCourseHandler))
	router.HandlerFunc(http.MethodGet, "/v1/courses/:id", app.requirePermission("courses:read", app.showCourseHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/courses/:id", app.requirePermission("courses:write", app.updateCourseHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/courses/:id", app.requirePermission("courses:write", app.deleteCourseHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
