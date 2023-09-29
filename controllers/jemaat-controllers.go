package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/Project_pasti/models"
	"github.com/josepnapitupulu/Project_pasti/utils"
)

var NewJemaat models.Jemaat
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetJemaat(w http.ResponseWriter, r *http.Request) {
	newJemaats := models.GetAllJemaats()
	res, _ := json.Marshal(newJemaats)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetJemaatById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatId := vars["jemaatId"]
	NIK, err := strconv.ParseInt(jemaatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaatDetails, _ := models.GetJemaatbyId(NIK)
	res, _ := json.Marshal(jemaatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateJemaat(w http.ResponseWriter, r *http.Request) {
	CreateJemaat := &models.Jemaat{}
	utils.ParseBody(r, CreateJemaat)
	b := CreateJemaat.CreateJemaat()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteJemaat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatId := vars["jemaatId"]
	ID, err := strconv.ParseInt(jemaatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaat := models.DeleteJemaat(ID)
	res, _ := json.Marshal(jemaat)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateJemaat(w http.ResponseWriter, r *http.Request) {
	var updateJemaat = &models.Jemaat{}
	utils.ParseBody(r, updateJemaat)
	vars := mux.Vars(r)
	jemaatId := vars["jemaatId"]
	ID, err := strconv.ParseInt(jemaatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jemaatDetails, db := models.GetJemaatbyId(ID)
	if updateJemaat.NamaJemaat != "" {
		jemaatDetails.NamaJemaat = updateJemaat.NamaJemaat
	}
	if updateJemaat.TempatLahir != "" {
		jemaatDetails.TempatLahir = updateJemaat.TempatLahir
	}
	if updateJemaat.JenisKelamin != "" {
		jemaatDetails.JenisKelamin = updateJemaat.JenisKelamin
	}
	if updateJemaat.Alamat != "" {
		jemaatDetails.Alamat = updateJemaat.Alamat
	}
	if updateJemaat.NoHP != "" {
		jemaatDetails.NoHP = updateJemaat.NoHP
	}
	db.Save(&jemaatDetails)
	res, _ := json.Marshal(jemaatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}