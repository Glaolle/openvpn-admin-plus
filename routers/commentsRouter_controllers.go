package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISessionController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISessionController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISessionController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISessionController"],
		beego.ControllerComments{
			Method:           "Kill",
			Router:           "/",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISignalController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISignalController"],
		beego.ControllerComments{
			Method:           "Send",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISysloadController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:APISysloadController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/certificates",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/certificates",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"],
		beego.ControllerComments{
			Method:           "Download",
			Router:           "/certificates/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"],
		beego.ControllerComments{
			Method:           "Remove",
			Router:           "/certificates/remove/:key/:serial",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"],
		beego.ControllerComments{
			Method:           "Revoke",
			Router:           "/certificates/revoke/:key/:serial",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/Glaolle/openvpn-admin-plus/controllers:CertificatesController"],
		beego.ControllerComments{
			Method:           "DownloadSingleConfig",
			Router:           "/certificates/single-config/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
