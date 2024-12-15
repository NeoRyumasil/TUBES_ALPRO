package main

func cariUser(username, password string, A tabAkunPengguna, n int) bool {
	/**Fungsi mencari username dan password menggunakan sequential search
	  akan mengembalikan true apabila username dan password ketemu
	  akan mengembalikan false apabila username atau password tidak ketemu **/
	var ketemu bool
	var i int

	ketemu = false
	for !ketemu && i <= n {
		if A[i].username == username && A[i].password == password {
			ketemu = true
		}
		i++
	}
	return ketemu
}

func roleCheck(username, password string, A tabAkunPengguna, n int) string {
	/**Fungsi mencari role rengguna menggunakan sequential search
	  akan mengembalikan role dari pengguna **/
	var role string
	var ketemu bool
	var i int

	ketemu = false
	for !ketemu && i <= n {
		if A[i].username == username && A[i].password == password {
			role = A[i].role
			ketemu = true
		}
		i++
	}

	return role
}

func cariUserBinary(id int, A tabAkunPengguna, n int) akunPengguna {
	/** Fungsi mencari username pengguna menggunakan binary search
		akan mengembalikan tipe data akunPengguna yang merupakan struct
	**/
	var dataUser akunPengguna
	var min, max, mid int
	dataUser.id = -1
	min = 0
	max = n
	mid = (min + max) / 2
	for min <= max && A[mid].id != id {
		if A[mid].id > id {
			max = mid - 1
		} else if A[mid].id < id {
			min = mid + 1
		}
		mid = (min + max) / 2
	}

	if A[mid].id == id {
		dataUser.id = A[mid].id
		dataUser.username = A[mid].username
		dataUser.password = A[mid].password
		dataUser.role = A[mid].role
	}
	return dataUser
}

func cariPertanyaan(id int, A tabPertanyaan, totalTag int) pertanyaan {
	/**Fungsi mencari pertanyaan berdasarkan ID menggunakan sequential search
		akan mengembalikan tipe data pertanyaan yang merupakan struct
	**/
	var ask pertanyaan
	var ketemu bool
	var i, j int

	i = 0
	ketemu = false
	for !ketemu && i < totalTag {
		j = 0
		for !ketemu && j < A[i].nPertanyaan {
			if A[i].id[j] == id {
				ketemu = true
			} else {
				j++
			}
		}
		i++
	}

	if ketemu {
		ask.id[0] = A[i-1].id[j]
		ask.nPertanyaan = A[i-1].nPertanyaan
		ask.tag = A[i-1].tag
		ask.deskripsi[0] = A[i-1].deskripsi[j]
		ask.tanggapan[0] = A[i-1].tanggapan[j]
	} else {
		ask.id[0] = -1
	}

	return ask
}
