package models

import (
	"crypto/md5"
	"encoding/hex"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/libra9z/orm"
)

type Pengguna struct {
	Idpengguna string
}

type Mahasiswa struct {
	Idmhs  string
	No_mhs string
	Nm_mhs string
	Tlp_hp string
	Photo  string
}

func Getpengguna() Result {
	o := orm.NewOrm()
	var ResultData Result
	var makul []orm.Params
	_, err := o.Raw("SELECT KodeMataKuliah, Nama FROM pppd.dbo.MasterMatakuliah").Values(&makul)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}
	data := make(map[string]interface{})
	data["item"] = makul
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func LoginUser(username, password, deviceid, merkdevice, tipedevice, jenisdevice string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var pengguna []orm.Params

	h := md5.New()
	h.Write([]byte(password))
	hashpassword := hex.EncodeToString(h.Sum(nil))

	_, err := o.Raw(`SELECT
	data.PHOTO AS fotouser,
	data.NM_MHS AS namauser,
	data.NO_MHS AS nomoruser,
	CONVERT ( nvarchar ( 50 ), data.MahasiswaID ) AS userid,
	data.TLP_HP AS tlpuser,
	data.Jenis,
	RBAC.dbo.MSTPENGGUNA.USERNAME AS username,
	RBAC.dbo.MSTPENGGUNA.PASSWORD AS password,
	aktivasi_apps.deviceid,
	aktivasi_apps.tipe_device,
	aktivasi_apps.merk_device,
	aktivasi_apps.nama_apps,
	aktivasi_apps.no_telepon,
	aktivasi_apps.jenis_device,
	aktivasi_apps.tanggal_aktivasi,
	aktivasi_apps.stat_data 
FROM
	(
	SELECT
		PHOTO,
		NM_MHS,
		NO_MHS,
		MahasiswaID,
		TLP_HP,
		'Mahasiswa' AS Jenis 
	FROM
		pppd.dbo.mastermahasiswa UNION ALL
	SELECT
		kepegawaian.dbo.MSTPEGAWAI.FOTO,
		kepegawaian.dbo.View_Dosen_Lengkap.NAMALENGKAP_GELAR,
		kepegawaian.dbo.View_Dosen_Lengkap.NIK,
		kepegawaian.dbo.View_Dosen_Lengkap.PEGAWAIID,
		kepegawaian.dbo.MSTPEGAWAI.TELPONHP,
		'Dosen' AS Jenis 
	FROM
		kepegawaian.dbo.View_Dosen_Lengkap
		LEFT OUTER JOIN kepegawaian.dbo.MSTPEGAWAI ON kepegawaian.dbo.View_Dosen_Lengkap.PEGAWAIID = kepegawaian.dbo.MSTPEGAWAI.PEGAWAIID 
	) AS data
	INNER JOIN RBAC.dbo.MSTPENGGUNA ON data.MahasiswaID = RBAC.dbo.MSTPENGGUNA.PERSONID
	LEFT OUTER JOIN (
	SELECT
		id_aktivasi_apps,
		person_id,
		no_telepon,
		nama_apps,
		deviceid,
		tipe_device,
		merk_device,
		jenis_device,
		tanggal_aktivasi,
		stat_data 
	FROM
		RBAC.dbo.aktivasi_apps AS aktivasi_apps_1 
	WHERE
		( stat_data = 'Aktif' ) 
	) AS aktivasi_apps ON RBAC.dbo.MSTPENGGUNA.PERSONID = aktivasi_apps.person_id 
WHERE
	RBAC.dbo.MSTPENGGUNA.USERNAME = ? 
	AND ( RBAC.dbo.MSTPENGGUNA.PASSWORD = ? OR RBAC.dbo.MSTPENGGUNA.PASSWORD = ? )`, username, password, hashpassword).Values(&pengguna)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	if pengguna[0]["deviceid"] != deviceid && pengguna[0]["deviceid"] != nil {
		ResultData.Code = 1
		ResultData.Message = "Gadget yang digunakan tidak sesuai"
		return ResultData
	} else if pengguna[0]["deviceid"] != deviceid && pengguna[0]["deviceid"] == nil {
		_, erraktivasi := o.Raw(`INSERT INTO RBAC.dbo.aktivasi_apps ( id_aktivasi_apps,person_id, nama_apps, deviceid, tipe_device, merk_device, jenis_device, tanggal_aktivasi, stat_data )
			VALUES
				( NEWID(), ?, 'Moloco Mobile App',?,?,?,?, GETDATE( ), 'Aktif' )`, pengguna[0]["userid"], deviceid, tipedevice, merkdevice, jenisdevice).Exec()
		if erraktivasi != nil {
			ResultData.Code = 1
			ResultData.Message = erraktivasi.Error()
			return ResultData
		}
	}

	data := make(map[string]interface{})
	data["item"] = pengguna
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}
