package main

import (
	"fmt"
	"strconv"
	"strings"
)

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

//	func removeElement(nums []int, val int) int {
//		i := 0
//		for j := 0; j < len(nums); j++ {
//			if nums[j] == val {
//				i++
//			}
//		}
//		return len(nums) - i
//	}
func soal(input string) int {
	var result int

	for i, v := range input {
		angka, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("Input harus berupa angka")
			return 0
		}
		fmt.Println(angka)
		if angka%2 == 0 {
			result += angka + (i + 1)
		} else {
			result += angka * (i + 1)
		}
		fmt.Println(result)
	}
	return result
}

// tinggi awal = 200
// lama hari = 5
// pertumbuhan harian = 0.05

func solution(tinggiAwal int, lamaHari int, pertumbuhan float64) float64 {
	tinggiAkhir := float64(tinggiAwal)
	for i := 0; i < lamaHari; i++ {
		tinggiAkhir += tinggiAkhir * pertumbuhan
	}
	return tinggiAkhir
}

// tinggi = 3
// proses
// tinggi ke 1 =     *
// tinggi ke 2 =    /**\
// tinggi ke 3 =   /****\
func solution2(tinggi int) string {
	var result strings.Builder
	for i := 1; i <= tinggi; i++ {
		spasi := strings.Repeat(" ", tinggi+1-i)
		if i == 1 {
			result.WriteString(spasi + "*" + "\n")
		} else {
			bintang := strings.Repeat("*", 2*i-2) // Jumlah bintang bertambah sesuai baris
			result.WriteString(spasi + "/" + bintang + "\\" + "\n")
		}
	}
	return result.String()
}

// kata1 = imagination
// output = imagnto

// kata2 = association
// output = asocitn
func solution3(kata string) string {
	terlihat := make(map[rune]bool)
	var result strings.Builder

	for _, huruf := range kata {
		if !terlihat[huruf] {
			result.WriteRune(huruf)
			terlihat[huruf] = true
		}
	}

	return result.String()
}

// harga produk = [2000, 50000, 100000]
// output
// totalharga = 152.000
// diskon = 7.600
// bonus = topi
// hargaakhir = 144.400

func solution4(hargaProduk []int) (totalHarga int, diskon int, bonus string, hargaAkhir int) {
	for _, harga := range hargaProduk {
		totalHarga += harga
	}
	diskon = totalHarga * 5 / 100
	if totalHarga > 100000 {
		bonus = "topi"
	} else {
		bonus = "tidak ada bonus"
	}
	hargaAkhir = totalHarga - diskon

	return
}

// noakun = [111, 211, 201]
// nominal = [200000, 50000, 150000]

// proses
// akun1
// debit = 200000
// kredit = 0

// akun2
// debit = 200000
// kredit = 50000

// akun3
// debit = 200000
// kredit = 200000
// status = balance
func solution5(noAkun, nominal []int) (akun int, debit int, kredit int, status string) {
	// Membuat map untuk menyimpan saldo setiap akun
	saldo := make(map[int]int)

	// Memproses setiap akun dan nominal
	for i, akunNomor := range noAkun {
		if i < len(nominal) { // Pastikan indeks nominal valid
			saldo[akunNomor] += nominal[i] // Tambahkan nominal ke saldo akun
		}
	}

	// Menghitung total debit dan kredit
	for akunNomor, jumlah := range saldo {
		akun = akunNomor // Menyimpan akun terakhir yang diproses
		debit += jumlah  // Semua nominal dianggap sebagai debit
		kredit = jumlah  // Untuk contoh ini, kita set kredit sama dengan jumlah nominal
	}

	// Menentukan status saldo
	if debit == kredit {
		status = "balance"
	} else {
		status = "not balance"
	}

	return
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	newLength := removeElement(nums, val)
	result := nums[:newLength] // Membuat slice hasil yang valid
	fmt.Println(result)        // Output yang diharapkan: [0, 1, 3, 0, 4]

	// hasil := soal("236")
	// fmt.Println(hasil)
	// tinggiAwal := 200
	// lamaHari := 5
	// pertumbuhan := 0.05

	// hasil := solution(tinggiAwal, lamaHari, pertumbuhan)
	// fmt.Printf("Tinggi akhir setelah %d hari adalah: %.2f\n", lamaHari, hasil)

	// tinggiAwal2 := 250
	// lamaHari2 := 2
	// pertumbuhan2 := 0.02

	// hasil2 := solution(tinggiAwal2, lamaHari2, pertumbuhan2)
	// fmt.Printf("Tinggi akhir setelah %d hari adalah: %.2f\n", lamaHari2, hasil2)

	// tinggi := 3
	// hasil := solution2(tinggi)
	// fmt.Println(hasil)

	// kata1 := "imagination"
	// kata2 := "association"

	// hasil1 := solution3(kata1)
	// hasil2 := solution3(kata2)

	// fmt.Println("Output kata1:", hasil1)
	// fmt.Println("Output kata2:", hasil2)

	// hargaProduk := []int{2000, 50000, 100000}

	// // Memanggil fungsi solution4
	// totalHarga, diskon, bonus, hargaAkhir := solution4(hargaProduk)

	// // Output hasil
	// fmt.Printf("Total Harga: %d\n", totalHarga)
	// fmt.Printf("Diskon: %d\n", diskon)
	// fmt.Printf("Bonus: %s\n", bonus)
	// fmt.Printf("Harga Akhir: %d\n", hargaAkhir)

	// noAkun := []int{111, 211, 201}
	// nominal := []int{200000, 50000, 150000}

	// // Memanggil fungsi solution5
	// akun, debit, kredit, status := solution5(noAkun, nominal)

	// // Output hasil
	// fmt.Printf("Akun: %d\n", akun)
	// fmt.Printf("Debit: %d\n", debit)
	// fmt.Printf("Kredit: %d\n", kredit)
	// fmt.Printf("Status: %s\n", status)

	// Ez
	// s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	// reverseString(s1)
	// fmt.Println(string(s1)) // Output: "olleh"

	// s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	// reverseString(s2)
	// fmt.Println(string(s2)) // Output: "hannaH"

	// nums1 := []int{2, 7, 11, 15}
	// target1 := 9
	// result1 := twoSum(nums1, target1)
	// fmt.Println(result1) // Output: [0, 1]

	// nums2 := []int{3, 2, 4}
	// target2 := 6
	// result2 := twoSum(nums2, target2)
	// fmt.Println(result2) // Output: [1, 2]

	// nums3 := []int{3, 3}
	// target3 := 6
	// result3 := twoSum(nums3, target3)
	// fmt.Println(result3) // Output: [0, 1]

	// Med
	// n1 := 3
	// result1 := generateParenthesis(n1)
	// fmt.Println(result1) // Output: ["((()))","(()())","(())()","()(())","()()()"]

	// n2 := 1
	// result2 := generateParenthesis(n2)
	// fmt.Println(result2) // Output: ["()"]

	// s1 := "coaching"
	// t1 := "coding"
	// fmt.Println(appendCharacters(s1, t1)) // Output: 4

	// s2 := "abcde"
	// t2 := "a"
	// fmt.Println(appendCharacters(s2, t2)) // Output: 0

	// s3 := "z"
	// t3 := "abcde"
	// fmt.Println(appendCharacters(s3, t3)) // Output: 5

	// Hard
	// satisfaction1 := []int{-1, -8, 0, 5, -9}
	// fmt.Println(maxSatisfaction(satisfaction1)) // Output: 14

	// satisfaction2 := []int{4, 3, 2}
	// fmt.Println(maxSatisfaction(satisfaction2)) // Output: 20

	// satisfaction3 := []int{-1, -4, -5}
	// fmt.Println(maxSatisfaction(satisfaction3)) // Output: 0

	// fmt.Println(minInsertions("zzazz"))    // Output: 0
	// fmt.Println(minInsertions("mbadm"))    // Output: 2
	// fmt.Println(minInsertions("leetcode")) // Output: 5

	// ===================
	// nums, target := []int{2, 7, 11, 15}, 9

	// fmt.Println(twoSum2(nums, target))
}
