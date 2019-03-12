package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"],
		beego.ControllerComments{
			Method: "AutoImage",
			Router: `/AutoImage`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"],
		beego.ControllerComments{
			Method: "AutoPatient",
			Router: `/AutoPatient`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"],
		beego.ControllerComments{
			Method: "AutoReport",
			Router: `/AutoReport`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"],
		beego.ControllerComments{
			Method: "AutoSeries",
			Router: `/AutoSeries`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:AutoSqlController"],
		beego.ControllerComments{
			Method: "AutoStudy",
			Router: `/AutoStudy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:DownLoadController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:DownLoadController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"],
		beego.ControllerComments{
			Method: "DelWeed",
			Router: `/DelWeed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"],
		beego.ControllerComments{
			Method: "Mysql2Redis",
			Router: `/Mysql2Redis`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"],
		beego.ControllerComments{
			Method: "AddData",
			Router: `/add/data`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/essearch`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"],
		beego.ControllerComments{
			Method: "ReadJson",
			Router: `/read/json`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"],
		beego.ControllerComments{
			Method: "SaveJson",
			Router: `/save/json`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:EsController"],
		beego.ControllerComments{
			Method: "UploadFile",
			Router: `/upload`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:FileUploadController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:FileUploadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ProducterController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ProducterController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ProducterController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:ProducterController"],
		beego.ControllerComments{
			Method: "DbTrans",
			Router: `/dbtransform`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:TestController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:TestController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"] = append(beego.GlobalControllerRouter["zdy.local/cecetl/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
