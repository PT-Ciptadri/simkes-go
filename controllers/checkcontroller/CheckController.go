package checkcontroller

import (
	"simkes-go/config/database"
	"simkes-go/models/entity"
)

func JurusanCheck(jurusanID int) bool {
	var jurusan entity.Jurusan
	result := database.DB.First(&jurusan, jurusanID)
	return result.Error == nil
}

func TahunAjarCheck(TahunAjarID int) bool {
	var tahunAjar entity.TahunAjar
	result := database.DB.First(&tahunAjar, TahunAjarID)
	return result.Error == nil
}
