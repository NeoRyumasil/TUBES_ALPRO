package main

import "fmt"

func DaftarAkun(totalPengguna *int, pengguna *tabAkunPengguna, udahDaftar bool, totalPertanyaan, totalTag int, question tabPertanyaan) {
	/**procedure untuk mendaftarkan akun dari sisi user, role akan otomatis terisi pasien
	  Setelah selesai mendaftar akan masuk ke tampilan login
	**/
	fmt.Println("-----------------------------------------")
	fmt.Println("DAFTAR AKUN APLIKASI KONSULTASI KESEHATAN")
	fmt.Println("-----------------------------------------")
	fmt.Print("Username	: ")
	fmt.Scan(&pengguna[*totalPengguna].username)
	fmt.Print("Password	: ")
	fmt.Scan(&pengguna[*totalPengguna].password)
	pengguna[*totalPengguna].role = "Pasien"
	pengguna[*totalPengguna].id = *totalPengguna
	*totalPengguna++

	//masuk ke dalam procedure login
	Login(*pengguna, *totalPengguna, totalPertanyaan, totalTag, question)
}

func postingPertanyaan(UdahDaftar bool, pengguna tabAkunPengguna, totalPertanyaan *int, totalPengguna, totalTag int, question *tabPertanyaan, user, role string) {
	/** procedure untuk memposting pertanyaan
		id setiap pertanyaan akan terisi otomatis sesuai totalpertanyaan
	**/
	var tag string
	fmt.Println("-------------------------------------")
	fmt.Print("Beri Tag kepada Keluhan Anda (Ringan/Kulit/THT): ")
	fmt.Scan(&tag)
	if tag == "Ringan" {
		fmt.Print("Masukkan Keluhan Anda : ")
		question[0].id[question[0].nPertanyaan] = *totalPertanyaan + 1
		question[0].tag[question[0].nPertanyaan] = "Ringan"
		fmt.Scan(&question[0].deskripsi[question[0].nPertanyaan])
		question[0].nPertanyaan++
	} else if tag == "Kulit" {
		fmt.Print("Masukkan Keluhan Anda : ")
		question[1].id[question[1].nPertanyaan] = *totalPertanyaan + 1
		question[1].tag[question[1].nPertanyaan] = "Kulit"
		fmt.Scan(&question[1].deskripsi[question[1].nPertanyaan])
		question[1].nPertanyaan++
	} else if tag == "THT" {
		fmt.Print("Masukkan Keluhan Anda : ")
		question[2].id[question[2].nPertanyaan] = *totalPertanyaan + 1
		question[2].tag[question[2].nPertanyaan] = "THT"
		fmt.Scan(&question[2].deskripsi[question[2].nPertanyaan])
		question[2].nPertanyaan++
	}
	*totalPertanyaan++

	// masuk ke dalam procedure TampilanKeluhan untuk melihat pertanyaan
	TampilanKeluhan(UdahDaftar, pengguna, totalPengguna, *totalPertanyaan, totalTag, *question, user, role)
}

func deleteAkun(A *tabAkunPengguna, n *int, id int) {
	// procedure untuk menghapus suatu akun pengguna berdasarkan ID user

	var index int

	for i := 0; i < *n; i++ {
		if A[i].id == id {
			index = i
		}
	}

	for i := index; i < *n; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func updateAkun(A *tabAkunPengguna, n int, id int) {
	// procedure untuk melakukan perubahan pada suatu user

	var jenis, update string
	var index int

	for i := 0; i < n; i++ {
		if A[i].id == id {
			index = i
		}
	}
	fmt.Print("Apa yang mau diubah (Username/Password/Role) : ")
	fmt.Scan(&jenis)
	if jenis == "Username" {
		fmt.Print("Masukkan Username Baru : ")
		fmt.Scan(&update)
		A[index].username = update
	} else if jenis == "Password" {
		fmt.Print("Masukkan Password Baru : ")
		fmt.Scan(&update)
		A[index].password = update
	} else if jenis == "Role" {
		fmt.Print("Masukkan Role Baru (Pasien/Dokter/Admin) : ")
		fmt.Scan(&update)
		if update == "Pasien" || update == "Dokter" || update == "Admin" {
			A[index].role = update
		} else {
			fmt.Print("Salah Input")
			updateAkun(A, n, id)
		}
	}

}

func beriTanggapan(A *tabPertanyaan, totalTag int, id int, user, role string) {
	// procedure untuk melakukan pemberian tanggapan pada suatu pertanyaan

	var update string
	var index, tag int

	for i := 0; i < totalTag; i++ {
		for j := 0; j < A[i].nPertanyaan; j++ {
			if A[i].id[j] == id {
				index = j
				tag = i
			}
		}
	}

	fmt.Println("Masukkan Tanggapan Anda Dibawah Ini")
	fmt.Print("Tanggapan : ")
	fmt.Scan(&update)
	A[tag].tanggapan[index].pesan[A[tag].tanggapan[index].nTanggapan] = update
	A[tag].tanggapan[index].user[A[tag].tanggapan[index].nTanggapan] = user
	A[tag].tanggapan[index].role[A[tag].tanggapan[index].nTanggapan] = role
	A[tag].tanggapan[index].nTanggapan++

}
