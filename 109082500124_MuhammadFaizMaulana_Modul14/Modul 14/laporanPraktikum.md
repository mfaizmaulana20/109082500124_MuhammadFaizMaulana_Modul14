# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">[MuhammadFaizMaulana] - [109082500124]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

const NMAX int = 100

func insertionSort(data *[NMAX]int, n int) {
	var i, j, key int

	for i = 1; i < n; i++ {
		key = data[i]
		j = i - 1

		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}
}

func cekJarak(data [NMAX]int, n int) {
	var i, jarak int
	var tetap bool = true

	if n < 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarak = data[1] - data[0]

	for i = 2; i < n; i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
		}
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

func main() {
	var data [NMAX]int
	var x, n int

	fmt.Scan(&x)

	for x >= 0 && n < NMAX {
		data[n] = x
		n++
		fmt.Scan(&x)
	}

	insertionSort(&data, n)

	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	cekJarak(data, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_Modul9/blob/main/modul%209/output/soal%201.png)
[Program ini membaca sekumpulan bilangan bulat positif dari input (hingga nilai negatif), kemudian mengurutkannya menggunakan algoritma insertion sort. Setelah diurutkan, program mencetak data yang sudah terurut, lalu melakukan pengecekan jarak (selisih) antara elemen-elemen yang berurutan untuk menentukan apakah data memiliki jarak yang konsisten atau tidak. Jika semua selisih antara elemen berurutan sama, maka program menampilkan nilai jaraknya; jika tidak konsisten, maka program menampilkan pesan "Data berjarak tidak tetap".]

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idxMax int

	idxMax = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}

	fmt.Println(pustaka[idxMax].judul, pustaka[idxMax].penulis, pustaka[idxMax].penerbit, pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var temp Buku

	for i = 1; i < n; i++ {
		temp = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	batas = 5
	if n < 5 {
		batas = n
	}

	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var kiri, kanan, tengah int
	var ditemukan bool

	kiri = 0
	kanan = n - 1
	ditemukan = false

	for kiri <= kanan && !ditemukan {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			ditemukan = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ditemukan {
		fmt.Println(
			pustaka[tengah].judul,
			pustaka[tengah].penulis,
			pustaka[tengah].penerbit,
			pustaka[tengah].tahun,
			pustaka[tengah].eksemplar,
			pustaka[tengah].rating,
		)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, ratingCari int

	fmt.Scan(&nPustaka)

	DaftarkanBuku(&pustaka, nPustaka)

	fmt.Scan(&ratingCari)

	CetakTerfavorit(pustaka, nPustaka)

	UrutBuku(&pustaka, nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)

	CariBuku(pustaka, nPustaka, ratingCari)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_Modul9/blob/main/modul%209/output/soal%202.png)
[Program ini merupakan sistem manajemen perpustakaan yang mengelola koleksi buku. Program membaca data buku berupa id, judul, penulis, penerbit, jumlah eksemplar, tahun, dan rating, kemudian melakukan beberapa operasi: mencari dan menampilkan buku dengan rating tertinggi, mengurutkan seluruh buku berdasarkan rating secara descending menggunakan insertion sort, menampilkan 5 buku teratas setelah pengurutan, dan melakukan pencarian binary search untuk menemukan buku dengan rating tertentu yang dimasukkan pengguna.]

