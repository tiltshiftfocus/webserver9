package Web

import (
	"net/http"
	"strings"

	"webserver/system/basecontroller"

	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Path   string
	Action string
}

func Router() *router {
	muxRouter := mux.NewRouter()
	instance := &router{muxRouter, mainRouteHandlerType{muxRouter, paramtype{}}}
	sys := systeminfo{muxRouter}
	instance.mrht.muxRouteExactly("/system/information", sys.pageShowRouteInfoHandler)
	return instance
}

type router struct {
	muxRouter *mux.Router
	mrht      mainRouteHandlerType
}

func (self *router) Route(routeConfig interface{}, icontroller interface{}) {
	var rc []RouteConfig
	switch routeConfig.(type) {
	case RouteConfig:
		rc = []RouteConfig{routeConfig.(RouteConfig)}
	case []RouteConfig:
		rc = routeConfig.([]RouteConfig)
	default:
		errorLog("Not yet support %T", routeConfig)
		return
	}
	for _, row := range rc {
		path := row.Path
		action := row.Action
		if !isMethodExist(&icontroller, action) {
			errorLog("Web.RouteConfig, path:%s controller:%T action:%s not found! ", path, icontroller, action)
		}
		if !isFieldExist(&icontroller, "Response") || !isFieldExist(&icontroller, "Request") {
			errorLog("Web.RouteConfig, controller:%T missing 'BaseController' ", icontroller)
		}
		get, post := retrieveMethodParams(&icontroller, action)
		self.mrht.addToRoute(path, icontroller, post, get, action)
	}
}
func (self *router) RouteExactly(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return self.mrht.muxRouteExactly(path, f)
}
func (self *router) RouteByController(path string, icontroller interface{}) {
	if !isFieldExist(&icontroller, "Response") || !isFieldExist(&icontroller, "Request") {
		errorLog("Web.RouteConfig, controller:%T missing 'BaseController' ", icontroller)
	}
	baseMethods := listAllMethods(new(basecontroller.BaseController))
	skipMethods := map[string]int{}
	for _, v := range baseMethods {
		skipMethods[v] = 1
	}
	methods := listAllMethods(icontroller)
	rc := []RouteConfig{}
	for _, name := range methods {
		if _, has := skipMethods[name]; has {
			continue
		}
		lowcase := path + "/" + strings.ToLower(name)
		rc = append(rc, RouteConfig{Path: lowcase, Action: name})
	}
	self.Route(rc, icontroller)
}
func (self *router) GetRouter() *mux.Router { return self.muxRouter }
func (self *router) AllowDomains(domains []string) {
	if len(domains) > 0 {
		self.muxRouter.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for idx := range domains {
					domain := &domains[idx]
					host := r.Host
					if i := strings.Index(host, ":"); i != -1 {
						host = host[:i]
					}
					if host == *domain {
						next.ServeHTTP(w, r)
						return
					}
				}
				http.Error(w, "Forbidden", http.StatusForbidden)
			})
		})
	}
}
func (self *router) SupportParameters(in ...interface{}) {
	self.mrht.pt.Process(in...)
	//self.mrht.pt.display()
}
