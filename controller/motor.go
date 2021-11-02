package controller

import (
	"fmt"
	"strconv"
	_ "time"

	. "jjwebserver/system/basecontroller"
)

type Motor struct {
	BaseController
	count int
}

func (u *Motor) Index() {
	u.Echo(`
	<html>
	<body>
		<div style="color:red">welcome</div>
	</body>
	</html>
	`)
}

func (u *Motor) Testing1(args struct {
	GET_name, GET_who   string
	POST_name, POST_who int
}) {
	u.Echo(args.GET_name)
	u.Echo(args.GET_who)
}

func (u *Motor) TestGetPost(args struct {
	GET_name string
	GET_who  string
	POST_id  int
}) {
	u.Echo(args.GET_name)
	u.Echo(args.GET_who)
	u.Echo(strconv.Itoa(args.POST_id))
}

func (u *Motor) TestPost(args struct {
	POST_name string
	POST_who  string
}) {
	u.Echo(args.POST_name)
	u.Echo(args.POST_who)
}
func (u *Motor) Process(args struct {
	GET_id   int
	GET_id2  int
	GET_name string
}) {
	fmt.Println("========= Process =============")
	u.count++
	u.Echo("starting\n")
	fmt.Println("get:", args.GET_id, " get:", args.GET_id2, " get:", args.GET_name, " count:", u.count)
	//msg := "count:" + string(u.count) + "\n";
	//for i := 0; i < 5; i++ {
	//	u.Echo(msg)
	//	fmt.Println(msg)
	//	time.Sleep(1 * time.Second)
	//}
}
func (u *Motor) Process1(args struct{ GET_id int }) {
	fmt.Println("========= Process1 =============")
	u.count++
	fmt.Println("get:", args.GET_id, " count:", u.count)
}
