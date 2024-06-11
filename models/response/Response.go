package response

import (
	"gorm.io/gorm"
	"time"
)

type Response struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type Users struct {
	Id         int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	UserroleId int    `json:"userrole_id"`
	SiswaId    int    `json:"siswa_id"`
	GuruId     int    `json:"guru_id"`
	Status     bool   `json:"status"`
}

type Siswa struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama          string    `json:"nama"`
	NISN          string    `json:"nisn"`
	Kelas         string    `json:"kelas"`
	TahunAjaranID int       `json:"tahun_ajaran_id"`
	JurusanID     int       `json:"jurusan_id"`
	NamaOrtu      string    `json:"nama_ortu"`
	NoTelpOrtu    string    `json:"no_telp_ortu"`
	NoTelpSiswa   string    `json:"no_telp_siswa"`
	Point         int       `json:"point"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	// Enables soft deletion by using the deleted_at column to mark records as deleted without permanently removing them.
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

// TableName sets the insert table name for this struct type
func (Siswa) TableName() string {
	return "siswa"
}
