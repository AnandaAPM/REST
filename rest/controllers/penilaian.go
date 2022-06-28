package controllers

import (
	"servicepspd/models"

	beego "github.com/beego/beego/v2/server/web"
)

type PenilaianController struct {
	beego.Controller
}

func (u *PenilaianController) Getkegiatanpenilaian() {
	idbagian := u.GetString("idbagian")
	data := models.Getkegiatan(idbagian)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getdosenpenilaian() {
	idrs := u.GetString("idrs")
	idbagian := u.GetString("idbagian")
	data := models.Getdokterklinik(idbagian, idrs)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getperiodepenilaian() {
	idmahasiswa := u.GetString("idmahasiswa")
	data := models.Getperiode(idmahasiswa)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getbagianpenilaian() {
	idmahasiswa := u.GetString("idmahasiswa")
	idperiode := u.GetString("idperiode")
	data := models.Getbagian(idmahasiswa, idperiode)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getkompetensipenilaian() {
	idbagian := u.GetString("idbagian")
	data := models.Getkompetensi(idbagian)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getkomponenpenilaian() {
	idbagian := u.GetString("idbagian")
	idkegiatan := u.GetString("idkegiatan")
	idkompetensi := u.GetString("idkompetensi")
	data := models.Getkomponen(idbagian, idkegiatan, idkompetensi)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getformpenilaian() {
	idbagian := u.GetString("idbagian")
	idkegiatan := u.GetString("idkegiatan")
	data := models.Getformpenilaian(idbagian, idkegiatan)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Createpenilaianmahasiswa() {
	idmahasiswa := u.GetString("idmahasiswa")
	idrs := u.GetString("idrs")
	idbagian := u.GetString("idbagian")
	idkegiatan := u.GetString("idkegiatan")
	idperiode := u.GetString("idperiode")
	iddosenklinik := u.GetString("iddosenklinik")
	tglkegiatan := u.GetString("tglkegiatan")
	idkompetensi := u.GetString("idkompetensi")
	dataform := u.GetString("dataform")
	data := models.Createpenilaianmhs(idmahasiswa, idrs, idbagian, idkegiatan, idperiode, tglkegiatan, iddosenklinik, idkompetensi, dataform)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Updatepenilaianmhs() {
	idnilai := u.GetString("idnilai")
	dataform := u.GetString("dataform")
	tglkegiatan := u.GetString("tglkegiatan")
	data := models.Updatepenilaian(idnilai, dataform, tglkegiatan)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Deletepenilaianmhs() {
	idnilai := u.GetString("idnilai")
	data := models.Deletepenilaian(idnilai)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getpenilaianmahasiswa() {
	idmahasiswa := u.GetString("idmahasiswa")
	offset := u.GetString("offset")
	idkegiatan := u.GetString("idkegiatan")
	idbagian := u.GetString("idbagian")
	statusapproval := u.GetString("statusapproval")
	data := models.Getpenilaianmhs(idmahasiswa, offset, idkegiatan, idbagian, statusapproval)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getpenilaiandosen() {
	iddosen := u.GetString("iddosen")
	offset := u.GetString("offset")
	search := u.GetString("search")
	idkegiatan := u.GetString("idkegiatan")
	statusapproval := u.GetString("statusapproval")
	data := models.Getpenilaiandsn(iddosen, offset, search, idkegiatan, statusapproval)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Createpenilaiandosen() {
	idnilai := u.GetString("idnilai")
	idbagian := u.GetString("idbagian")
	idkompetensi := u.GetString("idkompetensi")
	idmahasiswa := u.GetString("idmahasiswa")
	idkegiatan := u.GetString("idkegiatan")
	iddosenklinik := u.GetString("iddosenklinik")
	nilaikomponen := u.GetString("nilaikomponen")
	namadosen := u.GetString("namadosen")
	nilaijadi := u.GetString("nilaijadi")
	aspeksudahbagus := u.GetString("aspeksudahbagus")
	aspekdiperbaiki := u.GetString("aspekdiperbaiki")
	actionplan := u.GetString("actionplan")
	data := models.Createpenilaiandsn(idmahasiswa, idkegiatan, idnilai, idbagian, iddosenklinik, idkompetensi, namadosen, nilaikomponen, nilaijadi, aspeksudahbagus, aspekdiperbaiki, actionplan)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Updatepenilaiandosen() {
	idnilai := u.GetString("idnilai")
	idbagian := u.GetString("idbagian")
	idkompetensi := u.GetString("idkompetensi")
	idmahasiswa := u.GetString("idmahasiswa")
	idkegiatan := u.GetString("idkegiatan")
	iddosenklinik := u.GetString("iddosenklinik")
	nilaikomponen := u.GetString("nilaikomponen")
	namadosen := u.GetString("namadosen")
	nilaijadi := u.GetString("nilaijadi")
	aspeksudahbagus := u.GetString("aspeksudahbagus")
	aspekdiperbaiki := u.GetString("aspekdiperbaiki")
	actionplan := u.GetString("actionplan")
	data := models.Updatepenilaiandsn(idmahasiswa, idkegiatan, idnilai, idbagian, iddosenklinik, idkompetensi, namadosen, nilaikomponen, nilaijadi, aspeksudahbagus, aspekdiperbaiki, actionplan)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Showmore() {
	idnilai := u.GetString("idnilai")
	data := models.Getshowmore(idnilai)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Getrekapnilai() {
	idmahasiswa := u.GetString("idmahasiswa")
	idbagian := u.GetString("idbagian")
	data := models.Rekappenilaian(idmahasiswa, idbagian)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Validasipenilaian() {
	iddosen := u.GetString("iddosen")
	idnilai := u.GetString("idnilai")
	data := models.Validasipenilaian(iddosen, idnilai)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Countpenilaianmhs() {
	idmahasiswa := u.GetString("idmahasiswa")
	data := models.Countpenilaianmhs(idmahasiswa)
	u.Data["json"] = data
	u.ServeJSON()
}

func (u *PenilaianController) Countpenilaiandsn() {
	iddosen := u.GetString("iddosen")
	data := models.Countpenilaiandsn(iddosen)
	u.Data["json"] = data
	u.ServeJSON()
}
