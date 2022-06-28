// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"servicepspd/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSRouter("/get", &controllers.AuthController{}, "get:GetData"),
			beego.NSRouter("/loginuser", &controllers.AuthController{}, "post:Login"),
		),
		beego.NSNamespace("/profil",
			beego.NSRouter("/getprofil", &controllers.ProfilController{}, "post:Profil"),
		),
		beego.NSNamespace("/penilaian",
			beego.NSRouter("/getkegiatan", &controllers.PenilaianController{}, "post:Getkegiatanpenilaian"),
			beego.NSRouter("/getdosenklinik", &controllers.PenilaianController{}, "post:Getdosenpenilaian"),
			beego.NSRouter("/getperiode", &controllers.PenilaianController{}, "post:Getperiodepenilaian"),
			beego.NSRouter("/getbagian", &controllers.PenilaianController{}, "post:Getbagianpenilaian"),
			beego.NSRouter("/getkompetensi", &controllers.PenilaianController{}, "post:Getkompetensipenilaian"),
			beego.NSRouter("/getkomponen", &controllers.PenilaianController{}, "post:Getkomponenpenilaian"),
			beego.NSRouter("/createpenilaianmhs", &controllers.PenilaianController{}, "post:Createpenilaianmahasiswa"),
			beego.NSRouter("/updatepenilaianmhs", &controllers.PenilaianController{}, "post:Updatepenilaianmhs"),
			beego.NSRouter("/deletepenilaianmhs", &controllers.PenilaianController{}, "post:Deletepenilaianmhs"),
			beego.NSRouter("/updatepenilaiandsn", &controllers.PenilaianController{}, "post:Updatepenilaiandosen"),
			beego.NSRouter("/createpenilaiandsn", &controllers.PenilaianController{}, "post:Createpenilaiandosen"),
			beego.NSRouter("/getpenilaianmahasiswa", &controllers.PenilaianController{}, "post:Getpenilaianmahasiswa"),
			beego.NSRouter("/getpenilaiandosen", &controllers.PenilaianController{}, "post:Getpenilaiandosen"),
			beego.NSRouter("/getshowmore", &controllers.PenilaianController{}, "post:Showmore"),
			beego.NSRouter("/getrekapnilai", &controllers.PenilaianController{}, "post:Getrekapnilai"),
			beego.NSRouter("/getformpenilaian", &controllers.PenilaianController{}, "post:Getformpenilaian"),
			beego.NSRouter("/validasipenilaian", &controllers.PenilaianController{}, "post:Validasipenilaian"),
			beego.NSRouter("/countpenilaianmahasiswa", &controllers.PenilaianController{}, "post:Countpenilaianmhs"),
			beego.NSRouter("/countpenilaiandosen", &controllers.PenilaianController{}, "post:Countpenilaiandsn"),
		),
		beego.NSNamespace("/logbook",
			beego.NSRouter("/getbagianlog", &controllers.LogbookController{}, "post:GetBagianLog"),
			beego.NSRouter("/getkompetensilog", &controllers.LogbookController{}, "post:GetKompetensiLog"),
			beego.NSRouter("/getdosenkliniklog", &controllers.LogbookController{}, "post:GetDosenKlinikLog"),
			beego.NSRouter("/getlogmhs", &controllers.LogbookController{}, "post:GetLogMhs"),
			beego.NSRouter("/getlogdsn", &controllers.LogbookController{}, "post:GetLogDsn"),
			beego.NSRouter("/createlogmhs", &controllers.LogbookController{}, "post:CreateLogMhs"),
			beego.NSRouter("/createlogdsn", &controllers.LogbookController{}, "post:CreateLogDsn"),
			beego.NSRouter("/rekaplogmhs", &controllers.LogbookController{}, "post:RekapLogMhs"),
			beego.NSRouter("/gettindakankemampuan", &controllers.LogbookController{}, "post:GetTindakanKemampuan"),
			beego.NSRouter("/deletelogbook", &controllers.LogbookController{}, "post:Deletelogbook"),
			beego.NSRouter("/updatelogbook", &controllers.LogbookController{}, "post:Updatelogbook"),
			beego.NSRouter("/countlogbookmahasiswa", &controllers.LogbookController{}, "post:Countlogbookmhs"),
			beego.NSRouter("/countlogbookdosen", &controllers.LogbookController{}, "post:Countlogbookdsn"),
		),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
