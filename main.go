package main

// Jumlah Maksimal untuk Array
const NMAX = 20

// Tipe struct untuk menyimpan tanggapan, berisi user, role, dan pesan dengan tipe data string
type tanggapan struct {
	user, role, pesan [NMAX]string
	nTanggapan        int
}

// Tipe struct untuk menyimpan akun pengguna, berisi username, password, dan role pengguna dengan tipe data string
type akunPengguna struct {
	username, password, role string
	id                       int
}

// Tipe struct untuk menyimpan pertanyaan dari pasien
type pertanyaan struct {
	tag         [NMAX]string
	nPertanyaan int
	id          [NMAX]int
	deskripsi   [NMAX]string
	tanggapan   [NMAX]tanggapan
}

// Array yang menyimpan tipe data akun pengguna
type tabAkunPengguna = [NMAX]akunPengguna

// Array yang menyimpan tipe data pertanyaan
type tabPertanyaan = [NMAX]pertanyaan

func main() {
	var totalPengguna int = 0
	var totalPertanyaan int = 0
	var totalTag int = 3
	var pengguna tabAkunPengguna
	var question tabPertanyaan
	var username, role string

	// pemanggilan procedure data dummy
	DataUser(&pengguna, &totalPengguna)
	Datapertanyaan(&question, &totalPertanyaan)

	// masuk ke dalam procedure PunyaAkun untuk mengecek apakah sudah mendaftar
	PunyaAkun(totalPengguna, pengguna, totalPertanyaan, totalTag, question, username, role)
}
