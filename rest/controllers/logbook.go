package controllers

import (
	"servicepspd/models"

	beego "github.com/beego/beego/v2/server/web"
)

type LogbookController struct {
	beego.Controller
}

func (u *LogbookController) GetBagianLog() {
	data := models.Getbagianlog()
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) GetKompetensiLog() {
	idbagian := u.GetString("idbagian")
	data := models.Getkompetensilog(idbagian)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) GetDosenKlinikLog() {
	idbagian := u.GetString("idbagian")
	data := models.Getdosenkliniklog(idbagian)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) GetLogMhs() {
	idmahasiswa := u.GetString("idmahasiswa")
	offset := u.GetString("offset")
	idbagian := u.GetString("idbagian")
	idkompetensi := u.GetString("idkompetensi")
	statusapproval := u.GetString("statusapproval")
	data := models.Getlogmhs(idmahasiswa, idbagian, idkompetensi, offset, statusapproval)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) GetLogDsn() {
	iddosen := u.GetString("iddosen")
	offset := u.GetString("offset")
	search := u.GetString("search")
	idkompetensi := u.GetString("idkompetensi")
	statusapproval := u.GetString("statusapproval")
	data := models.Getlogdsn(iddosen, search, idkompetensi, offset, statusapproval)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) CreateLogMhs() {
	idmahasiswa := u.GetString("idmahasiswa")
	idbagian := u.GetString("idbagian")
	idkompetensi := u.GetString("idkompetensi")
	iddosen := u.GetString("iddosen")
	data := models.Createlogmhs(idmahasiswa, idbagian, idkompetensi, iddosen)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) CreateLogDsn() {
	idmahasiswa := u.GetString("idmahasiswa")
	iddosen := u.GetString("iddosen")
	idtranslogbook := u.GetString("idtranslogbook")
	namadosen := u.GetString("namadosen")
	tindakan := u.GetString("tindakan")
	idtindakan := u.GetString("idtindakan")
	data := models.Createlogdsn(idtranslogbook, iddosen, idmahasiswa, namadosen, tindakan, idtindakan)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) RekapLogMhs() {
	idmahasiswa := u.GetString("idmahasiswa")
	idbagian := u.GetString("idbagian")
	data := models.Rekaplogmhs(idmahasiswa, idbagian)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) GetTindakanKemampuan() {
	idbagian := u.GetString("idbagian")
	data := models.Gettindakankemampuan(idbagian)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) Deletelogbook() {
	idtranslogbook := u.GetString("idtranslogbook")
	data := models.Deletelogbook(idtranslogbook)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) Updatelogbook() {
	idtranslogbook := u.GetString("idtranslogbook")
	idbagian := u.GetString("idbagian")
	idkompetensi := u.GetString("idkompetensi")
	iddosen := u.GetString("iddosen")
	data := models.Updatelogbook(idtranslogbook, idbagian, iddosen, idkompetensi)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) Countlogbookmhs() {
	idmahasiswa := u.GetString("idmahasiswa")
	data := models.Countlogbookmhs(idmahasiswa)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *LogbookController) Countlogbookdsn() {
	iddosen := u.GetString("iddosen")
	data := models.Countlogbookdsn(iddosen)
	u.Data["json"] = data
	u.ServeJSON()
}
