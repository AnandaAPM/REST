package controllers

import (
	"servicepspd/models"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

func (u *AuthController) GetData() {
	data := models.Getpengguna()
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *AuthController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	deviceid := u.GetString("deviceid")
	merkdevice := u.GetString("merkdevice")
	tipedevice := u.GetString("tipedevice")
	jenisdevice := u.GetString("jenisdevice")
	data := models.LoginUser(username, password, deviceid, merkdevice, tipedevice, jenisdevice)
	u.Data["json"] = data
	u.ServeJSON()
}
