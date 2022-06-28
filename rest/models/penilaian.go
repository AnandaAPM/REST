package models

import (
	"encoding/json"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/libra9z/orm"
)

type Nilaicoas_pengajuan struct {
	IdNilai             string
	MahasiswaID_Pemohon string
	IdRs                string
	Idbagian            string
	Idkegiatan          string
	Idperiode           string
	Tglkegiatan         string
	Nilai               string
	Usr                 string
	Tglapprove          string
}

type Nilaikomponen struct {
	IdKomp   string
	Nilai    string
	Feedback string
}
type Nilaikomponenupdate struct {
	IdKomp   string
	Nilai    string
	Feedback string
}

type Forminput []struct {
	Idinput    string `json:"idinput"`
	Labelinput string `json:"labelinput"`
	Value      string `json:"value"`
}

type Formupdate []struct {
	Iddetail   string `json:"iddetail"`
	Labelinput string `json:"labelinput"`
	Value      string `json:"value"`
}

func Getkegiatan(idbagian string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var kegiatan []orm.Params

	_, err := o.Raw(`SELECT CONVERT
			( nvarchar ( 50 ), pppd.dbo.masterkegiatan.IdKegiatan) AS IdKegiatan,
			pppd.dbo.masterkegiatan.NamaKegiatan,
			pppd.dbo.masterkegiatan.Keterangan,
			pppd.dbo.masterkegiatan.Katagori 
		FROM
			pppd.dbo.settingformpeniaian
			LEFT JOIN pppd.dbo.masterkegiatan ON pppd.dbo.settingformpeniaian.idkegiatan = pppd.dbo.masterkegiatan.IdKegiatan 
		WHERE
			pppd.dbo.settingformpeniaian.Idbagian = ?
			AND pppd.dbo.settingformpeniaian.statusdata = ? 
		GROUP BY
			pppd.dbo.masterkegiatan.IdKegiatan,
			pppd.dbo.masterkegiatan.NamaKegiatan,
			pppd.dbo.masterkegiatan.Keterangan,
			pppd.dbo.masterkegiatan.Katagori`, idbagian, "Aktif").Values(&kegiatan)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = kegiatan
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData

}

func Getdokterklinik(idbagian, idrs string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var dokter []orm.Params

	_, err := o.Raw(`SELECT CONVERT(NVARCHAR(50),[PEGAWAIID]) as PEGAWAIID,[Nama],[NamaLengkap],[PANGKAT],[IDBAGIAN],[Bagian],[Status],[IDRS],[NamaRS],[Jenis]
FROM [pppd].[dbo].[DosenKlinik]
WHERE CONVERT(NVARCHAR(50),Jenis)='dosen' AND IDBAGIAN = ? AND (IDRS = ? OR IDRS = '01B1CE17-30F1-436E-8EED-37E0CD0B2A19')`, idbagian, idrs).Values(&dokter)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = dokter
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Getperiode(idmahasiswa string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var periode []orm.Params

	// SELECT convert(nvarchar(50), pppd.dbo.mastermahasiswa.MahasiswaID) as mhsid, convert(nvarchar(50), pppd.dbo.periodecoas.IdPeriode) as idperiode, pppd.dbo.mastermahasiswa.NO_MHS, pppd.dbo.mastermahasiswa.NM_MHS, pppd.dbo.periodecoas.NamaPeriode, pppd.dbo.periodecoas.AwalPeriode, pppd.dbo.periodecoas.AkhirPeriode, pppd.dbo.masterbagianrs.NamaBagian, pppd.dbo.masterrumahsakit.NamaRS
	// 	FROM  pppd.dbo.mastermahasiswa INNER JOIN
	// pppd.dbo.penempatan ON pppd.dbo.mastermahasiswa.MahasiswaID = pppd.dbo.penempatan.MahasiswaID INNER JOIN
	// pppd.dbo.masterbagianrs ON pppd.dbo.penempatan.IdBag = pppd.dbo.masterbagianrs.IdBag INNER JOIN
	// pppd.dbo.masterrumahsakit ON pppd.dbo.penempatan.IdRs = pppd.dbo.masterrumahsakit.IdRs INNER JOIN
	// pppd.dbo.periodecoas ON pppd.dbo.penempatan.IdPeriode = pppd.dbo.periodecoas.IdPeriode
	// 	WHERE (pppd.dbo.mastermahasiswa.MahasiswaID = ? AND pppd.dbo.periodecoas.AwalPeriode <= GETDATE() AND pppd.dbo.periodecoas.AkhirPeriode >= GETDATE())
	_, err := o.Raw(`SELECT convert(nvarchar(50), pppd.dbo.mastermahasiswa.MahasiswaID) as mhsid, convert(nvarchar(50), pppd.dbo.periodecoas.IdPeriode) as idperiode, pppd.dbo.mastermahasiswa.NO_MHS, pppd.dbo.mastermahasiswa.NM_MHS, pppd.dbo.periodecoas.NamaPeriode, pppd.dbo.periodecoas.AwalPeriode, pppd.dbo.periodecoas.AkhirPeriode
		FROM  pppd.dbo.mastermahasiswa INNER JOIN
	pppd.dbo.penempatan ON pppd.dbo.mastermahasiswa.MahasiswaID = pppd.dbo.penempatan.MahasiswaID INNER JOIN
	pppd.dbo.periodecoas ON pppd.dbo.penempatan.IdPeriode = pppd.dbo.periodecoas.IdPeriode
		WHERE (pppd.dbo.mastermahasiswa.MahasiswaID = ?) ORDER BY pppd.dbo.periodecoas.AwalPeriode DESC`, idmahasiswa).Values(&periode)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = periode
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Getbagian(idmahasiswa, idperiode string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var bagian []orm.Params

	_, err := o.Raw(`SELECT convert(nvarchar(50), pppd.dbo.mastermahasiswa.MahasiswaID) as mhsid, convert(nvarchar(50), pppd.dbo.periodecoas.IdPeriode) as idperiode, pppd.dbo.mastermahasiswa.NO_MHS, pppd.dbo.mastermahasiswa.NM_MHS, pppd.dbo.periodecoas.NamaPeriode, pppd.dbo.periodecoas.AwalPeriode, pppd.dbo.periodecoas.AkhirPeriode, pppd.dbo.masterbagianrs.NamaBagian, pppd.dbo.masterrumahsakit.NamaRS, convert(nvarchar(50), pppd.dbo.masterbagianrs.IdBag) as idbagian, convert(nvarchar(50), pppd.dbo.masterrumahsakit.IdRs) as idrs
		FROM  pppd.dbo.mastermahasiswa INNER JOIN
	pppd.dbo.penempatan ON pppd.dbo.penempatan.MahasiswaID = pppd.dbo.mastermahasiswa.MahasiswaID INNER JOIN
	pppd.dbo.masterbagianrs ON pppd.dbo.penempatan.IdBag = pppd.dbo.masterbagianrs.IdBag INNER JOIN
	pppd.dbo.masterrumahsakit ON pppd.dbo.penempatan.IdRs = pppd.dbo.masterrumahsakit.IdRs INNER JOIN
	pppd.dbo.periodecoas ON pppd.dbo.penempatan.IdPeriode = pppd.dbo.periodecoas.IdPeriode 
		WHERE (pppd.dbo.mastermahasiswa.MahasiswaID = ?) ORDER BY pppd.dbo.periodecoas.AwalPeriode DESC`, idmahasiswa).Values(&bagian)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = bagian
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData

}

func Getkompetensi(idbagian string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var kompetensi []orm.Params

	_, err := o.Raw(`SELECT convert(nvarchar(50), pppd.dbo.MasterKompetensiPenilaian.IdKompetensi) as idkompetensi, pppd.dbo.MasterKompetensiPenilaian.Keterangan, pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi 
	FROM pppd.dbo.settingformpeniaian
	LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON pppd.dbo.settingformpeniaian.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi
	WHERE pppd.dbo.settingformpeniaian.Idbagian = ? AND pppd.dbo.settingformpeniaian.statusdata = ?
	GROUP BY 
	pppd.dbo.MasterKompetensiPenilaian.IdKompetensi,
	pppd.dbo.MasterKompetensiPenilaian.Keterangan,
	pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi`, idbagian, "Aktif").Values(&kompetensi)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = kompetensi
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Getkomponen(idbagian, idkegiatan, idkompetensi string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var komponen []orm.Params

	if idkompetensi != "" {
		_, err := o.Raw(`SELECT convert(nvarchar(50), pppd.dbo.masterbagianrs.IdBag) as idbagian, convert(nvarchar(50), pppd.dbo.masterkegiatan.IdKegiatan) as idkegiatan, convert(nvarchar(50), pppd.dbo.MasterKompetensiPenilaian.IdKompetensi) as idkompetensi, convert(nvarchar(50), pppd.dbo.MasterKomponenKompdinilai.idkomponenkompdinilai) as idkomponen, pppd.dbo.MasterKomponenKompdinilai.namakompdinilai, pppd.dbo.MasterKomponenKompdinilai.nilaimax, pppd.dbo.MasterKomponenKompdinilai.keterangan, pppd.dbo.settingformpeniaian.KodePenghitungan
	FROM pppd.dbo.settingformpeniaian LEFT JOIN
	pppd.dbo.masterbagianrs ON pppd.dbo.settingformpeniaian.Idbagian = pppd.dbo.masterbagianrs.IdBag LEFT JOIN
	pppd.dbo.masterkegiatan ON pppd.dbo.settingformpeniaian.idkegiatan = pppd.dbo.masterkegiatan.IdKegiatan LEFT JOIN
	pppd.dbo.MasterKompetensiPenilaian ON pppd.dbo.settingformpeniaian.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi LEFT JOIN
	pppd.dbo.MasterKomponenKompdinilai ON pppd.dbo.settingformpeniaian.idkomponenkompdinilai = pppd.dbo.MasterKomponenKompdinilai.idkomponenkompdinilai
	WHERE pppd.dbo.settingformpeniaian.statusdata = 'Aktif' AND pppd.dbo.settingformpeniaian.Idbagian = ? AND pppd.dbo.settingformpeniaian.idkegiatan = ? AND pppd.dbo.settingformpeniaian.IdKompetensi = ? AND pppd.dbo.settingformpeniaian.statusdata = ?`, idbagian, idkegiatan, idkompetensi, "Aktif").Values(&komponen)
		if err != nil {
			ResultData.Code = 1
			ResultData.Message = err.Error()
			return ResultData
		}
	} else if idkompetensi == "" {
		_, err := o.Raw(`SELECT convert(nvarchar(50), pppd.dbo.masterbagianrs.IdBag) as idbagian, convert(nvarchar(50), pppd.dbo.masterkegiatan.IdKegiatan) as idkegiatan, convert(nvarchar(50), pppd.dbo.MasterKompetensiPenilaian.IdKompetensi) as idkompetensi, convert(nvarchar(50), pppd.dbo.MasterKomponenKompdinilai.idkomponenkompdinilai) as idkomponen, pppd.dbo.MasterKomponenKompdinilai.namakompdinilai, pppd.dbo.MasterKomponenKompdinilai.nilaimax, pppd.dbo.MasterKomponenKompdinilai.keterangan, pppd.dbo.settingformpeniaian.KodePenghitungan
	FROM pppd.dbo.settingformpeniaian LEFT JOIN
	pppd.dbo.masterbagianrs ON pppd.dbo.settingformpeniaian.Idbagian = pppd.dbo.masterbagianrs.IdBag LEFT JOIN
	pppd.dbo.masterkegiatan ON pppd.dbo.settingformpeniaian.idkegiatan = pppd.dbo.masterkegiatan.IdKegiatan LEFT JOIN
	pppd.dbo.MasterKompetensiPenilaian ON pppd.dbo.settingformpeniaian.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi LEFT JOIN
	pppd.dbo.MasterKomponenKompdinilai ON pppd.dbo.settingformpeniaian.idkomponenkompdinilai = pppd.dbo.MasterKomponenKompdinilai.idkomponenkompdinilai
	WHERE pppd.dbo.settingformpeniaian.statusdata = 'Aktif' AND pppd.dbo.settingformpeniaian.Idbagian = ? AND pppd.dbo.settingformpeniaian.idkegiatan = ? AND pppd.dbo.settingformpeniaian.statusdata = ?`, idbagian, idkegiatan, "Aktif").Values(&komponen)
		if err != nil {
			ResultData.Code = 1
			ResultData.Message = err.Error()
			return ResultData
		}
	}

	data := make(map[string]interface{})
	data["item"] = komponen
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Getformpenilaian(idbagian, idkegiatan string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var formnilai []orm.Params

	_, errget := o.Raw(`SELECT
	convert(nvarchar(50), pppd.dbo.SettingFormInput.IdBag) AS idbag,
	convert(nvarchar(50), pppd.dbo.SettingFormInput.IdKegiatan) AS idkegiatan,
	pppd.dbo.JenisInput.Labelnput,
	pppd.dbo.JenisInput.NameInput,
	pppd.dbo.JenisInput.InputType,
	pppd.dbo.JenisInput.IdJenisInput,
	parent = STUFF(
		(
		SELECT
			',' + '"' + md.value + '"' 
		FROM
			pppd.dbo.parentJenisInput md 
		WHERE
			md.IdJenisInput = pppd.dbo.JenisInput.IdJenisInput FOR XML PATH ( '' ),
			TYPE 
		).value ( '.', 'NVARCHAR(MAX)' ),
		1,
		1,
		'' 
	) 
FROM
	pppd.dbo.JenisInput
	INNER JOIN pppd.dbo.SettingFormInput ON pppd.dbo.JenisInput.IdJenisInput = pppd.dbo.SettingFormInput.IdJenisInput 
WHERE
	( pppd.dbo.SettingFormInput.IdBag = ? ) 
	AND ( pppd.dbo.SettingFormInput.IdKegiatan = ? ) 
ORDER BY
	NomorUrut`, idbagian, idkegiatan).Values(&formnilai)
	if errget != nil {
		ResultData.Code = 1
		ResultData.Message = errget.Error()
		return ResultData
	}
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = formnilai
	return ResultData
}

func Createpenilaianmhs(idmahasiswa, idrs, idbagian, idkegiatan, idperiode, tglkegiatan, iddosenklinik, idkompetensi, dataform string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var idpengajuan []orm.Params
	var datapengajuan []orm.Params
	statusdata := "Aktif"
	statusapproval := "Diajukan"
	dt := time.Now()

	_, errget := o.Raw(`SELECT * FROM pppd.dbo.nilaicoas_pengajuan 
	WHERE pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = ? 
	AND pppd.dbo.nilaicoas_pengajuan.IdPeriode = ? 
	AND pppd.dbo.nilaicoas_pengajuan.IdBag = ? 
	AND pppd.dbo.nilaicoas_pengajuan.IdKegiatan = ? 
	AND pppd.dbo.nilaicoas_pengajuan.Tgl_kegiatan = ? 
	AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif'`, idmahasiswa, idperiode, idbagian, idkegiatan, tglkegiatan).Values(&datapengajuan)
	if errget != nil {
		ResultData.Code = 1
		ResultData.Message = errget.Error()
		return ResultData
	}

	if datapengajuan != nil {
		ResultData.Code = 1
		ResultData.Message = "Anda Sudah Pernah Mengajukan di periode yang sama, bagian yang sama dan tanggal yang sama "
		return ResultData
	}

	if idkompetensi != "" {
		_, err := o.Raw(`INSERT INTO pppd.dbo.nilaicoas_pengajuan ( 
			IdNilai, 
			MahasiswaID_Pemohon, 
			IdRs, 
			IdBag, 
			IdKegiatan, 
			IdPeriode, 
			Tgl_kegiatan, 
			Tglmsk, 
			Status_Data, 
			Pegawaiid_aproval, 
			IdKompetensi, 
			Status_Approval ) OUTPUT CONVERT ( nvarchar ( 50 ), Inserted.IdNilai ) AS idnilai
			VALUES
				( NEWID( ),?,?,?,?,?,?,?,?,?,?,? )`,
			idmahasiswa, idrs, idbagian, idkegiatan, idperiode, tglkegiatan, dt.Format("01-02-2006 15:04:05"), statusdata, iddosenklinik, idkompetensi, statusapproval).Values(&idpengajuan)
		if err != nil {
			ResultData.Code = 1
			ResultData.Message = err.Error()
			return ResultData
		}
	} else if idkompetensi == "" {
		_, err := o.Raw(`INSERT INTO pppd.dbo.nilaicoas_pengajuan ( 
			IdNilai, 
			MahasiswaID_Pemohon, 
			IdRs, 
			IdBag, 
			IdKegiatan, 
			IdPeriode, 
			Tgl_kegiatan, 
			Tglmsk, 
			Status_Data, 
			Pegawaiid_aproval, 
			Status_Approval ) OUTPUT CONVERT ( nvarchar ( 50 ), Inserted.IdNilai ) AS idnilai
			VALUES
				( NEWID( ),?,?,?,?,?,?,?,?,?,? )`,
			idmahasiswa, idrs, idbagian, idkegiatan, idperiode, tglkegiatan, dt.Format("01-02-2006 15:04:05"), statusdata, iddosenklinik, statusapproval).Values(&idpengajuan)
		if err != nil {
			ResultData.Code = 1
			ResultData.Message = err.Error()
			return ResultData
		}

	}

	var form Forminput
	byt := []byte(dataform)
	if err := json.Unmarshal(byt, &form); err != nil {
		panic(err)
	}
	for i := range form {
		_, errval := o.Raw(`INSERT INTO pppd.dbo.DetailPengajuan ( 
			IdDetailPengajuan, 
			IdNilai, 
			IdJenisInput, 
			ValueInput,
			Created_at )
			VALUES
				( NEWID( ), ?, ?, ?, ? )`, idpengajuan[0]["idnilai"], form[i].Idinput, form[i].Value, dt.Format("01-02-2006 15:04:05")).Exec()
		if errval != nil {
			ResultData.Code = 1
			ResultData.Message = errval.Error()
			return ResultData
		}
	}

	data := make(map[string]interface{})
	data["item"] = idpengajuan
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Updatepenilaian(idnilai, dataform, tglkegiatan string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params
	dt := time.Now()

	_, errget := o.Raw(`SELECT * FROM pppd.dbo.nilaicoas_pengajuan 
	WHERE pppd.dbo.nilaicoas_pengajuan.IdNilai = ? `, idnilai).Values(&datanilai)
	if errget != nil {
		ResultData.Code = 1
		ResultData.Message = errget.Error()
		return ResultData
	}

	if datanilai[0]["Status_Approval"] != "Diajukan" {
		ResultData.Code = 1
		ResultData.Message = "Nilai Sudah Diproses, tidak dapat diubah"
		return ResultData
	}

	if tglkegiatan != "" {
		_, errval := o.Raw(`UPDATE pppd.dbo.nilaicoas_pengajuan 
		SET Tgl_kegiatan = ? 
		WHERE
			pppd.dbo.nilaicoas_pengajuan.IdNilai  = ?`, tglkegiatan, idnilai).Exec()
		if errval != nil {
			ResultData.Code = 1
			ResultData.Message = errval.Error()
			return ResultData
		}
	}

	var form Formupdate
	byt := []byte(dataform)
	if err := json.Unmarshal(byt, &form); err != nil {
		panic(err)
	}
	for i := range form {
		_, errval := o.Raw(`UPDATE pppd.dbo.DetailPengajuan 
		SET ValueInput = ? , Updated_at = ?
		WHERE
			pppd.dbo.DetailPengajuan.IdDetailPengajuan = ?`, form[i].Value, dt.Format("01-02-2006 15:04:05"), form[i].Iddetail).Exec()
		if errval != nil {
			ResultData.Code = 1
			ResultData.Message = errval.Error()
			return ResultData
		}
	}

	data := make(map[string]interface{})
	// data["item"] = datanilai
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Deletepenilaian(idnilai string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datadelete []orm.Params
	var datanilai []orm.Params

	_, errget := o.Raw(`SELECT * FROM pppd.dbo.nilaicoas_pengajuan 
	WHERE pppd.dbo.nilaicoas_pengajuan.IdNilai = ? `, idnilai).Values(&datanilai)
	if errget != nil {
		ResultData.Code = 1
		ResultData.Message = errget.Error()
		return ResultData
	}

	if datanilai[0]["Status_Approval"] != "Diajukan" {
		ResultData.Code = 1
		ResultData.Message = "Nilai Sudah Diproses, tidak dapat diubah"
		return ResultData
	}

	_, err := o.Raw(`UPDATE pppd.dbo.nilaicoas_pengajuan SET Status_Data = 'Tidak Aktif' WHERE pppd.dbo.nilaicoas_pengajuan.IdNilai = ? `, idnilai).Values(&datadelete)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	// data := make(map[string]interface{})
	// data["item"] = idpengajuan
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = datadelete
	return ResultData
}

func Getpenilaianmhs(idmahasiswa, offset, idkegiatan, idbagian, statusapproval string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params
	var datavalue []orm.Params

	var stringwhere string
	if idkegiatan == "" && idbagian == "" {
		stringwhere = ``
	} else if idkegiatan != "" && idbagian == "" {
		stringwhere = `AND pppd.dbo.masterkegiatan.IdKegiatan = '` + idkegiatan + `' `
	} else if idkegiatan == "" && idbagian != "" {
		stringwhere = `AND pppd.dbo.masterbagianrs.IdBag = '` + idbagian + `' `
	} else {
		stringwhere = `AND pppd.dbo.masterkegiatan.IdKegiatan = '` + idkegiatan + `' AND pppd.dbo.masterbagianrs.IdBag = '` + idbagian + `' `
	}

	var stringoffset string
	var stringlimit string
	if offset == "" {
		stringoffset = ` OFFSET COALESCE(0,0) ROWS `
		stringlimit = `FETCH FIRST COALESCE(10,0x7ffffff) ROWS ONLY `
	} else {
		stringoffset = ` OFFSET COALESCE(` + offset + `,0) ROWS `
		stringlimit = `FETCH NEXT COALESCE(10,0x7ffffff) ROWS ONLY `
	}

	var stringstatus string
	if statusapproval == "1" {
		stringstatus = ` AND pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diajukan' `
	} else if statusapproval == "2" {
		stringstatus = ` AND pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diinput Dosen' `
	} else if statusapproval == "3" {
		stringstatus = ` AND pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Validasi Admin' `
	} else {
		stringstatus = ``
	}

	_, err := o.Raw(`SELECT CONVERT
	( nvarchar ( 50 ), pppd.dbo.mastermahasiswa.MahasiswaID ) AS idmhs,
	pppd.dbo.mastermahasiswa.NO_MHS,
	pppd.dbo.mastermahasiswa.NM_MHS,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.periodecoas.IdPeriode ) AS idperiode,
	pppd.dbo.periodecoas.NamaPeriode,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterbagianrs.IdBag ) AS idbagian,
	pppd.dbo.masterbagianrs.NamaBagian,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterrumahsakit.IdRs ) AS idrs,
	pppd.dbo.masterrumahsakit.NamaRS,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterkegiatan.IdKegiatan ) AS idkegiatan,
	pppd.dbo.masterkegiatan.NamaKegiatan,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.nilaicoas_pengajuan.IdNilai ) AS idnilai,
	pppd.dbo.nilaicoas_pengajuan.Nilai,
	pppd.dbo.nilaicoas_pengajuan.USR,
	pppd.dbo.nilaicoas_pengajuan.Tglmsk,
	pppd.dbo.nilaicoas_pengajuan.Tglaproval,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval ) iddosen,
	pppd.dbo.nilaicoas_pengajuan.Status_Approval,
	pppd.dbo.nilaicoas_pengajuan.Tgl_kegiatan,
	kepegawaian.dbo.gelar ( pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval ) AS NAMALENGKAP_GELAR,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.MasterKompetensiPenilaian.IdKompetensi ) AS idkompetensi,
	pppd.dbo.MasterKompetensiPenilaian.Keterangan,
	pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi
FROM
	pppd.dbo.masterkegiatan
	INNER JOIN pppd.dbo.nilaicoas_pengajuan ON pppd.dbo.masterkegiatan.IdKegiatan = pppd.dbo.nilaicoas_pengajuan.IdKegiatan
	INNER JOIN pppd.dbo.mastermahasiswa
	INNER JOIN pppd.dbo.penempatan ON pppd.dbo.mastermahasiswa.MahasiswaID = pppd.dbo.penempatan.MahasiswaID
	INNER JOIN pppd.dbo.masterrumahsakit ON pppd.dbo.penempatan.IdRs = pppd.dbo.masterrumahsakit.IdRs
	INNER JOIN pppd.dbo.masterbagianrs ON pppd.dbo.penempatan.IdBag = pppd.dbo.masterbagianrs.IdBag
	INNER JOIN pppd.dbo.periodecoas ON pppd.dbo.penempatan.IdPeriode = pppd.dbo.periodecoas.IdPeriode ON pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = pppd.dbo.penempatan.MahasiswaID 
	AND pppd.dbo.nilaicoas_pengajuan.IdRs = pppd.dbo.penempatan.IdRs 
	AND pppd.dbo.nilaicoas_pengajuan.IdBag = pppd.dbo.penempatan.IdBag 
	AND pppd.dbo.nilaicoas_pengajuan.IdPeriode = pppd.dbo.penempatan.IdPeriode
	LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON (pppd.dbo.nilaicoas_pengajuan.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi OR (pppd.dbo.nilaicoas_pengajuan.IdKompetensi IS NULL AND pppd.dbo.MasterKompetensiPenilaian.IdKompetensi IS NULL))
WHERE
	( pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = ? AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' `+stringwhere+``+stringstatus+`)
	ORDER BY pppd.dbo.nilaicoas_pengajuan.Tglmsk DESC`+stringoffset+``+stringlimit+``, idmahasiswa).Values(&datanilai)

	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	for i := range datanilai {
		_, errval := o.Raw(`SELECT CONVERT
		( nvarchar ( 50 ), pppd.dbo.DetailPengajuan.IdDetailPengajuan ) AS idDetail,
		pppd.dbo.DetailPengajuan.ValueInput,
		pppd.dbo.JenisInput.Labelnput 
	FROM
		pppd.dbo.DetailPengajuan
		INNER JOIN pppd.dbo.JenisInput ON pppd.dbo.DetailPengajuan.IdJenisInput = pppd.dbo.JenisInput.IdJenisInput
		WHERE pppd.dbo.DetailPengajuan.IdNilai = ? `, datanilai[i]["idnilai"]).Values(&datavalue)
		if errval != nil {
			ResultData.Code = 1
			ResultData.Message = errval.Error()
			return ResultData
		}
		datanilai[i]["value"] = datavalue
	}

	data := make(map[string]interface{})
	data["item"] = datanilai
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData

}

func Getpenilaiandsn(iddosen, offset, search, idkegiatan, statusapproval string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params

	var stringwhere string
	if search == "" && idkegiatan == "" {
		stringwhere = ``
	} else if search != "" && idkegiatan == "" {
		stringwhere = ` AND pppd.dbo.mastermahasiswa.NM_MHS LIKE '%` + search + `%' `
	} else if search == "" && idkegiatan != "" {
		stringwhere = `AND pppd.dbo.masterkegiatan.IdKegiatan = '` + idkegiatan + `' `
	} else {
		stringwhere = `AND pppd.dbo.masterkegiatan.IdKegiatan = '` + idkegiatan + `' AND pppd.dbo.mastermahasiswa.NM_MHS LIKE '%` + search + `%' `
	}

	var stringoffset string
	var stringlimit string
	if offset == "" {
		stringoffset = ` OFFSET COALESCE(0,0) ROWS `
		stringlimit = `FETCH FIRST COALESCE(10,0x7ffffff) ROWS ONLY `
	} else {
		stringoffset = ` OFFSET COALESCE(` + offset + `,0) ROWS `
		stringlimit = `FETCH NEXT COALESCE(10,0x7ffffff) ROWS ONLY `
	}

	var stringstatus string
	if statusapproval == "1" {
		stringstatus = ` AND pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diajukan' `
	} else if statusapproval == "2" {
		stringstatus = ` AND pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diinput Dosen' `
	} else if statusapproval == "3" {
		stringstatus = ` AND pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Validasi Admin' `
	} else {
		stringstatus = ``
	}

	_, err := o.Raw(`SELECT CONVERT
	( nvarchar ( 50 ), pppd.dbo.mastermahasiswa.MahasiswaID ) AS idmhs,
	pppd.dbo.mastermahasiswa.NO_MHS,
	pppd.dbo.mastermahasiswa.NM_MHS,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.periodecoas.IdPeriode ) AS idperiode,
	pppd.dbo.periodecoas.NamaPeriode,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterbagianrs.IdBag ) AS idbagian,
	pppd.dbo.masterbagianrs.NamaBagian,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterrumahsakit.IdRs ) AS idrs,
	pppd.dbo.masterrumahsakit.NamaRS,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterkegiatan.IdKegiatan ) AS idkegiatan,
	pppd.dbo.masterkegiatan.NamaKegiatan,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.nilaicoas_pengajuan.IdNilai ) AS idnilai,
	pppd.dbo.nilaicoas_pengajuan.Nilai,
	pppd.dbo.nilaicoas_pengajuan.USR,
	pppd.dbo.nilaicoas_pengajuan.Tglmsk,
	pppd.dbo.nilaicoas_pengajuan.Tglaproval,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval ) iddosen,
	pppd.dbo.nilaicoas_pengajuan.Status_Approval,
	pppd.dbo.nilaicoas_pengajuan.Tgl_kegiatan,
	kepegawaian.dbo.gelar ( pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval ) AS NAMALENGKAP_GELAR,
	pppd.dbo.nilaicoas_pengajuan.NamaPasien,
	pppd.dbo.nilaicoas_pengajuan.JK_Pasien,
	pppd.dbo.nilaicoas_pengajuan.UsiaPasien,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.MasterKompetensiPenilaian.IdKompetensi ) AS idkompetensi,
	pppd.dbo.MasterKompetensiPenilaian.Keterangan,
	pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi 
	FROM
		pppd.dbo.masterkegiatan
		INNER JOIN pppd.dbo.nilaicoas_pengajuan ON pppd.dbo.masterkegiatan.IdKegiatan = pppd.dbo.nilaicoas_pengajuan.IdKegiatan
		INNER JOIN pppd.dbo.mastermahasiswa
		INNER JOIN pppd.dbo.penempatan ON pppd.dbo.mastermahasiswa.MahasiswaID = pppd.dbo.penempatan.MahasiswaID
		INNER JOIN pppd.dbo.masterrumahsakit ON pppd.dbo.penempatan.IdRs = pppd.dbo.masterrumahsakit.IdRs
		INNER JOIN pppd.dbo.masterbagianrs ON pppd.dbo.penempatan.IdBag = pppd.dbo.masterbagianrs.IdBag
		INNER JOIN pppd.dbo.periodecoas ON pppd.dbo.penempatan.IdPeriode = pppd.dbo.periodecoas.IdPeriode ON pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = pppd.dbo.penempatan.MahasiswaID 
		AND pppd.dbo.nilaicoas_pengajuan.IdRs = pppd.dbo.penempatan.IdRs 
		AND pppd.dbo.nilaicoas_pengajuan.IdBag = pppd.dbo.penempatan.IdBag 
		AND pppd.dbo.nilaicoas_pengajuan.IdPeriode = pppd.dbo.penempatan.IdPeriode
		LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON (
			pppd.dbo.nilaicoas_pengajuan.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi 
			OR ( pppd.dbo.nilaicoas_pengajuan.IdKompetensi IS NULL AND pppd.dbo.MasterKompetensiPenilaian.IdKompetensi IS NULL ) 
		) 
	WHERE
		( pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval = ? AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' `+stringwhere+``+stringstatus+` ) 
		ORDER BY pppd.dbo.nilaicoas_pengajuan.Tglmsk DESC
	`+stringoffset+``+stringlimit+`
	`, iddosen).Values(&datanilai)

	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = datanilai
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData

}

func Createpenilaiandsn(idmahasiswa, idkegiatan, idnilai, idbagian, iddosenklinik, idkompetensi, namadosen, nilaikomponen, nilaijadi, aspeksudahbagus, aspekdiperbaiki, actionplan string) Result {
	o := orm.NewOrm()
	var ResultData Result
	dt := time.Now()
	statusdata := "Diinput Dosen"
	i := 0
	var datnil []Nilaikomponen
	var stringkomp = nilaikomponen
	var err = json.Unmarshal([]byte(stringkomp), &datnil)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	if idkompetensi == "" {
		for {
			_, errin := o.Raw("INSERT INTO pppd.dbo.TransKompNilai(IdTransKompNilai, IdMahasiswa, IdKegiatan, Nilai,  IdBagian, Tgl_Input, Id_Pegawai, Usr, Status_Data, Idkomponenkompdinilai, IdNilai, feedback) VALUES (NEWID(),?,?,?,?,?,?,?,?,?,?,?)",
				idmahasiswa, idkegiatan, datnil[i].Nilai, idbagian, dt.Format("01-02-2006 15:04:05"), iddosenklinik, namadosen, statusdata, datnil[i].IdKomp, idnilai, datnil[i].Feedback).Exec()
			if errin != nil {
				ResultData.Code = 1
				ResultData.Message = errin.Error()
				return ResultData
			}
			i++
			if i == len(datnil) {
				break
			}
		}
	} else if idkompetensi != "" {
		for {
			_, errin := o.Raw("INSERT INTO pppd.dbo.TransKompNilai(IdTransKompNilai, IdMahasiswa, IdKegiatan, Nilai, IdKompetensi, IdBagian, Tgl_Input, Id_Pegawai, Usr, Status_Data, Idkomponenkompdinilai, IdNilai, feedback) VALUES (NEWID(),?,?,?,?,?,?,?,?,?,?,?,?)",
				idmahasiswa, idkegiatan, datnil[i].Nilai, idkompetensi, idbagian, dt.Format("01-02-2006 15:04:05"), iddosenklinik, namadosen, statusdata, datnil[i].IdKomp, idnilai, datnil[i].Feedback).Exec()
			if errin != nil {
				ResultData.Code = 1
				ResultData.Message = errin.Error()
				return ResultData
			}
			i++
			if i == len(datnil) {
				break
			}
		}
	}

	_, errupapp := o.Raw("UPDATE pppd.dbo.nilaicoas_pengajuan SET Nilai = ?, USR = ?, Tglaproval = ?, Status_Approval = ?, AspekSudahBagus = ?, AspekDiperbaiki = ?, ActionPlan = ? WHERE IdNilai = ?", nilaijadi, namadosen, dt.Format("01-02-2006 15:04:05"), statusdata, aspeksudahbagus, aspekdiperbaiki, actionplan, idnilai).Exec()
	if errupapp != nil {
		ResultData.Code = 1
		ResultData.Message = errupapp.Error()
		return ResultData
	}

	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = statusdata
	return ResultData
}

func Updatepenilaiandsn(idmahasiswa, idkegiatan, idnilai, idbagian, iddosenklinik, idkompetensi, namadosen, nilaikomponen, nilaijadi, aspeksudahbagus, aspekdiperbaiki, actionplan string) Result {
	o := orm.NewOrm()
	var ResultData Result
	i := 0
	dt := time.Now()
	statusdata := "Diinput Dosen"
	var datnil []Nilaikomponen
	var stringkomp = nilaikomponen
	var err = json.Unmarshal([]byte(stringkomp), &datnil)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	if idkompetensi == "" {
		for {
			_, errin := o.Raw(`UPDATE pppd.dbo.TransKompNilai 
			SET Nilai = ?,
			feedback = ? 
			WHERE
				IdMahasiswa = ? 
				AND IdKegiatan = ? 
				AND IdBagian = ? 
				AND Id_Pegawai = ? 
				AND Idkomponenkompdinilai = ? 
				AND IdNilai = ?`,
				datnil[i].Nilai, datnil[i].Feedback, idmahasiswa, idkegiatan, idbagian, iddosenklinik, datnil[i].IdKomp, idnilai).Exec()
			if errin != nil {
				ResultData.Code = 1
				ResultData.Message = errin.Error()
				return ResultData
			}
			i++
			if i == len(datnil) {
				break
			}
		}
	} else if idkompetensi != "" {
		for {
			_, errin := o.Raw(`UPDATE pppd.dbo.TransKompNilai 
			SET Nilai = ?,
			feedback = ? 
			WHERE
				IdMahasiswa = ? 
				AND IdKegiatan = ? 
				AND IdKompetensi = ?
				AND IdBagian = ? 
				AND Id_Pegawai = ? 
				AND Idkomponenkompdinilai = ? 
				AND IdNilai = ?`,
				datnil[i].Nilai, datnil[i].Feedback, idmahasiswa, idkegiatan, idkompetensi, idbagian, iddosenklinik, datnil[i].IdKomp, idnilai).Exec()
			if errin != nil {
				ResultData.Code = 1
				ResultData.Message = errin.Error()
				return ResultData
			}
			i++
			if i == len(datnil) {
				break
			}
		}
	}

	_, errupapp := o.Raw("UPDATE pppd.dbo.nilaicoas_pengajuan SET Nilai = ?, USR = ?, Tglaproval = ?, Status_Approval = ?, AspekSudahBagus = ?, AspekDiperbaiki = ?, ActionPlan = ? WHERE IdNilai = ?", nilaijadi, namadosen, dt.Format("01-02-2006 15:04:05"), statusdata, aspeksudahbagus, aspekdiperbaiki, actionplan, idnilai).Exec()
	if errupapp != nil {
		ResultData.Code = 1
		ResultData.Message = errupapp.Error()
		return ResultData
	}

	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = statusdata
	return ResultData
}

func Getshowmore(idnilai string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datakomp []orm.Params

	_, err := o.Raw(`SELECT pppd.dbo.MasterKomponenKompdinilai.namakompdinilai, pppd.dbo.TransKompNilai.Nilai, pppd.dbo.TransKompNilai.feedback, 
		CONVERT( nvarchar ( 50 ), pppd.dbo.TransKompNilai.Idkomponenkompdinilai) AS idkomponen
			FROM
				pppd.dbo.TransKompNilai
				LEFT JOIN pppd.dbo.MasterKomponenKompdinilai ON pppd.dbo.TransKompNilai.Idkomponenkompdinilai = pppd.dbo.MasterKomponenKompdinilai.idkomponenkompdinilai
				WHERE pppd.dbo.TransKompNilai.IdNilai = ?`, idnilai).Values(&datakomp)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = datakomp
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Rekappenilaian(idmahasiswa, idbagian string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datarekap []orm.Params

	_, err := o.Raw(`SELECT 
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransKompNilai.IdMahasiswa ) AS idmahasiswa,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransKompNilai.IdBagian ) AS idbagian,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransKompNilai.IdKegiatan ) AS idkegiatan,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransKompNilai.Idkomponenkompdinilai ) AS idkomponen,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransKompNilai.IdKompetensi ) AS idkompetensi,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransKompNilai.IdNilai ) AS IdNilai,
			pppd.dbo.nilaicoas_pengajuan.Nilai AS nilaijadi,
			pppd.dbo.masterbagianrs.NamaBagian AS namabagianrs,
			pppd.dbo.masterkegiatan.NamaKegiatan AS namakegiatanpenilaian,
			pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi AS namakompetensipenilaian,
			pppd.dbo.MasterKomponenKompdinilai.namakompdinilai AS namakomponen,
			pppd.dbo.TransKompNilai.Nilai AS nilai,
			pppd.dbo.TransKompNilai.Usr AS Dosen,
			pppd.dbo.TransKompNilai.Status_Data AS statusdata 
		FROM
			pppd.dbo.TransKompNilai
			LEFT JOIN pppd.dbo.nilaicoas_pengajuan ON pppd.dbo.TransKompNilai.IdNilai = pppd.dbo.nilaicoas_pengajuan.IdNilai
			LEFT JOIN pppd.dbo.masterbagianrs ON pppd.dbo.TransKompNilai.IdBagian = pppd.dbo.masterbagianrs.IdBag
			LEFT JOIN pppd.dbo.masterkegiatan ON pppd.dbo.TransKompNilai.IdKegiatan = pppd.dbo.masterkegiatan.IdKegiatan
			LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON pppd.dbo.TransKompNilai.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi
			LEFT JOIN pppd.dbo.MasterKomponenKompdinilai ON pppd.dbo.TransKompNilai.Idkomponenkompdinilai = pppd.dbo.MasterKomponenKompdinilai.idkomponenkompdinilai 
		WHERE
			pppd.dbo.TransKompNilai.IdMahasiswa = ?
			AND pppd.dbo.TransKompNilai.IdBagian = ?`, idmahasiswa, idbagian).Values(&datarekap)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = datarekap
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Validasipenilaian(iddosen, idnilai string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params

	_, err := o.Raw(`SELECT CONVERT
	( nvarchar ( 50 ), pppd.dbo.mastermahasiswa.MahasiswaID ) AS idmhs,
	pppd.dbo.mastermahasiswa.NO_MHS,
	pppd.dbo.mastermahasiswa.NM_MHS,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.periodecoas.IdPeriode ) AS idperiode,
	pppd.dbo.periodecoas.NamaPeriode,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterbagianrs.IdBag ) AS idbagian,
	pppd.dbo.masterbagianrs.NamaBagian,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterrumahsakit.IdRs ) AS idrs,
	pppd.dbo.masterrumahsakit.NamaRS,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.masterkegiatan.IdKegiatan ) AS idkegiatan,
	pppd.dbo.masterkegiatan.NamaKegiatan,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.nilaicoas_pengajuan.IdNilai ) AS idnilai,
	pppd.dbo.nilaicoas_pengajuan.Nilai,
	pppd.dbo.nilaicoas_pengajuan.USR,
	pppd.dbo.nilaicoas_pengajuan.Tglmsk,
	pppd.dbo.nilaicoas_pengajuan.Tglaproval,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval ) iddosen,
	pppd.dbo.nilaicoas_pengajuan.Status_Approval,
	pppd.dbo.nilaicoas_pengajuan.Tgl_kegiatan,
	kepegawaian.dbo.gelar ( pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval ) AS NAMALENGKAP_GELAR,
	pppd.dbo.nilaicoas_pengajuan.NamaPasien,
	pppd.dbo.nilaicoas_pengajuan.JK_Pasien,
	pppd.dbo.nilaicoas_pengajuan.UsiaPasien,
	CONVERT ( nvarchar ( 50 ), pppd.dbo.MasterKompetensiPenilaian.IdKompetensi ) AS idkompetensi,
	pppd.dbo.MasterKompetensiPenilaian.Keterangan,
	pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi 
	FROM
		pppd.dbo.masterkegiatan
		INNER JOIN pppd.dbo.nilaicoas_pengajuan ON pppd.dbo.masterkegiatan.IdKegiatan = pppd.dbo.nilaicoas_pengajuan.IdKegiatan
		INNER JOIN pppd.dbo.mastermahasiswa
		INNER JOIN pppd.dbo.penempatan ON pppd.dbo.mastermahasiswa.MahasiswaID = pppd.dbo.penempatan.MahasiswaID
		INNER JOIN pppd.dbo.masterrumahsakit ON pppd.dbo.penempatan.IdRs = pppd.dbo.masterrumahsakit.IdRs
		INNER JOIN pppd.dbo.masterbagianrs ON pppd.dbo.penempatan.IdBag = pppd.dbo.masterbagianrs.IdBag
		INNER JOIN pppd.dbo.periodecoas ON pppd.dbo.penempatan.IdPeriode = pppd.dbo.periodecoas.IdPeriode ON pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = pppd.dbo.penempatan.MahasiswaID 
		AND pppd.dbo.nilaicoas_pengajuan.IdRs = pppd.dbo.penempatan.IdRs 
		AND pppd.dbo.nilaicoas_pengajuan.IdBag = pppd.dbo.penempatan.IdBag 
		AND pppd.dbo.nilaicoas_pengajuan.IdPeriode = pppd.dbo.penempatan.IdPeriode
		LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON (
			pppd.dbo.nilaicoas_pengajuan.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi 
			OR ( pppd.dbo.nilaicoas_pengajuan.IdKompetensi IS NULL AND pppd.dbo.MasterKompetensiPenilaian.IdKompetensi IS NULL ) 
		) 
	WHERE
		( pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval = ? AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' AND pppd.dbo.nilaicoas_pengajuan.IdNilai = ? ) `, iddosen, idnilai).Values(&datanilai)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = datanilai
	return ResultData
}

func Countpenilaianmhs(idmahasiswa string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params

	_, errdiajukan := o.Raw(`SELECT
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.nilaicoas_pengajuan 
	WHERE
		pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diajukan' 
		AND pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = ? 
		AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' 
	) AS hasildiajukan,
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.nilaicoas_pengajuan 
	WHERE
		pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diinput Dosen' 
		AND pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = ? 
		AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' 
	) AS hasildinilai,
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.nilaicoas_pengajuan 
	WHERE
		pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Validasi Admin' 
		AND pppd.dbo.nilaicoas_pengajuan.MahasiswaID_Pemohon = ? 
	AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' 
	) AS hasilvalidasi`, idmahasiswa, idmahasiswa, idmahasiswa).Values(&datanilai)
	if errdiajukan != nil {
		ResultData.Code = 1
		ResultData.Message = errdiajukan.Error()
		return ResultData
	}

	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = datanilai
	return ResultData
}

func Countpenilaiandsn(iddosen string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params

	_, errdiajukan := o.Raw(`SELECT
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.nilaicoas_pengajuan 
	WHERE
		pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diajukan' 
		AND pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval = ? 
		AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' 
	) AS hasildiajukan,
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.nilaicoas_pengajuan 
	WHERE
		pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Diinput Dosen' 
		AND pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval = ? 
		AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' 
	) AS hasildinilai,
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.nilaicoas_pengajuan 
	WHERE
		pppd.dbo.nilaicoas_pengajuan.Status_Approval = 'Validasi Admin' 
		AND pppd.dbo.nilaicoas_pengajuan.Pegawaiid_aproval = ? 
	AND pppd.dbo.nilaicoas_pengajuan.Status_Data = 'Aktif' 
	) AS hasilvalidasi`, iddosen, iddosen, iddosen).Values(&datanilai)
	if errdiajukan != nil {
		ResultData.Code = 1
		ResultData.Message = errdiajukan.Error()
		return ResultData
	}

	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = datanilai
	return ResultData
}
