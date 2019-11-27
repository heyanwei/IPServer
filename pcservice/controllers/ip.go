package controllers

import (
	"pcservice/handler"

	"github.com/astaxie/beego"
)

type IPController struct {
	beego.Controller
}

func (o *IPController) Get() {

	o.Data["json"] = handler.GetDevices()
	o.ServeJSON()
}

func (o *IPController) Post() {
	deviceName := o.GetString("device")

	res := handler.SaveDeviceMessage(deviceName, o.Ctx.Request.RemoteAddr)

	if res {
		o.Data["json"] = o.Ctx.Request.RemoteAddr
	} else {
		o.Data["json"] = "Format Error"
	}

	o.ServeJSON()
}
