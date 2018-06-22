package routers

import (
	"study-go/study-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Index")
	beego.Router("/user",&controllers.UserController{})
}
