package main

import (
	"github.com/Glaolle/openvpn-admin-plus/lib"
	_ "github.com/Glaolle/openvpn-admin-plus/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	lib.AddFuncMaps()
	beego.Run()
}
