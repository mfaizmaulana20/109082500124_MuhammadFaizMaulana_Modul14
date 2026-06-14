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