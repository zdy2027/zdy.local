// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"zdy.local/cecetl/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
		beego.NSNamespace("/producter",
			beego.NSInclude(
				&controllers.ProducterController{},
			),
		),
		beego.NSNamespace("/consumer",
			beego.NSInclude(
				&controllers.DownLoadController{},
			),
		),
		beego.NSNamespace("/autosql",
			beego.NSInclude(
				&controllers.AutoSqlController{},
			),
		),
		beego.NSNamespace("/files",
			beego.NSInclude(
				&controllers.FileUploadController{},
			),
		),
		beego.NSNamespace("/utils",
			beego.NSInclude(
				&controllers.EsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
