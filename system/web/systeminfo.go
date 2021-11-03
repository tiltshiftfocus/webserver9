package Web
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strings"
)

type systeminfo struct {
	r *mux.Router
}
func (s *systeminfo) pageShowRouteInfoHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || !(user == "user" && pass == "1234") {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized login", http.StatusUnauthorized)
		return
	}
	_, paths := s.getRouteInfo()
	w.Write([]byte(`
		<html>
		<body>
	`));
	w.Write([]byte(paths));
	w.Write([]byte(`
		</body>
		</html>
	`));
}
func (s *systeminfo) getRouteInfo() (string,string) {
	var msg string
	var paths string
	err := s.r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			paths += fmt.Sprintf("<div>%s</div>", pathTemplate)
			msg += fmt.Sprintln("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			msg += fmt.Sprintln("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			msg += fmt.Sprintln("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			msg += fmt.Sprintln("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			msg += fmt.Sprintln("Methods:", strings.Join(methods, ","))
		}
		return nil
	})
	_ = err
	return msg, paths
}
