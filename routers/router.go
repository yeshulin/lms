package routers

import (
	"webproject/lms/controllers"
	"webproject/lms/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &admin.UserController{}, "get:Login")
	beego.Router("/user/login", &admin.UserController{}, "post:Post")
	beego.Router("/user/reg", &admin.UserController{}, "post:Reg")
	beego.Router("/register", &admin.UserController{}, "get:Register")
	beego.Router("/admin", &admin.IndexController{})
	beego.Router("/admin/user/find", &admin.UserController{}, "get:Get")
	beego.Router("/admin/userlist", &admin.UserController{}, "get:UserList")
	beego.Router("/admin/user/add", &admin.UserController{}, "get:Add;post:Add")
	beego.Router("/admin/user/addpost", &admin.UserController{}, "post:AddPost")
	beego.Router("/admin/user/delete", &admin.UserController{}, "post:Delete")
	beego.Router("/admin/user/view", &admin.UserController{}, "get:View")
	beego.Router("/admin/user/edit", &admin.UserController{}, "get:Edit")
	beego.Router("/admin/user/editpost", &admin.UserController{}, "post:EditPost")
	beego.Router("/admin/role", &admin.RoleController{}, "get:RoleList")
	beego.Router("/admin/role/find", &admin.RoleController{}, "get:Get")
	beego.Router("/admin/role/add", &admin.RoleController{}, "get:Add")
	beego.Router("/admin/role/addpost", &admin.RoleController{}, "post:AddPost")
	beego.Router("/admin/role/delete", &admin.RoleController{}, "post:Delete")
	beego.Router("/admin/role/view", &admin.RoleController{}, "get:View")
	beego.Router("/admin/role/edit", &admin.RoleController{}, "get:Edit")
	beego.Router("/admin/role/editpost", &admin.RoleController{}, "post:EditPost")
	beego.Router("/admin/node", &admin.NodeController{}, "get:NodeList")
	beego.Router("/admin/node/find", &admin.NodeController{}, "get:Get")
	beego.Router("/admin/node/add", &admin.NodeController{}, "get:Add")
	beego.Router("/admin/node/addpost", &admin.NodeController{}, "post:AddPost")
	beego.Router("/admin/node/delete", &admin.NodeController{}, "post:Delete")
	beego.Router("/admin/node/view", &admin.NodeController{}, "get:View")
	beego.Router("/admin/node/edit", &admin.NodeController{}, "get:Edit")
	beego.Router("/admin/node/editpost", &admin.NodeController{}, "post:EditPost")
	beego.Router("/admin/menu", &admin.MenuController{})
	beego.Router("/admin/news", &admin.NewsController{})
	beego.Router("/admin/newstype", &admin.NewsTypeController{})
	beego.Router("/admin/course", &admin.CourseController{})
	beego.Router("/admin/coursetype", &admin.CourseTypeController{})
	beego.Router("/admin/index/welcome", &admin.IndexController{}, "get:Welcome")
}
