package main

/* TODO
* /system/information display more information
* View template
* WebSocket design
* DB design
 */

import (
	"fmt"
	"net/http"
	"time"
	"webserver/controller"
	"webserver/parameter"
	Web "webserver/system/web"
)

type RC = Web.RouteConfig

func main() {
	web := Web.Router()
	web.SupportParameters(
		new(parameter.Username),
		new(parameter.Password))

	web.AllowDomains([]string{"52.77.146.102"})
	web.Route([]RC{
		{"/testinfo", "Info"},
		{"/testinfo2", "Info2"},
	}, new(controller.Info))
	web.RouteByController("/info/{path}", new(controller.Info))
	web.RouteByController("/allinone", new(controller.Motor))
	web.Route(RC{Path: "/testpost", Action: "TestPost"}, new(controller.Motor))
	web.Route([]RC{
		{"/testgetpost", "TestGetPost"},
		{"/matchtest/{n:.*}", "Index"},
	}, new(controller.Motor))

	web.AllowDomains([]string{"{subdomain}.grannygame.io"})
	web.Route([]RC{
		{"/singleinfo", "Info"},
	}, new(controller.Info))

	web.AllowAllDomains()
	web.Route([]RC{
		{"/defaultinfo", "Info"},
	}, new(controller.Info))


	web.RouteExactly("/{n:.*}", Web.PageNotFoundHandler)
	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        web.GetRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Http Server...")
	server.ListenAndServe()
}
