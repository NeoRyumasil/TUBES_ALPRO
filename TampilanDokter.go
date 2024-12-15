package main

import (
	"fmt"
)

func tampilanDokter(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, user, role string) {
	// procedure tampilan untuk user yang memiliki role dokter

	var pilihan int
	fmt.Println("----------------------------------------")
	fmt.Println("MENU APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("----------------------------------------")
	fmt.Println("1. FORUM KONSULTASI KESEHATAN")
	fmt.Println("2. LOG OUT")
	fmt.Println("3. EXIT")
	fmt.Print("Tentukan Pilihanmu (1/2) : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		ForumKonsultasiKesehatan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 2 {
		Login(pengguna, totalPengguna, totalPertanyaan, totalTag, question)
	} else {
		Exit()
	}
}

func TanggapiKeluhan(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question *tabPertanyaan, user, role string) {
	// procedure untuk menanggapi keluhan pasien

	var id int
	var ask pertanyaan
	fmt.Print("Masukkan Id Pertanyaan yang ingin ditanggapi : ")
	fmt.Scan(&id)

	// memanggil fungsi untuk mencari pertanyaan berdasarkan id
	ask = cariPertanyaan(id, *question, totalPertanyaan)
	if ask.id[0] == -1 {
		fmt.Print("Pertanyaan yang anda cari tidak ada")
	} else {
		// memanggil procedure untuk memberikan tanggapan pada suatu pertanyaan
		beriTanggapan(question, totalPertanyaan, ask.id[0], user, role)
		// memanggil procedure untuk menampilkan keluhan pasien
		TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, *question, user, role)
	}
}
