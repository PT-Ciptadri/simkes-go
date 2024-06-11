package migration

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"simkes-go/config/database"
	"simkes-go/models/entity"
	"time"
)

func RunMigration() {
	// Migrasi entity model user
	err := database.DB.AutoMigrate(
		&entity.Users{},
		&entity.Siswa{},
		&entity.SiswaPelanggar{},
		&entity.Guru{},
		&entity.PasalPelanggaran{},
		&entity.Jurusan{},
		&entity.UserRole{},
		&entity.TahunAjar{},
		&entity.Sanksi{},
	)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Database Migrated")

	seedData()
}

func seedData() {
	// DB CONNECT
	db := database.DB
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Cek apakah data sudah ada
	var count int64

	// Seed Jurusan
	db.Model(&entity.Jurusan{}).Count(&count)
	if count == 0 {
		jurusan := []entity.Jurusan{
			{Nama: "PH 1"},
			{Nama: "PH 2"},
			{Nama: "PPLG 1"},
			{Nama: "PPLG 2"},
			{Nama: "DKV 1"},
			{Nama: "DKV 2"},
			{Nama: "TO 1"},
			{Nama: "TO 2"},
			{Nama: "TO 3"},
			{Nama: "TO 4"},
			{Nama: "AKL 1"},
			{Nama: "AKL 2"},
		}
		err := db.WithContext(ctx).Create(&jurusan).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Seed TahunAjar
	db.Model(&entity.TahunAjar{}).Count(&count)
	if count == 0 {
		tahunAjar := []entity.TahunAjar{
			{Tahun: "2022/2023"},
			{Tahun: "2023/2024"},
			{Tahun: "2024/2025"},
		}
		err := db.Create(&tahunAjar).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Seed UserRole
	db.Model(&entity.UserRole{}).Count(&count)
	if count == 0 {
		userRole := []entity.UserRole{
			{Role: "Super Admin"},
			{Role: "Admin"},
			{Role: "Operator"},
			{Role: "User"},
		}
		err := db.Create(&userRole).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	fmt.Println("Data seed created successfully")
}
