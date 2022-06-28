package models

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/libra9z/orm"
)

func Getprofil(idpengguna, kategori string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var profil []orm.Params

	if kategori == "1" {
		_, err := o.Raw("SELECT NAMA, NIK, KLASIFIKASIPEGAWAI, STATUS FROM pppd.dbo.masterdosen WHERE PEGAWAIID = ?", idpengguna).Values(&profil)
		if err != nil {
			ResultData.Code = 1
			ResultData.Message = err.Error()
			return ResultData
		}
	} else if kategori == "2" {
		_, err := o.Raw("SELECT NO_MHS, NM_MHS, PHOTO FROM pppd.dbo.mastermahasiswa WHERE MahasiswaID = ?", idpengguna).Values(&profil)
		if err != nil {
			ResultData.Code = 1
			ResultData.Message = err.Error()
			return ResultData
		}
	}

	data := make(map[string]interface{})
	data["item"] = profil
	data["idpengguna"] = idpengguna
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData

}
