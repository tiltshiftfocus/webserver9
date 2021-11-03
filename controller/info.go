package controller

import (
	_ "strconv"
	"webserver/parameter"
	. "webserver/system/basecontroller"
)

type Info struct{ BaseController }

func (self *Info) Info(arg struct {
	GET_name      *parameter.Username
	POST_name     *parameter.Username
	POST_password *parameter.Password
}) {

	if arg.GET_name != nil {
		if arg.GET_name.Valid {
			self.Echo("GET name => " + arg.GET_name.Value + "\n")
		} else {
			self.Echo("Get name is invalid\n")
		}
	}
	if arg.POST_name != nil {
		if arg.POST_name.Valid {
			self.Echo("POST name => " + arg.POST_name.Value + "\n")
		} else {
			self.Echo("Get name is invalid\n")
		}
	}
}

func (self *Info) Info2(arg struct {
	GET_name  parameter.Username
	POST_name parameter.Username
}) {
	self.Echo("GET name => " + arg.GET_name.Value + "\n")
	self.Echo("POST name => " + arg.POST_name.Value + "\n")
}
