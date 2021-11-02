package main

/* TODO
* add PathPrefix
* /system/information display more information
* View template
* WebSocket design
* DB design
 */

import (
	"fmt"
	"net/http"
	"time"

	"jjwebserver/controller"
	"jjwebserver/parameter"
	Web "jjwebserver/system/Web"
)

type RC = Web.RouteConfig

func main() {
	web := Web.Router()
	web.AllowDomains([]string{"test1.grannygame.io", "52.77.146.102"})
	web.SupportParameters(
		new(parameter.Username),
		new(parameter.Password))

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
