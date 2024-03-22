package controllers

import (
	"github.com/Glaolle/openvpn-admin-plus/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller

	Userinfo *models.User
	IsLogin  bool
}

type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

func (c *BaseController) Prepare() {
	c.SetParams()

	c.IsLogin = c.GetSession("userinfo") != nil
	if c.IsLogin {
		c.Userinfo = c.GetLogin()
	}

	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.Userinfo

	//c.Data["HeadStyles"] = []string{}
	//c.Data["HeadScripts"] = []string{}

	//c.Layout = "base.tpl"
	//c.LayoutSections = make(map[string]string)
	//c.LayoutSections["BaseHeader"] = "header.tpl"
	//c.LayoutSections["BaseFooter"] = "footer.tpl"

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *BaseController) GetLogin() *models.User {
	u := &models.User{Id: c.GetSession("userinfo").(int64)}
	u.Read()
	return u
}

func (c *BaseController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *BaseController) SetLogin(user *models.User) {
	c.SetSession("userinfo", user.Id)
}

func (c *BaseController) LoginPath() string {
	return c.URLFor("LoginController.Login")
}

func (c *BaseController) SetParams() {
	c.Data["Params"] = make(map[string]string)
	xInput, err := c.Input()
	if err != nil {
		logs.Error(err)
	}
	for k, v := range xInput {
		c.Data["Params"].(map[string]string)[k] = v[0]
	}
}

type BreadCrumbs struct {
	Title    string
	Subtitle string
}
