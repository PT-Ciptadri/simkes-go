package siswacontroller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"simkes-go/config/database"
	"simkes-go/helpers"
	"simkes-go/models/response"
	"time"
)

func GetSiswa(c *fiber.Ctx) error {
	var (
		Siswa []response.Siswa
		count int64
	)

	db := database.DB
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := db.WithContext(ctx).Model(&Siswa).Find(&Siswa).Count(&count).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// Memeriksa role dari token yang sudah di ekstrak menjadi claims
	role := c.Locals("RoleId")
	roleStr, ok := role.(string)
	if !ok || roleStr == "4" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  403,
			"message": "forbidden access",
		})
	}

	res := helpers.GetResponse(200, fiber.Map{
		"Message": "Data Fetch Successfully",
		"Data":    Siswa,
	}, nil)
	return c.JSON(res)

}

//func PostSiswa(c *fiber.Ctx) error {
//	db := database.DB
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	siswaReq := request.Siswa{}
//	if err := c.BodyParser(&siswaReq); err != nil {
//		return c.JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	insertSiswa := request.Siswa{
//		Nama:          siswaReq.Nama,
//		NISN:          siswaReq.NISN,
//		Kelas:         siswaReq.Kelas,
//		TahunAjaranID: siswaReq.TahunAjaranID,
//		JurusanID:     siswaReq.JurusanID,
//		NamaOrtu:      siswaReq.NamaOrtu,
//		NoTelpOrtu:    siswaReq.NoTelpOrtu,
//		NoTelpSiswa:   siswaReq.NoTelpSiswa,
//		Point:         siswaReq.Point,
//		CreatedAt:     time.Now(),
//	}
//
//	err := db.WithContext(ctx).Create(&insertSiswa).Error
//	if err != nil {
//		return c.Status(500).JSON(fiber.Map{
//			"Message": "Failed To Create Data",
//		})
//	}
//	return c.JSON(helpers.GetResponse(200, fiber.Map{
//		"Message": "Data Created Successfully",
//		"Data":    insertSiswa,
//	}, nil))
//}
//
//func UpdateSiswa(c *fiber.Ctx) error {
//	db := database.DB
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	// Parse request body untuk mendapatkan data siswa yang ingin di-update
//	siswaReq := request.Siswa{}
//	if err := c.BodyParser(&siswaReq); err != nil {
//		return c.JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	// Buat map dari data siswa yang ingin di-update
//	updateData := map[string]interface{}{
//		"Nama":          siswaReq.Nama,
//		"NISN":          siswaReq.NISN,
//		"Kelas":         siswaReq.Kelas,
//		"TahunAjaranID": siswaReq.TahunAjaranID,
//		"JurusanID":     siswaReq.JurusanID,
//		"NamaOrtu":      siswaReq.NamaOrtu,
//		"NoTelpOrtu":    siswaReq.NoTelpOrtu,
//		"NoTelpSiswa":   siswaReq.NoTelpSiswa,
//		"Point":         siswaReq.Point,
//		"CreatedAt":     time.Now(), // Anda mungkin tidak perlu memperbarui CreatedAt
//	}
//
//	// Lakukan update data siswa berdasarkan ID
//	if err := db.WithContext(ctx).Model(&response.Siswa{}).Where("id = ?", siswaReq.ID).Updates(updateData).Error; err != nil {
//		return c.Status(500).JSON(fiber.Map{
//			"Message": "Failed To Update Data",
//			"Error":   err.Error(),
//		})
//	}
//
//	// Jika update berhasil, kembalikan respons yang sesuai
//	return c.JSON(fiber.Map{
//		"Message": "Data Updated Successfully",
//		"Data":    updateData, // atau Anda bisa memberikan respons yang berisi data yang baru di-update
//	})
//}
//
//func DeleteSiswa(c *fiber.Ctx) error {
//	var (
//		Siswa []response.Siswa
//	)
//
//	db := database.DB
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	// Parse request body untuk mendapatkan data siswa yang ingin di-update
//	siswaReq := request.Siswa{}
//	if err := c.BodyParser(&siswaReq); err != nil {
//		return c.JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	err := db.WithContext(ctx).Where("id = ?", siswaReq.ID).Delete(&Siswa).Error
//	if err != nil {
//		return c.Status(500).JSON(fiber.Map{
//			"Message": "Failed To Delete Data",
//		})
//	}
//	return c.JSON(fiber.Map{
//		"Message": "Data Deleted Successfully",
//	})
//}
