package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get(){
	c.Data["Message"] = "Hello world!"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName="user/index.tpl"
}
