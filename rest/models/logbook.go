package models

import (
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/libra9z/orm"
)

func Getbagianlog() Result {
	o := orm.NewOrm()
	var ResultData Result
	var bagian []orm.Params

	_, err := o.Raw(`SELECT 
			CONVERT(nvarchar( 50 ),pppd.dbo.masterbagianrs.IdBag) AS idbagian,
			pppd.dbo.masterbagianrs.NamaBagian 
		FROM
			pppd.dbo.masterbagianrs
			WHERE pppd.dbo.masterbagianrs.Status = 'Aktif'`).Values(&bagian)

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

func Getkompetensilog(idbagian string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var kompetensi []orm.Params
	statusdata := "Aktif"

	_, err := o.Raw(`SELECT 
				CONVERT(nvarchar(50),pppd.dbo.MasterKompetensiPenilaian.IdKompetensi) AS idkompetensi,
				pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi,
				pppd.dbo.MasterKompetensiPenilaian.Keterangan,
				pppd.dbo.MasterKompetensiPenilaian.Loc
			FROM
				pppd.dbo.MasterKompetensiPenilaian
			WHERE pppd.dbo.MasterKompetensiPenilaian.Status_Data = ? 
			AND pppd.dbo.MasterKompetensiPenilaian.IdBag = ?`, statusdata, idbagian).Values(&kompetensi)

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

func Getdosenkliniklog(idbagian string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var dosenklinik []orm.Params
	status := "Aktif"

	_, err := o.Raw(`SELECT CONVERT(NVARCHAR(50),[PEGAWAIID]) as PEGAWAIID,[Nama],[NamaLengkap],[PANGKAT],[IDBAGIAN],[Bagian],[Status],[IDRS],[NamaRS],[Jenis]
	FROM [pppd].[dbo].[DosenKlinik]
	WHERE CONVERT(NVARCHAR(50),Jenis)='dosen' AND IDBAGIAN = ? AND Status = ?`, idbagian, status).Values(&dosenklinik)

	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = dosenklinik
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Getlogmhs(idmahasiswa, idbagian, idkompetensi, offset, statusapproval string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datalogbook []orm.Params
	status := "Aktif"

	var stringwhere string
	if idkompetensi == "" && idbagian == "" {
		stringwhere = ``
	} else if idkompetensi != "" && idbagian == "" {
		stringwhere = `AND pppd.dbo.TransLogBook.IdKompetensi = '` + idkompetensi + `' `
	} else if idkompetensi == "" && idbagian != "" {
		stringwhere = `AND pppd.dbo.TransLogBook.IdBagian = '` + idbagian + `' `
	} else {
		stringwhere = `AND pppd.dbo.TransLogBook.IdKompetensi = '` + idkompetensi + `' AND pppd.dbo.TransLogBook.IdBagian = '` + idbagian + `' `
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
		stringstatus = ` AND pppd.dbo.TransLogBook.StatusPengajuan = 'Diajukan' `
	} else if statusapproval == "2" {
		stringstatus = ` AND pppd.dbo.TransLogBook.StatusPengajuan = 'Diinput Dosen' `
	} else {
		stringstatus = ``
	}

	_, err := o.Raw(`SELECT CONVERT
				( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdTransLogBook ) AS idtranslogbook,
				CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdBagian ) AS idbagian,
				CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdMahasiswa ) AS idmahasiswa,
				CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdKompetensi ) AS idkompetensi,
				CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.Id_Pegawai ) AS idpegawai,
				pppd.dbo.TransLogBook.NamaPasien,
				pppd.dbo.TransLogBook.ProblemPasien,
				pppd.dbo.TransLogBook.JenisKelaminPasien,
				pppd.dbo.TransLogBook.SituasiRuangan,
				pppd.dbo.TransLogBook.StatusPengajuan,
				pppd.dbo.TransLogBook.Tgl_Input,
				pppd.dbo.TransLogBook.TanggalPengajuan,
				pppd.dbo.TransLogBook.TanggalApproval,
				pppd.dbo.TransLogBook.UmurPasien,
				pppd.dbo.TransLogBook.Usr,
				pppd.dbo.masterbagianrs.NamaBagian,
				pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi,
				pppd.dbo.MasterKompetensiPenilaian.Keterangan,
				pppd.dbo.MasterKompetensiPenilaian.Loc,
				pppd.dbo.DosenKlinik.NamaLengkap,
				CONVERT ( nvarchar ( 50 ), pppd.dbo.MasterTindakanKemampuan.IdTindakanKemampuan) AS IdTindakan,
				pppd.dbo.MasterTindakanKemampuan.TindakanKemampuan,
				pppd.dbo.MasterTindakanKemampuan.Keterangan AS KeteranganKemampuan
			FROM
			pppd.dbo.TransLogBook
				LEFT JOIN pppd.dbo.mastermahasiswa ON pppd.dbo.TransLogBook.IdMahasiswa = pppd.dbo.mastermahasiswa.MahasiswaID
				LEFT JOIN pppd.dbo.masterbagianrs ON pppd.dbo.TransLogBook.IdBagian = pppd.dbo.masterbagianrs.IdBag
				LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON pppd.dbo.TransLogBook.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi
				LEFT JOIN pppd.dbo.DosenKlinik ON pppd.dbo.TransLogBook.Id_Pegawai = pppd.dbo.DosenKlinik.PEGAWAIID 
				LEFT JOIN pppd.dbo.MasterTindakanKemampuan ON pppd.dbo.TransLogBook.IdTindakanKemampuan = pppd.dbo.MasterTindakanKemampuan.IdTindakanKemampuan
			WHERE
				pppd.dbo.TransLogBook.IdMahasiswa = ?
				AND pppd.dbo.TransLogBook.Status_Data = ? `+stringwhere+``+stringstatus+`
			ORDER BY pppd.dbo.TransLogBook.Tgl_Input DESC
				`+stringoffset+``+stringlimit+``, idmahasiswa, status).Values(&datalogbook)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = datalogbook
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Getlogdsn(iddosen, search, idkompetensi, offset, statusapproval string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datalogbook []orm.Params
	status := "Aktif"

	var stringwhere string
	if search == "" && idkompetensi == "" {
		stringwhere = ``
	} else if search != "" && idkompetensi == "" {
		stringwhere = ` AND pppd.dbo.mastermahasiswa.NM_MHS LIKE '%` + search + `%' `
	} else if search == "" && idkompetensi != "" {
		stringwhere = `AND pppd.dbo.TransLogBook.IdKompetensi = '` + idkompetensi + `' `
	} else {
		stringwhere = `AND pppd.dbo.TransLogBook.IdKompetensi = '` + idkompetensi + `' AND pppd.dbo.mastermahasiswa.NM_MHS LIKE '%` + search + `%' `
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
		stringstatus = ` AND pppd.dbo.TransLogBook.StatusPengajuan = 'Diajukan' `
	} else if statusapproval == "2" {
		stringstatus = ` AND pppd.dbo.TransLogBook.StatusPengajuan = 'Diinput Dosen' `
	} else {
		stringstatus = ``
	}

	_, err := o.Raw(`SELECT CONVERT
			( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdTransLogBook ) AS idtranslogbook,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdBagian ) AS idbagian,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdMahasiswa ) AS idmahasiswa,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.IdKompetensi ) AS idkompetensi,
			CONVERT ( nvarchar ( 50 ), pppd.dbo.TransLogBook.Id_Pegawai ) AS idpegawai,
			pppd.dbo.mastermahasiswa.NM_MHS,
			pppd.dbo.TransLogBook.NamaPasien,
			pppd.dbo.TransLogBook.ProblemPasien,
			pppd.dbo.TransLogBook.JenisKelaminPasien,
			pppd.dbo.TransLogBook.SituasiRuangan,
			pppd.dbo.TransLogBook.StatusPengajuan,
			pppd.dbo.TransLogBook.Tgl_Input,
			pppd.dbo.TransLogBook.TanggalPengajuan,
			pppd.dbo.TransLogBook.TanggalApproval,
			pppd.dbo.TransLogBook.UmurPasien,
			pppd.dbo.TransLogBook.Usr,
			pppd.dbo.masterbagianrs.NamaBagian,
			pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi,
			pppd.dbo.MasterKompetensiPenilaian.Keterangan,
			pppd.dbo.MasterKompetensiPenilaian.Loc,
			pppd.dbo.DosenKlinik.NamaLengkap, 
			CONVERT ( nvarchar ( 50 ), pppd.dbo.MasterTindakanKemampuan.IdTindakanKemampuan) AS IdTindakan,
			pppd.dbo.MasterTindakanKemampuan.TindakanKemampuan,
			pppd.dbo.MasterTindakanKemampuan.Keterangan AS KeteranganKemampuan
		FROM
			pppd.dbo.TransLogBook
			LEFT JOIN pppd.dbo.mastermahasiswa ON pppd.dbo.TransLogBook.IdMahasiswa = pppd.dbo.mastermahasiswa.MahasiswaID
			LEFT JOIN pppd.dbo.masterbagianrs ON pppd.dbo.TransLogBook.IdBagian = pppd.dbo.masterbagianrs.IdBag
			LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON pppd.dbo.TransLogBook.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi
			LEFT JOIN pppd.dbo.DosenKlinik ON pppd.dbo.TransLogBook.Id_Pegawai = pppd.dbo.DosenKlinik.PEGAWAIID 
			LEFT JOIN pppd.dbo.MasterTindakanKemampuan ON pppd.dbo.TransLogBook.IdTindakanKemampuan = pppd.dbo.MasterTindakanKemampuan.IdTindakanKemampuan
		WHERE
			pppd.dbo.TransLogBook.Id_Pegawai = ?
			AND pppd.dbo.TransLogBook.Status_Data = ? `+stringwhere+``+stringstatus+`
		ORDER BY pppd.dbo.TransLogBook.Tgl_Input DESC
			`+stringoffset+``+stringlimit+``, iddosen, status).Values(&datalogbook)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = datalogbook
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Createlogmhs(idmahasiswa, idbagian, idkompetensi, iddosen string) Result {
	o := orm.NewOrm()
	var ResultData Result
	statusapproval := "Diajukan"
	statusdata := "Aktif"
	dt := time.Now()

	_, err := o.Raw(`INSERT INTO pppd.dbo.TransLogBook ( 
		IdTransLogBook, 
		IdMahasiswa, 
		IdKompetensi, 
		IdBagian, 
		Tgl_Input, 
		Id_Pegawai, 
		Status_Data,  
		StatusPengajuan, 
		TanggalPengajuan )
		VALUES
			( NEWID( ),?,?,?,?,?,?,?,? )`, idmahasiswa, idkompetensi, idbagian, dt.Format("01-02-2006 15:04:05"), iddosen, statusdata, statusapproval, dt.Format("01-02-2006 15:04:05")).Exec()

	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = statusapproval
	return ResultData
}

func Createlogdsn(idtranslogbook, iddosen, idmahasiswa, namadosen, tindakan, idtindakan string) Result {
	o := orm.NewOrm()
	var ResultData Result
	statusapproval := "Diinput Dosen"
	statusdata := "Aktif"
	dt := time.Now()

	_, err := o.Raw(`UPDATE pppd.dbo.TransLogBook SET Usr = ?, TanggalApproval = ?, StatusPengajuan = ?, TindakanKemampuan = ?, IdTindakanKemampuan = ? 
	WHERE IdTransLogBook = ? 
	AND IdMahasiswa = ? 
	AND Id_Pegawai = ? 
	AND Status_Data = ?`, namadosen, dt.Format("01-02-2006 15:04:05"), statusapproval, tindakan, idtindakan, idtranslogbook, idmahasiswa, iddosen, statusdata).Exec()

	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = statusapproval
	return ResultData
}

func Rekaplogmhs(idmahasiswa, idbagian string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datarekap []orm.Params
	status := "Aktif"

	_, err := o.Raw(`SELECT
			CONVERT(nvarchar(50),pppd.dbo.TransLogBook.IdBagian) AS idbagian,
			CONVERT(nvarchar(50),pppd.dbo.TransLogBook.IdKompetensi) AS idkompetensi,
			CONVERT(nvarchar(50),pppd.dbo.TransLogBook.IdMahasiswa) AS idmahasiswa,
			pppd.dbo.TransLogBook.TindakanKemampuan,
			COUNT ( TindakanKemampuan ) AS total,
			pppd.dbo.masterbagianrs.NamaBagian,
			pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi 
		FROM
			pppd.dbo.TransLogBook
			LEFT JOIN pppd.dbo.masterbagianrs ON pppd.dbo.TransLogBook.IdBagian = pppd.dbo.masterbagianrs.IdBag
			LEFT JOIN pppd.dbo.MasterKompetensiPenilaian ON pppd.dbo.TransLogBook.IdKompetensi = pppd.dbo.MasterKompetensiPenilaian.IdKompetensi 
		WHERE
			pppd.dbo.TransLogBook.IdMahasiswa = ?
			AND pppd.dbo.TransLogBook.IdBagian = ? 
			AND pppd.dbo.TransLogBook.Status_Data = ?
			AND pppd.dbo.TransLogBook.TindakanKemampuan IS NOT NULL 
		GROUP BY
			pppd.dbo.TransLogBook.IdBagian,
			pppd.dbo.TransLogBook.IdKompetensi,
			pppd.dbo.TransLogBook.IdMahasiswa,
			pppd.dbo.TransLogBook.TindakanKemampuan,
			pppd.dbo.masterbagianrs.NamaBagian,
			pppd.dbo.MasterKompetensiPenilaian.NamaKompetensi`, idmahasiswa, idbagian, status).Values(&datarekap)
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

func Gettindakankemampuan(idbagian string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datatindakan []orm.Params

	_, err := o.Raw(`SELECT CONVERT(nvarchar(50),pppd.dbo.MasterTindakanKemampuan.IdTindakanKemampuan) AS idtindakan, 
			pppd.dbo.MasterTindakanKemampuan.TindakanKemampuan,
			pppd.dbo.MasterTindakanKemampuan.Keterangan
			FROM pppd.dbo.MasterTindakanKemampuan WHERE IdBagian =?`, idbagian).Values(&datatindakan)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	data["item"] = datatindakan
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData
}

func Deletelogbook(idtranslogbook string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datadelete []orm.Params
	var datalogbook []orm.Params

	_, errget := o.Raw(`SELECT * FROM pppd.dbo.TransLogBook WHERE pppd.dbo.TransLogBook.IdTransLogBook = ?`, idtranslogbook).Values(&datalogbook)
	if errget != nil {
		ResultData.Code = 1
		ResultData.Message = errget.Error()
		return ResultData
	}

	if datalogbook[0]["StatusPengajuan"] != "Diajukan" {
		ResultData.Code = 1
		ResultData.Message = "Logbook Sudah Diproses, tidak dapat dihapus"
		return ResultData
	}

	_, err := o.Raw(`UPDATE pppd.dbo.TransLogBook SET Status_Data = 'Tidak Aktif' WHERE pppd.dbo.TransLogBook.IdTransLogBook = ?`, idtranslogbook).Exec()
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = datadelete
	return ResultData

}

func Updatelogbook(idtranslogbook, idbagian, iddosen, idkompetensi string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datalogbook []orm.Params
	var dataupdate []orm.Params

	_, errget := o.Raw(`Select * FROM pppd.dbo.TransLogBook WHERE TransLogBook.IdTransLogBook = ?`, idtranslogbook).Values(&datalogbook)
	if errget != nil {
		ResultData.Code = 1
		ResultData.Message = errget.Error()
		return ResultData
	}

	if datalogbook[0]["StatusPengajuan"] != "Diajukan" {
		ResultData.Code = 1
		ResultData.Message = "Logbook Sudah Diproses, tidak dapat diupdate"
		return ResultData
	}

	_, err := o.Raw(`UPDATE pppd.dbo.TransLogBook SET IdKompetensi =?, Id_Pegawai =?, IdBagian =? WHERE TransLogBook.IdTransLogBook = ?`, idkompetensi, iddosen, idbagian, idtranslogbook).Values(&dataupdate)
	if err != nil {
		ResultData.Code = 1
		ResultData.Message = err.Error()
		return ResultData
	}

	data := make(map[string]interface{})
	ResultData.Code = 0
	ResultData.Message = "Success"
	ResultData.Data = data
	return ResultData

}

func Countlogbookdsn(iddosen string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params

	_, errdiajukan := o.Raw(`SELECT
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.TransLogBook 
	WHERE
		pppd.dbo.TransLogBook.StatusPengajuan = 'Diajukan' 
		AND pppd.dbo.TransLogBook.Id_Pegawai = ? 
		AND pppd.dbo.TransLogBook.Status_Data = 'Aktif' 
	) AS hasildiajukan,
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.TransLogBook 
	WHERE
		pppd.dbo.TransLogBook.StatusPengajuan = 'Diinput Dosen' 
		AND pppd.dbo.TransLogBook.Id_Pegawai = ? 
	AND pppd.dbo.TransLogBook.Status_Data = 'Aktif' 
	) AS hasildinilai`, iddosen, iddosen).Values(&datanilai)
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

func Countlogbookmhs(idmahasiswa string) Result {
	o := orm.NewOrm()
	var ResultData Result
	var datanilai []orm.Params

	_, errdiajukan := o.Raw(`SELECT
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.TransLogBook 
	WHERE
		pppd.dbo.TransLogBook.StatusPengajuan = 'Diajukan' 
		AND pppd.dbo.TransLogBook.IdMahasiswa = ? 
		AND pppd.dbo.TransLogBook.Status_Data = 'Aktif' 
	) AS hasildiajukan,
	(
	SELECT COUNT
		( * ) 
	FROM
		pppd.dbo.TransLogBook 
	WHERE
		pppd.dbo.TransLogBook.StatusPengajuan = 'Diinput Dosen' 
		AND pppd.dbo.TransLogBook.IdMahasiswa = ? 
	AND pppd.dbo.TransLogBook.Status_Data = 'Aktif' 
	) AS hasildinilai`, idmahasiswa, idmahasiswa).Values(&datanilai)
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
