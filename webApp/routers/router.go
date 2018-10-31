package routers

import (
	"webApp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.MainController{}, "get:Showlogin;post:HandleLogin")
	beego.Router("/product/insert", &controllers.MainController{}, "get:ShowInsertPro;post:InserProduct")
	beego.Router("/product/productShow", &controllers.MainController{}, "get:ProductShow")        //方法第一个字母必须大写Pro not pro
	beego.Router("/product/productManage", &controllers.MainController{}, "get:ProductManage")    //方法第一个字母必须大写Pro not pro
	beego.Router("/product/detail", &controllers.MainController{}, "get:ProductDetail")           //方法第一个字母必须大写Pro not pro
	beego.Router("/product/detailquery", &controllers.MainController{}, "get:ProductDetailquery") //方法第一个字母必须大写Pro not pro
	beego.Router("/product/delete", &controllers.MainController{}, "get:ProductDel")
	beego.Router("/product/update", &controllers.MainController{}, "post:ProductUpdate")
}
