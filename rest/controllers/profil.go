package controllers

import (
	"servicepspd/models"

	beego "github.com/beego/beego/v2/server/web"
)

type ProfilController struct {
	beego.Controller
}

func (u *ProfilController) Profil() {
	idpengguna := u.GetString("idpengguna")
	kategori := u.GetString("kategori")
	data := models.Getprofil(idpengguna, kategori)
	u.Data["json"] = data
	u.ServeJSON()
}
