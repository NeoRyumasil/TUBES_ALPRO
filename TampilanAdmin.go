package main

import (
	"fmt"
)

func tampilanAdmin(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, user, role string) {
	// procedure berisi tampilan untuk akun yang memiliki role admin

	var pilihan int
	fmt.Println("----------------------------------------")
	fmt.Println("MENU ADMIN APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("----------------------------------------")
	fmt.Println("1. FORUM KONSULTASI KESEHATAN")
	fmt.Println("2. KELOLA AKUN PENGGUNA")
	fmt.Println("3. LOG OUT")
	fmt.Println("4. EXIT")
	fmt.Print("Tentukan Pilihanmu (1/2/3/4) : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		ForumKonsultasiKesehatan(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 2 {
		kelolaAkunPengguna(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 3 {
		Login(pengguna, totalPengguna, totalPertanyaan, totalTag, question)
	} else {
		Exit()
	}
}

func kelolaAkunPengguna(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, user, role string) {
	// procedure yang berisi tampilan menu untuk mengelola pengguna untuk akun dengan role admin

	var pilihan int
	fmt.Println("----------------------------------------------")
	fmt.Println("KELOLA PENGGUNA APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("----------------------------------------------")
	fmt.Println("1. LIHAT AKUN PENGGUNA")
	fmt.Println("2. HAPUS AKUN PENGGUNA")
	fmt.Println("3. TAMBAH AKUN PENGGUNA")
	fmt.Println("4. EDIT AKUN PENGGUNA")
	fmt.Println("5. KEMBALI")
	fmt.Print("Tentukan Pilihanmu (1/2/3/4/5) : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		lihatAkun(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 2 {
		hapusAkun(UdahDaftar, &pengguna, &totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 3 {
		tambahAkun(UdahDaftar, &pengguna, &totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 4 {
		editAkun(UdahDaftar, &pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else if pilihan == 5 {
		tampilanAdmin(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	} else {
		fmt.Println("Salah Input")
		kelolaAkunPengguna(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	}
}

func lihatAkun(UdahDaftar bool, pengguna tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, user, role string) {
	// procedure untuk melihat seluruh user yang terdaftar
	var kembali string
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("%-4s | %-12s | %-16s | %-12s | \n", "ID", "Username", "Password", "Role")
	fmt.Println("-------------------------------------------------------")
	for i := 0; i < totalPengguna; i++ {
		fmt.Printf("%-4d | %-12s | %-16s | %-12s | \n", pengguna[i].id, pengguna[i].username, pengguna[i].password, pengguna[i].role)
	}
	fmt.Println("-------------------------------------------------------")
	fmt.Print("Kembali (Y/T) : ")
	fmt.Scan(&kembali)
	if kembali == "Y" {
		kelolaAkunPengguna(UdahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, user, role)
	}
}

func hapusAkun(UdahDaftar bool, pengguna *tabAkunPengguna, totalpengguna *int, totalPertanyaan, totalTag int, question tabPertanyaan, username, role string) {
	// procedure untuk input id pengguna yang akan digunakan untuk menghapus user

	var ID int
	var user akunPengguna
	var pilihan string
	fmt.Print("Masukkan ID akun yang ingin dihapus : ")
	fmt.Scan(&ID)

	// pemanggilan fungsi searching untuk mencari id user
	user = cariUserBinary(ID, *pengguna, *totalpengguna)
	if user.id == -1 {
		fmt.Print("User yang anda cari tidak ada, cari lagi? (Y/T) : ")
		fmt.Scan(&pilihan)
		if pilihan == "Y" {
			hapusAkun(UdahDaftar, pengguna, totalpengguna, totalPertanyaan, totalTag, question, username, role)
		}
	} else {
		fmt.Printf("Akun yang akan anda hapus adalah %s dengan role sebagai %s\n", user.username, user.role)
		fmt.Print("Apakah anda yakin ingin menghapusnya (Y/T) : ")
		fmt.Scan(&pilihan)
		if pilihan == "Y" {

			// pemanggilan procedure deleteAkun untuk menghapus suatu user berdasarkan id
			deleteAkun(pengguna, totalpengguna, user.id)
			fmt.Println("user sudah dihapus")

			// pemanggilan procedure lihatAkun untuk melihat seluruh user
			lihatAkun(UdahDaftar, *pengguna, *totalpengguna, totalPertanyaan, totalTag, question, username, role)
		} else {
			hapusAkun(UdahDaftar, pengguna, totalpengguna, totalPertanyaan, totalTag, question, username, role)
		}
	}
}

func tambahAkun(udahDaftar bool, pengguna *tabAkunPengguna, totalPengguna *int, totalPertanyaan, totalTag int, question tabPertanyaan, user, roles string) {
	/** procedure untuk mendaftarkan akun dari sisi admin
	   	username, password, dan role akan ditentukan oleh admin
	**/

	var username, password, roless string
	fmt.Println("-----------------------------------------")
	fmt.Println("DAFTAR AKUN APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("-----------------------------------------")

	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)
	fmt.Print("Masukkan role pengguna (Pasien/Admin/Dokter) : ")
	fmt.Scan(&roless)

	pengguna[*totalPengguna] = akunPengguna{
		id:       *totalPengguna + 1,
		username: username,
		password: password,
		role:     roless,
	}
	*totalPengguna++

	fmt.Println("user sudah ditambahkan")

	// pemanggilan procedure lihatAkun untuk melihat seluruh user
	lihatAkun(udahDaftar, *pengguna, *totalPengguna, totalPertanyaan, totalTag, question, user, roles)
}

func editAkun(udahDaftar bool, pengguna *tabAkunPengguna, totalPengguna, totalPertanyaan, totalTag int, question tabPertanyaan, users, role string) {
	/** procedure untuk melakukan perubahan pada suatu akun user berdasarkan id user
		dapat merubah username, password, dan role user
	**/

	var ID int
	var user akunPengguna
	var pilihan string
	fmt.Print("Masukkan ID akun yang ingin diedit : ")
	fmt.Scan(&ID)

	// pemanggilan fungsi searching untuk mencari id user
	user = cariUserBinary(ID, *pengguna, totalPengguna)
	if user.id == -1 {
		fmt.Println("User yang anda cari tidak ada")
		editAkun(udahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, users, role)
	} else {
		fmt.Printf("Akun yang akan anda edit adalah %s dengan role sebagai %s\n", user.username, user.role)
		fmt.Print("Apakah anda yakin ingin mengeditnya (Y/T) : ")
		fmt.Scan(&pilihan)
		if pilihan == "Y" {

			//pemanggilan procedure updateAkun untuk melakukan perubahan pada user
			updateAkun(pengguna, totalPengguna, ID)
			fmt.Println("user sudah diedit")
		} else {
			editAkun(udahDaftar, pengguna, totalPengguna, totalPertanyaan, totalTag, question, users, role)
		}
	}
	// pemanggilan procedure lihatAkun untuk melihat data akun
	lihatAkun(udahDaftar, *pengguna, totalPengguna, totalPertanyaan, totalTag, question, users, role)
}
