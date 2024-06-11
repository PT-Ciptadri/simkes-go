package entity

import (
	"gorm.io/gorm"
	"time"
)

// NOT MASTER DATA START

type Users struct {
	Id         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string    `json:"username"`
	Password   string    `json:"password" gorm:"size:255;not null"`
	UserRoleId int       `json:"user_role_id" gorm:"default:4"` // Intentionally not defining a foreign key constraint
	SiswaId    int       `json:"siswa_id"`                      // Intentionally not defining a foreign key constraint
	GuruId     int       `json:"guru_id"`                       // Intentionally not defining a foreign key constraint
	Status     bool      `json:"status" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// Enables soft deletion by using the deleted_at column to mark records as deleted without permanently removing them.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

type Siswa struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama        string    `json:"nama"`
	NISN        string    `json:"nisn"`
	Kelas       string    `json:"kelas"`
	JurusanID   int       `json:"jurusan_id"`
	NamaOrtu    string    `json:"nama_ortu"`
	NoTelpOrtu  string    `json:"no_telp_ortu"`
	NoTelpSiswa string    `json:"no_telp_siswa"`
	TahunAjarID int       `json:"tahun_ajar_id"` // Intentionally not defining a foreign key constraint
	Point       int       `json:"point" gorm:"default:100"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Enables soft deletion by using the deleted_at column to mark records as deleted without permanently removing them.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

func (Siswa) TableName() string {
	return "siswa"
}

type SiswaPelanggar struct {
	ID                 uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SiswaId            int       `json:"siswa_id"` // Intentionally not defining a foreign key constraint
	PasalPelanggaranId int       `json:"pasal_pelanggaran_id"`
	GuruId             int       `json:"guru_id"`
	SanksiId           int       `json:"sanksi_id"` // Intentionally not defining a foreign key constraint
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	// Enables soft deletion by using the deleted_at column to mark records as deleted without permanently removing them.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

func (SiswaPelanggar) TableName() string {
	return "siswa_pelanggar"
}

type Guru struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama   string `json:"nama"`
	NIK    string `json:"nik"`
	Status bool   `json:"status" gorm:"default:false"`
}

func (Guru) TableName() string {
	return "guru"
}

// NOT MASTER DATA END

// MASTER DATA START

type PasalPelanggaran struct {
	ID               uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Kode             string `json:"kode"`
	Keterangan       string `json:"keterangan"`
	PointPelanggaran int    `json:"point_pelanggaran"`
}

func (PasalPelanggaran) TableName() string {
	return "pasal_pelanggaran"
}

type Jurusan struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama string `json:"nama"`
}

func (Jurusan) TableName() string {
	return "jurusan"
}

type UserRole struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Role string `json:"role"`
}

func (UserRole) TableName() string {
	return "user_role"
}

type TahunAjar struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Tahun string `json:"tahun"`
}

func (TahunAjar) TableName() string { return "tahun_ajar" }

type Sanksi struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Sanksi string `json:"sanksi"`
}

func (Sanksi) TableName() string { return "sanksi" }

// MASTER DATA END
