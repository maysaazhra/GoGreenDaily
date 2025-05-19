KELOMPOK 2 :
1. Elok Elysia Dewi (103132400013)
2. Keyla Azzahra (103132400017)
3. Isna Farhah Umami (103132430004)
4. Maysa Azhra Fauziyyah (103132430005)


# 🌿 Aplikasi Pelacak Gaya Hidup Ramah Lingkungan

Aplikasi ini berbasis Go untuk mencatat dan mengelola aktivitas harian pengguna yang berdampak terhadap lingkungan, baik positif maupun negatif. Aplikasi ini membantu pengguna mengevaluasi gaya hidup mereka berdasarkan skor ramah lingkungan.

---

## Fitur Aplikasi

- ✅ Tambah aktivitas baru berdasarkan kategori
- 📄 Tampilkan daftar semua aktivitas
- ✏️ Edit nama, kategori, atau frekuensi aktivitas
- ❌ Hapus aktivitas tertentu
- 🔍 Pencarian aktivitas menggunakan:
  - Sequential Search
  - Binary Search (dengan pengurutan terlebih dahulu)
- 📊 Hitung total skor dari seluruh aktivitas
- 📑 Laporan bulanan lengkap dengan rekomendasi
- 🔢 Urutkan aktivitas berdasarkan nama atau skor 
  
---

## 📦 Struktur Data

### `Aktivitas`
```go
type Aktivitas struct {
    Nama      string
    Kategori  string
    Skor      int
    Frekuensi int
}
### `Kategori`
```go
type Kategori struct {
    Nama    string
    Positif bool
    Skor    int
}

---

## 💡 Daftar Kategori Aktivitas

| Kategori                        | Positif? | Skor |
|-------------------------------|----------|------|
| Hemat Energi                  | ✅       | 80   |
| Hemat Air                     | ✅       | 70   |
| Mengurangi Sampah             | ✅       | 95   |
| Transportasi Ramah Lingkungan | ✅       | 75   |
| Pola Makan Sehat              | ✅       | 85   |
| Boros Energi                  | ❌       | 20   |
| Buang Sampah Sembarangan      | ❌       | 5    |

---

## 📂 Menu Aplikasi

| No | Menu                                      | Fungsi                                                                 |
|----|-------------------------------------------|------------------------------------------------------------------------|
| 1  | **Tambah Aktivitas**                      | Menambahkan aktivitas ke daftar                                       |
| 2  | **Tampilkan Aktivitas**                   | Menampilkan seluruh aktivitas yang tercatat                           |
| 3  | **Update Aktivitas**                      | Mengubah nama, kategori, dan frekuensi aktivitas tertentu             |
| 4  | **Hapus Aktivitas**                       | Menghapus aktivitas berdasarkan nama                                  |
| 5  | **Cari Aktivitas (Sequential Search)**    | Mencari aktivitas berdasarkan nama menggunakan pencarian linear       |
| 6  | **Cari Aktivitas (Binary Search)**        | Mencari aktivitas setelah data diurutkan berdasarkan nama             |
| 7  | **Urutkan Berdasarkan Skor**              | Mengurutkan menggunakan algoritma *Selection Sort*                    |
| 8  | **Urutkan Berdasarkan Nama**              | Mengurutkan menggunakan algoritma *Insertion Sort*                    |
| 9  | **Hitung Total Skor**                     | Menjumlahkan seluruh skor dari aktivitas yang dicatat                 |
| 10 | **Cetak Laporan Bulanan**                 | Menampilkan ringkasan dan rekomendasi berdasarkan kebiasaan pengguna |
| 11 | **Keluar Aplikasi**                       | Menutup aplikasi                                                      |

---

## 🧠 Algoritma yang Digunakan

- 🔎 **Sequential Search** – Untuk mencari aktivitas berdasarkan nama secara linear.
- 🔍 **Binary Search** – Untuk pencarian cepat setelah data diurutkan berdasarkan nama.
- 📈 **Selection Sort** – Untuk mengurutkan daftar aktivitas berdasarkan skor.
- 📊 **Insertion Sort** – Untuk mengurutkan daftar aktivitas berdasarkan nama.

---

## 🖥️ Cara Menjalankan Aplikasi

### 1. Clone Repository
```bash
git clone https://github.com/maysaazhra/nama-repo.git
cd nama-repo
go run main.go

---

## 📊 Contoh Output Laporan Bulanan
=== Laporan Bulanan ===
Total aktivitas: 4
Aktivitas Positif: 3
Aktivitas Negatif: 1
Rata-rata Skor Aktivitas: 78

Rekomendasi:
- Gaya hidup Anda cukup ramah lingkungan. Pertahankan dan tingkatkan lagi!
