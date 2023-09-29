package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/Project_pasti/config"
	)

var db *gorm.DB

type Jemaat struct {
	gorm.Model
	NIK string `gorm:""json:"nik"`
	NamaJemaat string `json:"nama_jemaat"`
	TempatLahir string `json:"tempat_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat string `json:"alamat"`
	NoHP string `json:"no_hp"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Jemaat{})
}

func (b *Jemaat) CreateJemaat() *Jemaat {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllJemaats() []Jemaat {
	var Jemaats []Jemaat
	db.Find(&Jemaats)
	return Jemaats
}

func GetJemaatbyId(nik int64) (*Jemaat, *gorm.DB) {
	var getJemaat Jemaat
	db := db.Where("NIK=?", nik).Find(&getJemaat)
	return &getJemaat, db
}

func DeleteJemaat(nik int64) Jemaat {
	var jemaat Jemaat
	db.Where("NIK=?", nik).Delete(jemaat)
	return jemaat 
}