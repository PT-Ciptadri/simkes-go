package request

type Users struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type Siswa struct {
	Nama        string `json:"nama" validate:"required"`
	NISN        string `json:"nisn" validate:"required"`
	Kelas       string `json:"kelas" validate:"required"`
	NoTelpSiswa string `json:"no_telp_siswa" validate:"required,max=13"`
	JurusanID   int    `json:"jurusan_id" validate:"required"`
	TahunAjarID int    `json:"tahun_ajar_id" validate:"required"`
	NamaOrtu    string `json:"nama_ortu" validate:"required"`
	NoTelpOrtu  string `json:"no_telp_ortu" validate:"required,max=13"`
}

type Guru struct {
	Nama string `json:"nama" validate:"required"`
	NIK  string `json:"nik" validate:"required"`
}

type Jurusan struct {
	Id int `json:"id"`
}

type LoginAs struct {
	Login string `json:"login" validate:"required"`
}

func (Siswa) TableName() string {
	return "siswa"
}
