package main

import (
	"fmt"
)

func PunyaAkun(totalPengguna int, pengguna tabAkunPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, username, role string) {
	//procedure untuk mengetahui apakah pengguna punya akun atau bukan, dan apakah ingin mendaftarkan akun atau tidak.
	var PunyaAkun, Daftar string
	var UdahDaftar bool
	//tampilan awal procedure
	fmt.Println("-----------------------------------")
	fmt.Println("LOGIN APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("-----------------------------------")
	fmt.Print("Punya Akun? (Ya/Tidak) : ")
	fmt.Scan(&PunyaAkun)

	/**If statement untuk mengetahui apakah pengguna yang masuk sudah punya akun atau tidak
	Apabila sudah punya akun akan masuk ke tampilan login
	Apabila tidak punya akan ditanya apakah ingin mendaftar atau tidak
	Apabila ingin mendaftar maka akan ke tampilan DaftarAkun, apabila tidak akan masuk sebagai guest**/
	if PunyaAkun == "Ya" || PunyaAkun == "ya" {
		UdahDaftar = true
		Login(pengguna, totalPengguna, totalPertanyaan, totalTag, question)
	} else if PunyaAkun == "Tidak" || PunyaAkun == "tidak" {
		fmt.Print("Apakah Anda Ingin Mendaftar (Ya/Tidak) : ")
		fmt.Scan(&Daftar)
		if Daftar == "Ya" || Daftar == "ya" {
			DaftarAkun(&totalPengguna, &pengguna, UdahDaftar, totalPertanyaan, totalTag, question)
		} else {
			UdahDaftar = false
			TampilanMenu(UdahDaftar, pengguna, totalPengguna, totalPengguna, totalTag, question, username, role)
		}
	}
}

func Login(pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan) {
	/**procedure login, masuk menggunakan username dan password.
	  apabila username dan password benar akan masuk ke tampilan menu
	  apabila username atau password salah akan diminta kembali memasukkan username dan password**/
	var UdahDaftar bool
	var username, password, role string
	fmt.Println("-----------------------------------")
	fmt.Println("LOGIN APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("-----------------------------------")
	fmt.Print("Username 	: ")
	fmt.Scan(&username)
	fmt.Print("Password	: ")
	fmt.Scan(&password)
	if cariUser(username, password, pengguna, totalPengguna) {
		UdahDaftar = true
		role = roleCheck(username, password, pengguna, totalPengguna)
		if role == "Pasien" {
			TampilanMenu(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, username, role)
		} else if role == "Admin" {
			tampilanAdmin(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, username, role)
		} else if role == "Dokter" {
			tampilanDokter(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, username, role)
		}
	}
}

func TampilanMenu(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, user, role string) {
	/**procedure untuk menampilkan menu utama aplikasi**/
	// clearscreen terminal
	var pilihan int
	fmt.Println("----------------------------------")
	fmt.Println("MENU APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("----------------------------------")
	fmt.Println("1. FORUM KONSULTASI KESEHATAN")
	fmt.Println("2. Log Out")
	fmt.Println("3. Exit")
	fmt.Print("Tentukan Pilihanmu (1/2/3) : ")
	fmt.Scan(&pilihan)

	/** if statement
	   apabila inputnya 1 maka menampilkan procedure ForumKonsultasiKesehatan
	   apabila inputnya 2 maka akan kembali ke tampilan login
	   apabila inputnya 3 akan keluar dari aplikasi
	**/

	if pilihan == 1 {
		ForumKonsultasiKesehatan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 2 {
		Login(pengguna, totalPengguna, totalPertanyaan, totalTag, question)
	} else {
		Exit()
	}
}

func ForumKonsultasiKesehatan(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, user, role string) {
	/**procedure untuk menampilkan Forum Konsultasi Kesehatan**/
	var pilihan int
	fmt.Println("--------------------------")
	fmt.Println("FORUM KONSULTASI KESEHATAN")
	fmt.Println("--------------------------")
	fmt.Println("1. CARI KELUHANMU")
	fmt.Println("2. POSTING DAN TANYAKAN KELUHANMU (Harus Login Dulu)")
	fmt.Println("3. EXIT")
	fmt.Print("Tentukan Pilihanmu (1/2/3) : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 2 {
		if UdahDaftar {
			postingPertanyaan(UdahDaftar, pengguna, &totalPertanyaan, totalPengguna, totalTag, &question, user, role)
		} else {
			var mauDaftar string
			fmt.Println("Anda belum bisa menggunakan fitur ini karena belum mendaftar")
			fmt.Print("Apakah Anda ingin mendaftar (Ya/Tidak) : ")
			fmt.Scan(&mauDaftar)
			if mauDaftar == "Ya" {
				DaftarAkun(&totalPengguna, &pengguna, UdahDaftar, totalPertanyaan, totalTag, question)
			}
		}
	} else {
		Exit()
	}
}

func TampilanKeluhan(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, user, role string) {
	/** procedure untuk menampilkan keluhan-keluhan pasien
		bisa diurutkan berdasarkan tag terbanyk, tag tersedikit, maupun tang yang dicari
	**/
	var pilihan, id int
	var tag, Tanggapi string
	fmt.Println("-----------------------------------")
	fmt.Println("   FORUM KONSULTASI KESEHATAN      ")
	fmt.Println("-----------------------------------")
	fmt.Println("1. LIHAT BERDASARKAN TAG TERBANYAK")
	fmt.Println("2. LIHAT BERDASARKAN TAG TERSEDIKIT")
	fmt.Println("3. LIHAT BERDASARKAN TAG")
	fmt.Println("4. KEMBALI")
	fmt.Print("Tentukan Pilihanmu (1/2/3/4) : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {

		// memanggil procedure untuk mengurutkan secara ascending berdasarkan banyak tag
		sortPertanyaanAscending(&question, totalTag)
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Printf("%-6s | %-6s | %-54s | %-9s | \n", "ID", "Tag", "Deskripsi", "Tanggapan")
		fmt.Println("--------------------------------------------------------------------------------------")
		for i := 0; i < totalTag; i++ {
			for j := 0; j < question[i].nPertanyaan; j++ {
				fmt.Printf("%-6d | %-6s | %-54s | %-9d | \n", question[i].id[j], question[i].tag[j], question[i].deskripsi[j], question[i].tanggapan[j].nTanggapan)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Print("Lihat Tanggapan? (Y/T) : ")
		fmt.Scan(&Tanggapi)
		if Tanggapi == "Y" && UdahDaftar {
			// memanggil procedure tanggapiKeluhan untuk memberikan tanggapan
			fmt.Print("Masukkan ID Keluhan : ")
			fmt.Scan(&id)
			TampilanTanggapan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, id, question, user, role)
		} else if !UdahDaftar {
			fmt.Println("Kamu Belum Login")
			TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
		} else {
			TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
		}
	} else if pilihan == 2 {

		// memanggil procedure untuk mengurutkan secara descending berdasarkan banyak tag
		sortPertanyaanDescending(&question, totalTag)
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Printf("%-6s | %-6s | %-54s | %-9s | \n", "ID", "Tag", "Deskripsi", "Tanggapan")
		fmt.Println("---------------------------------------------------------------------------------------")
		for i := 0; i < totalTag; i++ {
			for j := 0; j < question[i].nPertanyaan; j++ {
				fmt.Printf("%-6d | %-6s | %-54s | %-9d | \n", question[i].id[j], question[i].tag[j], question[i].deskripsi[j], question[i].tanggapan[j].nTanggapan)
			}
		}
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Print("Lihat Tanggapan? (Y/T) : ")
		fmt.Scan(&Tanggapi)
		if Tanggapi == "Y" && UdahDaftar {
			// memanggil procedure tanggapiKeluhan untuk memberikan tanggapan
			fmt.Print("Masukkan ID Keluhan : ")
			fmt.Scan(&id)
			TampilanTanggapan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, id, question, user, role)
		} else if !UdahDaftar {
			fmt.Println("Kamu Belum Login")
			TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
		} else {
			TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
		}
	} else if pilihan == 3 {
		fmt.Print("Masukkan Tag (Ringan/Kulit/THT) : ")
		fmt.Scan(&tag)

		if tag == "Ringan" {
			fmt.Println("--------------------------------------------------------------------------------")
			fmt.Printf("%-6s | %-6s | %-54s | %-9s | \n", "ID", "Tag", "Deskripsi", "Tanggapan")
			fmt.Println("--------------------------------------------------------------------------------")
			for i := 0; i < question[0].nPertanyaan; i++ {
				fmt.Printf("%-6d | %-6s | %-54s | %-9d | \n", question[0].id[i], question[0].tag[i], question[0].deskripsi[i], question[0].tanggapan[i].nTanggapan)
			}
			fmt.Println("--------------------------------------------------------------------------------")
		} else if tag == "Kulit" {
			fmt.Println("--------------------------------------------------------------------------------")
			fmt.Printf("%-6s | %-6s | %-54s | %-9s | \n", "ID", "Tag", "Deskripsi", "Tanggapan")
			fmt.Println("--------------------------- ----------------------------------------------------")
			for i := 0; i < question[1].nPertanyaan; i++ {
				fmt.Printf("%-6d | %-6s | %-54s | %-9d | \n", question[1].id[i], question[1].tag[i], question[1].deskripsi[i], question[1].tanggapan[i].nTanggapan)
			}
			fmt.Println("--------------------------------------------------------------------------------")
		} else if tag == "THT" {
			fmt.Println("--------------------------------------------------------------------------------")
			fmt.Printf("%-6s | %-6s | %-54s | %-9s | \n", "ID", "Tag", "Deskripsi", "Tanggapan")
			fmt.Println("--------------------------------------------------------------------------------")
			for i := 0; i < question[2].nPertanyaan; i++ {
				fmt.Printf("%-6d | %-6s | %-54s | %-9d | \n", question[2].id[i], question[2].tag[i], question[2].deskripsi[i], question[2].tanggapan[i].nTanggapan)
			}
			fmt.Println("---------------------------------------------------------------------------------")
		}
		fmt.Print("Lihat Tanggapan? (Y/T) : ")
		fmt.Scan(&Tanggapi)
		if Tanggapi == "Y" && UdahDaftar {
			fmt.Print("Masukkan ID Keluhan : ")
			fmt.Scan(&id)
			TampilanTanggapan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, id, question, user, role)
		} else if !UdahDaftar {
			fmt.Println("Kamu Belum Login")
			TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
		} else {
			TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
		}
	} else if pilihan == 4 {
		ForumKonsultasiKesehatan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	}
}

func TampilanTanggapan(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag, id int, question tabPertanyaan, user, role string) {
	// procedure untuk menampilkan tanggapan suatu keluhan
	var ask pertanyaan
	var Tanggapi string

	ask = cariPertanyaan(id, question, totalTag)

	if ask.id[0] != -1 {
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("	Tanggapan Keluhan Pasien						 ")
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("Keluhan : ", ask.deskripsi)
		fmt.Println("--------------------------------------------------------------------")
		for i := 0; i <= ask.tanggapan[ask.nPertanyaan].nTanggapan; i++ {
			for j := 0; j < ask.tanggapan[i].nTanggapan; j++ {
				fmt.Printf("%s (%s) \n", ask.tanggapan[i].user[j], ask.tanggapan[i].role[j])
				fmt.Printf("= %s \n", ask.tanggapan[i].pesan[j])
				fmt.Println("----------------------------------------------------------------")
			}
		}
		fmt.Print("Tanggapi? (Y/T) : ")
		fmt.Scan(&Tanggapi)
		if Tanggapi == "Y" {
			// memanggil procedure tanggapiKeluhan untuk memberikan tanggapan
			TanggapiKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, &question, user, role)
		} else {
			TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
		}
	} else {
		fmt.Print("Keluhan yang ada cari tidak ada")
	}

}

func Exit() {
	// keluar dari program
	fmt.Print("Terimakasih Telah Berkunjung")
}
