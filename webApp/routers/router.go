package routers

import (
	"webApp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.MainController{}, "get:Showlogin;post:HandleLogin")
	beego.Router("/product", &controllers.MainController{}, "get:ShowInsertPro;post:InserProduct")
	beego.Router("/productShow", &controllers.MainController{}, "get:ProductShow")     //方法第一个字母必须大写Pro not pro
	beego.Router("/productManage", &controllers.MainController{}, "get:ProductManage") //方法第一个字母必须大写Pro not pro
}
