package main

import (
	"fmt"
	"strconv"
)

func minimum(seq string) string {
	// result menyimpan string output
	result := ""

	// buat slice kosong dari bilanganß bulat untuk mensimulasikan stack (first in last out)
	var stack []int

	for i := 0; i <= len(seq); i++ {
		// tambahkan nomor i+1 ke dalam stack
		// index dimulai dari 0
		// karena aturannya angka yg ditampilkan hanya boleh 1-9 (index + 1)
		stack = append(stack, i+1)

		// jika i == panjang karakter (karakter terkahir), atau karakter saat ini adalah 'N' / naik
		if i == len(seq) || seq[i] == 'N' {
			// jalankan sampai stack kosong
			for len(stack) > 0 {
				// hapus elemen terakhir dari stack dan
				// tambahkan ke solusi
				result += strconv.Itoa(stack[len(stack)-1])
				result += " "
				stack = stack[:len(stack)-1]
			}
		}
	}

	return result
}

func main() {
	in := "NNMMM"
	fmt.Println(minimum(in))
}

/*

contoh karakternya = NNMMM
harapan outputnya = 126543

perulangan 1:
	karna karakter 1 adalah N, maka kita push angka 1 ke stack (i + 1) = 0 + 1 = 1
	stack = [1]

	krna dia N, kita pop stack
	output = 1

	stack kita jadi kosong
perulangan 2:
	karna karakter 2 adalah N, maka kita push angka 2 ke stack (i + 1) = 1 + 1 = 2
	stack = [2]

	krna dia N, kita pop stack
	output = 1 2

	stack kita jadi kosong
perulangan 3:
	karna karakter 3 adalah M, maka kita push angka 3 ke stack (i + 1) = 2 + 1 = 3
	stack = [3]
	krna dia M, maka percabangan tidak dijalankan
	output tetap = 1 2

perulangan 4:
	karna karakter 4 adalah M, maka kita push angka 4 ke stack (i + 1) = 3 + 1 = 4
	stack = [3, 4]
	krna dia M, maka percabangan tidak dijalankan
	output tetap = 1 2

perulangan 5:
	karna karakter 5 adalah M, maka kita push angka 5 ke stack (i + 1) = 4 + 1 = 5
	stack = [3, 4, 5]

	krna index sekarang sudah mencapai panjang karakter, maka kita pop stack sampai stack kosong
	output = 1 2 5 4 3
*/
