package controllers

import (
	"log"
	"myproject/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) ShowLogin() {
	beego.Info("hello world")
	c.TplName = "login.tpl"
}

func (c *AuthController) HandleLogin() {
	//1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("输入数据不合法")
		c.TplName = "login.tpl"
		return
	}
	//3.查询账号密码是否正确
	o := orm.NewOrm()
	user := models.Users{}
	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败1001")
		c.TplName = "login.tpl"
		return
	}
	if user.Password != pwd {
		beego.Info("查询失败1002")
		c.TplName = "login.tpl"
		return
	}
	c.SetSession("uid", user.Id)

	//4.跳转
	c.Ctx.WriteString("登陆成功，欢迎您")
}

func (c *AuthController) ShowRegister() {
	c.TplName = "register.tpl"
}

func (c *AuthController) HandleRegister() {
	//1.拿到数据
	userName := c.GetString("userName") //因为是从前端拿数据，这个名称需要跟前端中的名称保持一致
	pwd := c.GetString("pwd")           //因为是从前端拿数据，这个名称需要跟前端中的名称保持一致
	valid := validation.Validation{}

	//2.对数据进行校验
	valid.MinSize(userName, 1, "nameMin")
	valid.MinSize(pwd, 1, "pwdMin")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		beego.Info("数据不能为空")
		c.Redirect("/register", 302) //重定向函数，如果发生错误页面重新回到/register，并返回错误码302
		return
	}

	//3.插入数据库
	o := orm.NewOrm()
	user := models.Users{}
	user.Name = userName
	user.Password = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据库失败")
		c.Redirect("/register", 302)
		return
	}

	//4.返回登陆界面
	c.TplName = "login.tpl"   //指定视图文件，同时可以给这个视图传递一些数据如在c.Data["errmsg"]，优点就是能够传递数据
	c.Redirect("/login", 302) //跳转，不能传递数据。优点是速度快
}
