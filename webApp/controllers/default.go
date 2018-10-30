package controllers

import (
	"strconv"
	"webApp/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "login.tpl"
}
func (c *MainController) Post() {
	c.Ctx.WriteString("ok")
}
func (c *MainController) Showlogin() {
	c.TplName = "login.tpl"
}
func (c *MainController) ProductShow() {
	c.TplName = "productManage.tpl"
}
func (c *MainController) HandleLogin() {
	o := orm.NewOrm()
	user := models.User{}
	user.Name = c.GetString("user")
	user.Pwd = c.GetString("password")
	err := o.Read(&user, "Name")
	//beego.Info(user, user.Name, user.Pwd)
	if err != nil || user.Pwd != c.GetString("password") {
		beego.Info("query fail")
		c.TplName = "index.tpl"
		return
	}
	c.Data["User"] = user.Name

	c.TplName = "index.tpl"
	//c.Redirect("/index.tpl", 302)
	//c.Ctx.WriteString("登陆成功")
}

func (c *MainController) ShowInsertPro() { //前端产品录入界面
	c.TplName = "product.tpl"
}
func (c *MainController) InserProduct() { //产品录入提交后进入此action，入数据库
	o := orm.NewOrm()
	pro := models.Product{}
	pro.Name = c.GetString("Name")
	Order_price, _ := strconv.ParseFloat(c.GetString("Order_price"), 64)
	Origin_price, _ := strconv.ParseFloat(c.GetString("Origin_price"), 64)
	Min_num, _ := strconv.Atoi(c.GetString("Min_num"))
	Discount, _ := strconv.ParseFloat(c.GetString("Discount"), 64)
	pro.Order_price = Order_price
	pro.Origin_price = Origin_price
	pro.Min_num = Min_num
	pro.Discount = Discount
	pro.Unit = c.GetString("Unit")
	_, err := o.Insert(&pro)
	if err != nil {
		beego.Info("insert product fail")
		return
	}
	c.TplName = "index.tpl"
	//c.Redirect("/index.tpl", 302)
	//c.Ctx.WriteString("登陆成功")
}

func (c *MainController) ProductManage() {

	o := orm.NewOrm()
	var products []*models.Product
	o.QueryTable("product").All(&products)

	//c.Ctx.WriteString("展示成果")
	//c.TplName = "index.tpl"
	//c.Redirect("/index.tpl", 302)
	c.Data["json"] = products
	c.ServeJSON()
}

//	beego.Info("ddsds")

//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//	orm.Debug = true
//	c := models.Customer{}

//	cp := models.Cus_pro{}
//	p := models.Product{}

//	//p.Name = "ferment"
//	p.Order_price = 198
//	p.Discount = 0.8
//	p.Min_num = 100
//	p.Id = 1
//	c.Id = 1
//	c.Name = "liyuling"
//	c.Phone = 13818571412
//	c.CreateTime = time.Now()
//	c.Birthday = "20181030"
//	cp.Customer_num = c.Phone
//	cp.Product_name = p.Name
//o := orm.NewOrm()
//	var cus_pro []*models.Cus_pro
//	_, err := o.QueryTable("cus_pro").Filter("Customer_num", c.Phone).All(&cus_pro)

//	if err == nil {
//		for _, order := range cus_pro {

//			beego.Info(order)
//		}
//	}
//_, err := o.Insert(&cp)
//o.Read(&cp, "Customer_num")
//	if err != nil {
//		beego.Info(err)
//		return
//	}
//beego.Info(cp)
//beego.Run()

//	if _, err := o.Update(&obj); err == nil {
//		beego.Info(obj)
//	}
