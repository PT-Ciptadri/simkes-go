package authcontroller

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"os"
	"simkes-go/config/database"
	"simkes-go/controllers/checkcontroller"
	"simkes-go/helpers"
	"simkes-go/models/entity"
	"simkes-go/models/request"
	"simkes-go/models/response"
	"strings"
	"time"
)

//func Register(c *fiber.Ctx) error {
//
//	db := database.DB
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	usersReq := request.Users{}
//	if err := c.BodyParser(&usersReq); err != nil {
//		return c.JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	// Validasi kustom untuk username
//	if usersReq.Username == "" {
//		return c.JSON(fiber.Map{
//			"status": 401,
//			"error":  "pastikan username sudah benar",
//		})
//	}
//	var count int64
//	err := db.WithContext(ctx).Model(&response.Users{}).Where("username = ?", usersReq.Username).Find(&response.Users{}).Count(&count).Error
//	if err != nil {
//		return c.JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	if count > 0 {
//		return c.JSON(fiber.Map{
//			"status": 409,
//			"error":  "username telah digunakan",
//		})
//	}
//
//	password, err := helpers.HashPassword(usersReq.Password, helpers.GenerateSalt())
//	if err != nil {
//		return c.JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	insertUser := request.Users{
//		Username:   usersReq.Username,
//		Password:   password,
//		UserroleId: usersReq.UserroleId,
//		SiswaId:    usersReq.SiswaId,
//		GuruId:     usersReq.GuruId,
//		Status:     false,
//	}
//
//	errCreateUser := db.WithContext(ctx).Create(&insertUser).Error
//	if errCreateUser != nil {
//		return c.Status(500).JSON(fiber.Map{
//			"Message": "Failed To Create Data",
//		})
//	}
//
//	return c.JSON(fiber.Map{
//		"Message": "Data Created Successfully",
//		"Data":    insertUser,
//	})
//}

//func RegisterSiswa(c *fiber.Ctx) error {
//	// DB CONNECT
//	db := database.DB
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	// START USER
//
//	// USER REQUEST
//	usersReq := request.Users{}
//
//	// USER PARSING
//	if err := c.BodyParser(&usersReq); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	// VALIDATE USER
//	validate := validator.New()
//	errValidate := validate.Struct(usersReq)
//	if errValidate != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"message": "failed",
//			"error":   errValidate.Error(),
//		})
//	}
//
//	// CHECK USER DUPLICATE
//	var count int64
//	err := db.WithContext(ctx).Model(&response.Users{}).Where("username = ?", usersReq.Username).Count(&count).Error
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	if count > 0 {
//		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
//			"status": 409,
//			"error":  "username telah digunakan",
//		})
//	}
//
//	// HASH PASSWORD
//	password, err := utils.HashingPassword(usersReq.Password)
//	if err != nil {
//		log.Println(err)
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"message": "internal server error",
//		})
//	}
//
//	// END USER
//
//	// START SISWA
//
//	// SISWA REQUEST
//	siswaReq := request.Siswa{}
//
//	// SISWA PARSING
//	if err := c.BodyParser(&siswaReq); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	// VALIDATE SISWA
//	errValidate = validate.Struct(siswaReq)
//	if errValidate != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"message": "failed",
//			"error":   errValidate.Error(),
//		})
//	}
//
//	// CHECK SISWA DUPLICATE
//	err = db.WithContext(ctx).Model(&response.Siswa{}).Where("nama = ?", siswaReq.Nama).Count(&count).Error
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	if count > 0 {
//		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
//			"status": 409,
//			"error":  "nama siswa telah terdaftar",
//		})
//	}
//
//	// CHECK JURUSAN ID
//	if !checkcontroller.JurusanCheck(siswaReq.JurusanID) {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": "Invalid JurusanID",
//		})
//	}
//
//	// CHECK TAHUN AJAR ID
//	if !checkcontroller.TahunAjarCheck(siswaReq.TahunAjarID) {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": "Invalid TahunAjarID",
//		})
//	}
//
//	// END SISWA
//
//	// START INSERT
//
//	// INSERT SISWA
//	insertSiswa := entity.Siswa{
//		Nama:        siswaReq.Nama,
//		NISN:        siswaReq.NISN,
//		Kelas:       siswaReq.Kelas,
//		NoTelpSiswa: siswaReq.NoTelpSiswa,
//		JurusanID:   siswaReq.JurusanID,
//		TahunAjarID: siswaReq.TahunAjarID,
//		NamaOrtu:    siswaReq.NamaOrtu,
//		NoTelpOrtu:  siswaReq.NoTelpOrtu,
//	}
//
//	errCreateSiswa := db.WithContext(ctx).Create(&insertSiswa).Error
//	if errCreateSiswa != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"message": "Failed to create siswa",
//			"error":   errCreateSiswa.Error(),
//		})
//	}
//
//	// INSERT USER
//	insertUser := entity.Users{
//		Username: usersReq.Username,
//		Password: password,
//		SiswaId:  int(insertSiswa.ID),
//	}
//
//	errCreateUser := db.WithContext(ctx).Create(&insertUser).Error
//	if errCreateUser != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"message": "Failed to create user",
//			"error":   errCreateUser.Error(),
//		})
//	}
//
//	// END INSERT
//
//	return c.JSON(fiber.Map{
//		"message": "Data created successfully",
//	})
//}
//
//func RegisterGuru(c *fiber.Ctx) error {
//	// DB CONNECT
//	db := database.DB
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	// START USER
//
//	// USER REQUEST
//	usersReq := request.Users{}
//
//	// USER PARSING
//	if err := c.BodyParser(&usersReq); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	// VALIDATE USER
//	validate := validator.New()
//	errValidate := validate.Struct(usersReq)
//	if errValidate != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"message": "failed",
//			"error":   errValidate.Error(),
//		})
//	}
//
//	// CHECK USER DUPLICATE
//	var count int64
//	err := db.WithContext(ctx).Model(&response.Users{}).Where("username = ?", usersReq.Username).Count(&count).Error
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	if count > 0 {
//		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
//			"status": 409,
//			"error":  "username telah digunakan",
//		})
//	}
//
//	// HASH PASSWORD
//	password, err := utils.HashingPassword(usersReq.Password)
//	if err != nil {
//		log.Println(err)
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"message": "internal server error",
//		})
//	}
//
//	// END USER
//
//	// START GURU
//
//	// GURU REQUEST
//	guruReq := request.Guru{}
//
//	// GURU PARSING
//	if err := c.BodyParser(&guruReq); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	// VALIDATE GURU
//	errValidate = validate.Struct(guruReq)
//	if errValidate != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"message": "failed",
//			"error":   errValidate.Error(),
//		})
//	}
//
//	// CHECK GURU DUPLICATE
//	err = db.WithContext(ctx).Model(&response.Siswa{}).Where("nama = ?", guruReq.Nama).Count(&count).Error
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": 500,
//			"error":  err.Error(),
//		})
//	}
//
//	if count > 0 {
//		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
//			"status": 409,
//			"error":  "nama siswa telah terdaftar",
//		})
//	}
//
//	// END GURU
//
//	// START INSERT
//
//	// INSERT GURU
//	insertGuru := entity.Guru{
//		Nama: guruReq.Nama,
//		NIK:  guruReq.NIK,
//	}
//
//	errCreateSiswa := db.WithContext(ctx).Create(&insertGuru).Error
//	if errCreateSiswa != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"message": "Failed to create siswa",
//			"error":   errCreateSiswa.Error(),
//		})
//	}
//
//	// INSERT USER
//	insertUser := entity.Users{
//		Username: usersReq.Username,
//		Password: password,
//		GuruId:   int(insertGuru.ID),
//	}
//
//	errCreateUser := db.WithContext(ctx).Create(&insertUser).Error
//	if errCreateUser != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"message": "Failed to create user",
//			"error":   errCreateUser.Error(),
//		})
//	}
//
//	// END INSERT
//
//	return c.JSON(fiber.Map{
//		"message": "Data created successfully",
//	})
//}

func Register(c *fiber.Ctx) error {
	// DB CONNECT
	db := database.DB
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// START USER

	// USER REQUEST
	usersReq := request.Users{}

	// USER PARSING
	if err := c.BodyParser(&usersReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// VALIDATE USER
	validate := validator.New()
	errValidate := validate.Struct(usersReq)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// CHECK USER DUPLICATE
	var count int64
	err := db.WithContext(ctx).Model(&response.Users{}).Where("username = ?", usersReq.Username).Count(&count).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status": 409,
			"error":  "username telah digunakan",
		})
	}

	// HASH PASSWORD
	password, err := helpers.HashPassword(usersReq.Password, helpers.GenerateSalt())
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// END USER

	// CHECK WHO IS LOGIN (LOGIN AS)
	loginAsReq := request.LoginAs{}

	// LOGIN AS PARSING
	if err := c.BodyParser(&loginAsReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// VALIDATE LOGIN AS
	errValidate = validate.Struct(loginAsReq)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	if strings.ToLower(loginAsReq.Login) == "guru" {
		// START GURU

		// GURU REQUEST
		guruReq := request.Guru{}

		// GURU PARSING
		if err := c.BodyParser(&guruReq); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": 500,
				"error":  err.Error(),
			})
		}

		// VALIDATE GURU
		errValidate = validate.Struct(guruReq)
		if errValidate != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "failed",
				"error":   errValidate.Error(),
			})
		}

		// CHECK GURU DUPLICATE
		err = db.WithContext(ctx).Model(&response.Siswa{}).Where("nama = ?", guruReq.Nama).Count(&count).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": 500,
				"error":  err.Error(),
			})
		}

		if count > 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status": 409,
				"error":  "nama siswa telah terdaftar",
			})
		}

		// END GURU

		// START INSERT

		// INSERT GURU
		insertGuru := entity.Guru{
			Nama: guruReq.Nama,
			NIK:  guruReq.NIK,
		}

		errCreateSiswa := db.WithContext(ctx).Create(&insertGuru).Error
		if errCreateSiswa != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create siswa",
				"error":   errCreateSiswa.Error(),
			})
		}

		// INSERT USER
		insertUser := entity.Users{
			Username: usersReq.Username,
			Password: password,
			GuruId:   int(insertGuru.ID),
		}

		errCreateUser := db.WithContext(ctx).Create(&insertUser).Error
		if errCreateUser != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create user",
				"error":   errCreateUser.Error(),
			})
		}

		// END INSERT
	} else if strings.ToLower(loginAsReq.Login) == "siswa" {
		// START SISWA

		// SISWA REQUEST
		siswaReq := request.Siswa{}

		// SISWA PARSING
		if err := c.BodyParser(&siswaReq); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": 500,
				"error":  err.Error(),
			})
		}

		// VALIDATE SISWA
		errValidate = validate.Struct(siswaReq)
		if errValidate != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "failed",
				"error":   errValidate.Error(),
			})
		}

		// CHECK SISWA DUPLICATE
		err = db.WithContext(ctx).Model(&response.Siswa{}).Where("nama = ?", siswaReq.Nama).Count(&count).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": 500,
				"error":  err.Error(),
			})
		}

		if count > 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status": 409,
				"error":  "nama siswa telah terdaftar",
			})
		}

		// CHECK JURUSAN ID
		if !checkcontroller.JurusanCheck(siswaReq.JurusanID) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JurusanID",
			})
		}

		// CHECK TAHUN AJAR ID
		if !checkcontroller.TahunAjarCheck(siswaReq.TahunAjarID) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid TahunAjarID",
			})
		}

		// END SISWA

		// START INSERT

		// INSERT SISWA
		insertSiswa := entity.Siswa{
			Nama:        siswaReq.Nama,
			NISN:        siswaReq.NISN,
			Kelas:       siswaReq.Kelas,
			NoTelpSiswa: siswaReq.NoTelpSiswa,
			JurusanID:   siswaReq.JurusanID,
			TahunAjarID: siswaReq.TahunAjarID,
			NamaOrtu:    siswaReq.NamaOrtu,
			NoTelpOrtu:  siswaReq.NoTelpOrtu,
		}

		errCreateSiswa := db.WithContext(ctx).Create(&insertSiswa).Error
		if errCreateSiswa != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create siswa",
				"error":   errCreateSiswa.Error(),
			})
		}

		// INSERT USER

		insertUser := entity.Users{
			Username: usersReq.Username,
			Password: password,
			SiswaId:  int(insertSiswa.ID),
		}

		errCreateUser := db.WithContext(ctx).Create(&insertUser).Error
		if errCreateUser != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create user",
				"error":   errCreateUser.Error(),
			})
		}

		// END INSERT
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Wrong login access",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data created successfully",
	})

}

func Login(c *fiber.Ctx) error {
	// DB CONNECT
	db := database.DB
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// USER REQUEST
	usersReq := request.Users{}

	// USER PARSING
	if err := c.BodyParser(&usersReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// VALIDATE USER
	validate := validator.New()
	errValidate := validate.Struct(usersReq)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.Users
	err := db.WithContext(ctx).Model(&response.Users{}).Where("username = ?", usersReq.Username).Find(&user).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(usersReq.Password))
	if err != nil {
		res := helpers.GetResponse(fiber.StatusUnauthorized, nil, errors.New("unauthorized"))
		return c.Status(res.Status).JSON(res)
	}

	claims := jwt.MapClaims{
		"userid":     user.Id,
		"username":   user.Username,
		"userroleid": user.UserRoleId,
		"guruid":     user.GuruId,
		"siswaid":    user.SiswaId,
		"exp":        time.Now().Add(time.Hour * 500000).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if user.Status == true {
		res := helpers.GetResponse(200, fiber.Map{
			"UserId":     user.Id,
			"Username":   user.Username,
			"Token":      t,
			"UserRoleId": user.UserRoleId,
			"GuruId":     user.GuruId,
			"SiswaId":    user.SiswaId,
		}, err)
		return c.Status(res.Status).JSON(res)
	} else {
		res := helpers.GetResponse(400, nil, errors.New("akun anda belum di verifikasi"))
		return c.Status(res.Status).JSON(res)
	}

}
