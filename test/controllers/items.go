package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

 
type ItemsController struct {
    beego.Controller    
}
 
func (this *ItemsController) Get() {
    this.TplName = "itemsTemplate.html"
}
 
func (this *ItemsController) Post() {

    
    this.Data["id"] = this.GetString("id")
    this.Data["name"] = this.GetString("name")
    this.Data["quantity"] = this.GetString("quantity")

    quantity, _ := strconv.Atoi(this.Input().Get("quantity"))

    if quantity <= 100{
    	this.TplName = "availabilityTemplate.html"
    }else{
    	this.TplName = "itemsTemplate.html"//if quantity < 100 redirected to same page..
    }
}