package main

func sortPertanyaanAscending(A *tabPertanyaan, n int) {
	/** Sorting pertanyaan secara ascending (menaik)
		berdasarkan jumlah tag
		menggunakan selection sort
	**/
	pass := 1
	for pass < n {
		idx := pass - 1
		i := pass
		for i < n {
			if A[idx].nPertanyaan < A[i].nPertanyaan {
				idx = i
			}
			i++
		}
		temp := A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func sortPertanyaanDescending(A *tabPertanyaan, n int) {
	/** Sorting pertanyaan secara descending (menurun)
		berdasarkan jumlah tag
		menggunakan selection sort
	**/
	pass := 1
	for pass < n {
		idx := pass - 1
		i := pass
		for i < n {
			if A[idx].nPertanyaan > A[i].nPertanyaan {
				idx = i
			}
			i++
		}
		temp := A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
