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
	beego.Router("/admin/role", &admin.RoleController{})
	beego.Router("/admin/menu", &admin.MenuController{})
	beego.Router("/admin/news", &admin.NewsController{})
	beego.Router("/admin/newstype", &admin.NewsTypeController{})
	beego.Router("/admin/course", &admin.CourseController{})
	beego.Router("/admin/coursetype", &admin.CourseTypeController{})
	beego.Router("/admin/index/welcome", &admin.IndexController{}, "get:Welcome")
}
