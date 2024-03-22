package controllers

import (
	"html/template"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bnhf/go-openvpn/server/config"

	"github.com/Glaolle/openvpn-admin-plus/lib"
	"github.com/Glaolle/openvpn-admin-plus/models"
)

type OVConfigController struct {
	BaseController
}

func (c *OVConfigController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
	c.Data["breadcrumbs"] = &BreadCrumbs{
		Title: "OpenVPN config",
	}
}

func (c *OVConfigController) Get() {
	c.TplName = "ovconfig.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	cfg := models.OVConfig{Profile: "default"}
	cfg.Read("Profile")
	c.Data["Settings"] = &cfg

}

func (c *OVConfigController) Post() {
	c.TplName = "ovconfig.html"
	flash := beego.NewFlash()
	cfg := models.OVConfig{Profile: "default"}
	cfg.Read("Profile")
	if err := c.ParseForm(&cfg); err != nil {
		logs.Warning(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}
	lib.Dump(cfg)
	c.Data["Settings"] = &cfg

	destPath := models.GlobalCfg.OVConfigPath + "/" + os.Getenv("PIVPN_CONF")
	err := config.SaveToFile("conf/openvpn-server-config.tpl", cfg.Config, destPath)
	if err != nil {
		logs.Warning(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	// Attempts to restart OpenVPN service after config update
	// below caused issues in bridged (TAP) setups
	o := orm.NewOrm()
	if _, err := o.Update(&cfg); err != nil {
		flash.Error(err.Error())
	} else {
		flash.Success("Config has been updated -- reboot required for changes to take effect!")
		//		client := mi.NewClient(models.GlobalCfg.MINetwork, models.GlobalCfg.MIAddress)
		//		if err := client.Signal("SIGTERM"); err != nil {
		//if err != nil { //Above two lines causing MI to stop responding
		//	flash.Warning("Config has been updated but OpenVPN server was NOT reloaded: " + err.Error())
		//}
	}
	flash.Store(&c.Controller)
}
