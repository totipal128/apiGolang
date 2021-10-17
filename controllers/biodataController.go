package controllers

import (
	"fmt"
	"golang3/koneksi"
	"golang3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var(
	db = koneksi.GetDatabases()
)

func ViewBiodata(contex *gin.Context){
	var viewBiodata []models.Biodata
	contex.JSON(http.StatusOK, gin.H{
		"tittle" : "data berhasil di ambil",
		"data"		:db.Find(&viewBiodata) ,
 	})

	//  fmt.Println(viewBiodata.)
}

func InsertBiodata(contex *gin.Context){
	var insertBiodata models.Biodata

	fmt.Println(insertBiodata)
	
	if err := contex.ShouldBindJSON(&insertBiodata); err != nil {
		contex.SecureJSON(http.StatusUnauthorized,gin.H{
			"message" : err.Error(),
		})
		return
	}

	result := db.Create(&insertBiodata)
	
	if result.Error != nil {
		contex.SecureJSON(http.StatusOK, gin.H{
			"message " : result.Error,
		})
	}

	contex.JSON(http.StatusOK, gin.H{
		"message" : "data Berhasil Di simpan",
		"data "		: insertBiodata,
	})
}

func ViewOneBiodata(context *gin.Context){
	var viewOneBiodata models.Biodata

	id := context.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&viewOneBiodata).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : "data yang anda cari tidak di temukan",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message"  : "data berhasil di Ambil",
		"data"     : db.Find(&viewOneBiodata),
		"status"   : http.StatusOK,
	})	
}

func UpdateBiodata(contex *gin.Context){
	var updateBiodata models.Biodata

	id := contex.Params.ByName("id")

	if err:=db.Where("id = ? ", id).First(&updateBiodata).Error; err !=nil {
		contex.JSON(http.StatusBadRequest, gin.H{
			"message"  : "data tidak di temukan",
		})
		return
	}

	if err := contex.ShouldBindJSON(&updateBiodata); err != nil {
		contex.JSON(http.StatusUnauthorized, gin.H{
			"error"  : err.Error(),
			"tittle" : "1",
		})
		return
	}

	result :=db.Save(&updateBiodata)
	if result.Error != nil {
		contex.JSON(http.StatusOK, gin.H{
			"message"  : result.Error,
			"tittle"   : "2",
		})
		return
	}

	contex.JSON(http.StatusOK, gin.H{
		"message"  : "data berhasil di update",
		"data"     : updateBiodata,
		"status"   : http.StatusOK, 
	})
}

func DeleteBiodata(contex *gin.Context){
	var deleteBiodata models.Biodata

	id := contex.Params.ByName("id")

	if err:= db.Where("id = ?", id).Delete(&deleteBiodata).Error; err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{
			"message"  : "data tidak ditemukan",
		})
		return
	}

	contex.JSON(http.StatusOK,gin.H{
		"message"  : "data berhasil di hapus",
	})
}