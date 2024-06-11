package routes

import (
	"github.com/gofiber/fiber/v2"
	"simkes-go/controllers/authcontroller"
	"simkes-go/controllers/siswacontroller"
	"simkes-go/helpers"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/Api")
	auth := api.Group("/Auth")
	siswa := api.Group("/Siswa")

	// AUTH //
	//auth.Post("/RegisterSiswa", authcontroller.RegisterSiswa)
	//auth.Post("/RegisterGuru", authcontroller.RegisterGuru)
	auth.Post("/Register", authcontroller.Register)
	auth.Post("/Login", authcontroller.Login)

	// SISWA //
	siswa.Post("/Get", helpers.Middleware, siswacontroller.GetSiswa)
	//siswa.Post("/", siswacontroller.PostSiswa)
	//siswa.Put("/", siswacontroller.UpdateSiswa)
	//siswa.Delete("/", siswacontroller.DeleteSiswa)

}
