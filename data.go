package main

func DataUser(pengguna *tabAkunPengguna, totalPengguna *int) {
	// procedure yang berisi data dummy untuk user

	pengguna[*totalPengguna].username = "Admin"
	pengguna[*totalPengguna].password = "AkuAdmint123"
	pengguna[*totalPengguna].role = "Admin"
	pengguna[*totalPengguna].id = *totalPengguna
	*totalPengguna++

	pengguna[*totalPengguna].username = "Alvin"
	pengguna[*totalPengguna].password = "Alvin123"
	pengguna[*totalPengguna].role = "Dokter"
	pengguna[*totalPengguna].id = *totalPengguna
	*totalPengguna++

	pengguna[*totalPengguna].username = "Naren"
	pengguna[*totalPengguna].password = "Naren123"
	pengguna[*totalPengguna].role = "Pasien"
	pengguna[*totalPengguna].id = *totalPengguna
	*totalPengguna++

}

func Datapertanyaan(question *tabPertanyaan, totalPertanyaan *int) {
	// procedure yang berisi data dummy untuk pertannyaan

	question[0].id[question[0].nPertanyaan] = *totalPertanyaan
	question[0].tag[question[0].nPertanyaan] = "Ringan"
	question[0].deskripsi[question[0].nPertanyaan] = "Batuk kering sudah 1 minggu obatnya apa?"
	question[0].nPertanyaan++
	*totalPertanyaan++

	question[0].id[question[0].nPertanyaan] = *totalPertanyaan
	question[0].tag[question[0].nPertanyaan] = "Ringan"
	question[0].deskripsi[question[0].nPertanyaan] = "Anak umur 4 bulan demam cara penanganannya gimana?"
	question[0].nPertanyaan++
	*totalPertanyaan++

	question[0].id[question[0].nPertanyaan] = *totalPertanyaan
	question[0].tag[question[0].nPertanyaan] = "Ringan"
	question[0].deskripsi[question[0].nPertanyaan] = "Saya gak bisa tidur, mana besok ujian"
	question[0].nPertanyaan++
	*totalPertanyaan++

	question[1].id[question[1].nPertanyaan] = *totalPertanyaan
	question[1].tag[question[1].nPertanyaan] = "Kulit"
	question[1].deskripsi[question[1].nPertanyaan] = "Kulit gatal-gatal harus diapain ya?"
	question[1].nPertanyaan++
	*totalPertanyaan++

	question[1].id[question[1].nPertanyaan] = *totalPertanyaan
	question[1].tag[question[1].nPertanyaan] = "Kulit"
	question[1].deskripsi[question[1].nPertanyaan] = "Kulit saya panuan obatnya apa?"
	question[1].nPertanyaan++
	*totalPertanyaan++

	question[2].id[question[2].nPertanyaan] = *totalPertanyaan
	question[2].tag[question[2].nPertanyaan] = "THT"
	question[2].deskripsi[question[2].nPertanyaan] = "Telinga terasa gatal sehingga tidak nyaman"
	question[2].tanggapan[question[2].nPertanyaan].pesan[question[2].tanggapan[question[2].nPertanyaan].nTanggapan] = "Cek dokter untuk pemeriksaan telinga"
	question[2].tanggapan[question[2].nPertanyaan].user[question[2].tanggapan[question[2].nPertanyaan].nTanggapan] = "Alvin"
	question[2].tanggapan[question[2].nPertanyaan].role[question[2].tanggapan[question[2].nPertanyaan].nTanggapan] = "Dokter"
	question[2].tanggapan[question[2].nPertanyaan].nTanggapan++
	question[2].nPertanyaan++
}
