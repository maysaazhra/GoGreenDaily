package main

import "fmt"

// Struct Aktivitas
type Aktivitas struct {
	Nama      string
	Kategori  string
	Skor      int
	Frekuensi int
}

// Struct Kategori
type Kategori struct {
	Nama    string
	Positif bool
	Skor    int
}

// Array untuk menyimpan data aktivitas
var Daftar_Aktivitas []Aktivitas

// Data Kategori Aktivitas
var DataKategori = []Kategori{
	{"hemat energi", true, 80},
	{"hemat air", true, 70},
	{"mengurangi sampah", true, 95},
	{"transportasi ramah lingkungan", true, 75},
	{"pola makan", true, 85},
	{"boros energi", false, 20},
	{"pencemaran lingkungan", false, 5},
}

// Fungsi tambah aktivitas
func tambah_aktivitas() {
	var nama string
	var pilihan, frekuensi int

	fmt.Print("Nama Aktivitas: ")
	fmt.Scanln(&nama)

	tampilkanKategori()
	fmt.Print("Masukkan nomor kategori: ")
	fmt.Scanln(&pilihan)

	if pilihan < 1 || pilihan > len(DataKategori) {
		fmt.Println("Kategori tidak valid.")
		return
	}

	fmt.Print("Jumlah aktivitas yang dilakukan dalam sebulan: ")
	fmt.Scanln(&frekuensi)

	kategoriDipilih := DataKategori[pilihan-1]
	aktivitas := Aktivitas{
		Nama:      nama,
		Kategori:  kategoriDipilih.Nama,
		Skor:      kategoriDipilih.Skor,
		Frekuensi: frekuensi,
	}

	Daftar_Aktivitas = append(Daftar_Aktivitas, aktivitas)
	fmt.Printf("Aktivitas berhasil ditambahkan dengan skor: %d\n", aktivitas.Skor)
}

// Fungsi tampilkan aktivitas
func tampilkan_aktivitas() {
	fmt.Println("=== Daftar Aktivitas ===")
	for i := 0; i < len(Daftar_Aktivitas); i++ {
		a := Daftar_Aktivitas[i]
		fmt.Printf("%d. Nama: %s | Kategori: %s | Skor: %d\n", i+1, a.Nama, a.Kategori, a.Skor)
	}
	fmt.Println()
}

// Fungsi update aktivitas
func edit_aktivitas() {
	var nama, namaBaru string
	var pilihan, frekuensiBaru int

	fmt.Print("Masukkan nama aktivitas yang ingin diedit: ")
	fmt.Scanln(&nama)

	index := -1
	for i := 0; i < len(Daftar_Aktivitas); i++ {
		if Daftar_Aktivitas[i].Nama == nama {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Aktivitas tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama aktivitas baru: ")
	fmt.Scanln(&namaBaru)

	tampilkanKategori()
	fmt.Print("Masukkan nomor kategori baru: ")
	fmt.Scanln(&pilihan)

	if pilihan < 1 || pilihan > len(DataKategori) {
		fmt.Println("Kategori tidak valid.")
		return
	}

	fmt.Print("Jumlah aktivitas yang dilakukan dalam sebulan: ")
	fmt.Scanln(&frekuensiBaru)

	kategoriDipilih := DataKategori[pilihan-1]

	Daftar_Aktivitas[index].Nama = namaBaru
	Daftar_Aktivitas[index].Kategori = kategoriDipilih.Nama
	Daftar_Aktivitas[index].Skor = kategoriDipilih.Skor
	Daftar_Aktivitas[index].Frekuensi = frekuensiBaru

	fmt.Println("Aktivitas berhasil diperbarui.")
}

// Fungsi hapus aktivitas
func hapus_aktivitas() {
	var Nama string
	fmt.Print("Masukkan Nama Aktivitas yang ingin dihapus: ")
	fmt.Scanln(&Nama)

	for i := 0; i < len(Daftar_Aktivitas); i++ {
		if Daftar_Aktivitas[i].Nama == Nama {
			Daftar_Aktivitas = append(Daftar_Aktivitas[:i], Daftar_Aktivitas[i+1:]...)
			fmt.Println("Aktivitas berhasil dihapus.")
			return
		}
	}
	fmt.Println("Aktivitas tidak ditemukan.")
}

// Fungsi cari aktivitas dengan data terurut
func SequentialSearch() {
	var Nama string
	fmt.Print("Masukkan Nama Aktivitas yang ingin dicari: ")
	fmt.Scanln(&Nama)

	for i := 0; i < len(Daftar_Aktivitas); i++ {
		a := Daftar_Aktivitas[i]
		if a.Nama == Nama {
			fmt.Printf("Aktivitas ditemukan:\nNama: %s\nKategori: %s\nSkor: %d\n\n", a.Nama, a.Kategori, a.Skor)
			return
		}
	}
	fmt.Println("Aktivitas tidak ditemukan.")
}

// Fungsi cari aktivitas dengan data tidak terurut
func BinarySearch() {
	var nama string
	fmt.Print("Masukkan Nama Aktivitas yang ingin dicari (Binary Search): ")
	fmt.Scanln(&nama)

	// Urutkan dulu sebelum binary search
	sort_aktivitas_by_nama()

	low := 0
	high := len(Daftar_Aktivitas) - 1

	for low <= high {
		mid := (low + high) / 2
		if Daftar_Aktivitas[mid].Nama == nama {
			a := Daftar_Aktivitas[mid]
			fmt.Printf("Aktivitas ditemukan:\nNama: %s\nKategori: %s\nSkor: %d\n\n", a.Nama, a.Kategori, a.Skor)
			return
		} else if Daftar_Aktivitas[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Aktivitas tidak ditemukan.")
}

// Fungsi Selection Sort berdasarkan skor
func selectionSort() {
	var opsi string

	fmt.Print("Urutkan skor secara descending? (yes or no): ")
	fmt.Scanln(&opsi)
	desc := (opsi == "yes" || opsi == "Yes")
	n := len(Daftar_Aktivitas)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if desc {
				if Daftar_Aktivitas[j].Skor > Daftar_Aktivitas[idx].Skor {
					idx = j
				}
			} else {
				if Daftar_Aktivitas[j].Skor < Daftar_Aktivitas[idx].Skor {
					idx = j
				}
			}
		}
		if idx != i {
			Daftar_Aktivitas[i], Daftar_Aktivitas[idx] = Daftar_Aktivitas[idx], Daftar_Aktivitas[i]
		}
	}
	fmt.Println("Aktivitas berhasil diurutkan berdasarkan skor.")
}

// Fungsi insertion sort berdasarkan nama
func InsertionSort() {
	n := len(Daftar_Aktivitas)
	for i := 1; i < n; i++ {
		key := Daftar_Aktivitas[i]
		j := i - 1

		for j >= 0 && Daftar_Aktivitas[j].Nama > key.Nama {
			Daftar_Aktivitas[j+1] = Daftar_Aktivitas[j]
			j--
		}
		Daftar_Aktivitas[j+1] = key
	}

	fmt.Println("Aktivitas berhasil diurutkan berdasarkan nama (A-Z).")
}

// Fungsi untuk menampilkan laporan bulanan
func laporan_bulanan() {
	totalSkor := 0
	jumlahPositif := 0
	jumlahNegatif := 0

	for i := 0; i < len(Daftar_Aktivitas); i++ {
		a := Daftar_Aktivitas[i]
		totalSkor += a.Skor

		for j := 0; j < len(DataKategori); j++ {
			if a.Kategori == DataKategori[j].Nama {
				if DataKategori[j].Positif {
					jumlahPositif++
				} else {
					jumlahNegatif++
				}
				break
			}
		}
	}

	totalAktivitas := len(Daftar_Aktivitas)
	rataRata := 0
	if totalAktivitas > 0 {
		rataRata = totalSkor / totalAktivitas
	}

	fmt.Println("=== Laporan Bulanan ===")
	fmt.Printf("Total aktivitas: %d\n", totalAktivitas)
	fmt.Printf("Aktivitas Positif: %d\n", jumlahPositif)
	fmt.Printf("Aktivitas Negatif: %d\n", jumlahNegatif)
	fmt.Printf("Rata-rata Skor Aktivitas: %d\n", rataRata)

	fmt.Println("\nRekomendasi:")
	if jumlahNegatif > jumlahPositif {
		fmt.Println("- Kurangi aktivitas negatif, seperti membuang sampah sembarangan atau boros energi.")
	}
	if totalAktivitas == 0 {
		fmt.Println("- Tambahkan aktivitas ramah lingkungan untuk memulai gaya hidup hijau.")
	} else if rataRata < 60 {
		fmt.Println("- Tingkatkan frekuensi atau jumlah aktivitas dengan skor tinggi, seperti tanam pohon atau mengurangi sampah.")
	} else {
		fmt.Println("- Gaya hidup Anda cukup ramah lingkungan. Pertahankan dan tingkatkan lagi!")
	}
	fmt.Println()
}

func main() { //menampilkan menu utama dalam bentuk looping dan memanggil fungsi-fungsi sesuai menu
	var pilihan int //untuk menyimpan input user (menu yang dipilih).

	for {
		fmt.Println("=== Aplikasi Pelacak Gaya Hidup Ramah Lingkungan ===")
		fmt.Println("1. Tambah aktivitas")
		fmt.Println("2. Tampilkan aktivitas")
		fmt.Println("3. Update aktivitas")
		fmt.Println("4. Hapus aktivitas")
		fmt.Println("5. Cari aktivitas (SequentialSearch)")
		fmt.Println("6. Cari aktivitas (BinarySearch)")
		fmt.Println("7. Urutkan Aktivitas berdasarkan Skor (Selection Sort)")
		fmt.Println("8. Urutkan Aktivitas berdasarkan Nama (Insertion Sort)")
		fmt.Println("9. Hitung total skor")
		fmt.Println("10. Cetak laporan bulanan")
		fmt.Println("11. Keluar aplikasi")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan { //digunakan untuk mengevaluasi nilai dari variable pilihan (biasanya input dari pengguna)
		case 1:
			tambah_aktivitas()
		case 2:
			tampilkan_aktivitas()
		case 3:
			edit_aktivitas()
		case 4:
			hapus_aktivitas()
		case 5:
			SequentialSearch()
		case 6:
			BinarySearch()
		case 7:
			selectionSort()
		case 8:
			insertionSort()
		case 9:
			hitung_total_skor()
		case 10:
			laporan_bulanan()
		case 11:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

// Fungsi untuk mengurutkan data nama aktivitas dari terurut
func sort_aktivitas_by_nama() {
	for i := 1; i < len(Daftar_Aktivitas); i++ {
		key := Daftar_Aktivitas[i]
		j := i - 1
		for j >= 0 && Daftar_Aktivitas[j].Nama > key.Nama {
			Daftar_Aktivitas[j+1] = Daftar_Aktivitas[j]
			j--
		}
		Daftar_Aktivitas[j+1] = key
	}
}

// Fungsi menampilkan kategori
func tampilkanKategori() {
	fmt.Println("Pilih kategori aktivitas:")
	for i := 0; i < len(DataKategori); i++ {
		fmt.Printf("%d. %s\n", i+1, DataKategori[i].Nama)
	}
}

// Fungsi hitung total skor
func hitung_total_skor() {
	total := 0
	for i := 0; i < len(Daftar_Aktivitas); i++ {
		total += Daftar_Aktivitas[i].Skor
	}
	fmt.Printf("Total skor dari semua aktivitas: %d\n\n", total)
}

// Fungsi hitung skor otomatis
func hitungSkorOtomatis(kategori string) int {
	if kategori == "hemat energi" {
		return 80
	} else if kategori == "hemat air" {
		return 70
	} else if kategori == "mengurangi sampah" {
		return 95
	} else if kategori == "transportasi ramah lingkungan" {
		return 75
	} else if kategori == "pola makan" {
		return 99
	} else if kategori == "boros energi" {
		return 20
	} else if kategori == "pencemaran lingkungan" {
		return 5
	}
	return 50
}
