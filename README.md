Berikut adalah penjelasan yang telah disesuaikan untuk ditambahkan ke file README di GitHub agar lebih rapi dan terstruktur:

````markdown
# go-learn-basic

## 1. Inisialisasi Proyek

1.1 Gunakan perintah berikut untuk menginisialisasi proyek saat mulai menggunakan package:

```bash
go mod init project
```
````

1.2 Untuk menyesuaikan package yang digunakan, jalankan:

```bash
go mod tidy
```

## 2. Variabel Golang

Variabel di Golang digunakan untuk menyimpan data yang dapat digunakan dan dimanipulasi dalam program. Berikut adalah beberapa poin penting mengenai variabel di Golang:

Tentu! Berikut adalah penjelasan lebih mendetail tentang variabel di Golang yang bisa kamu masukkan ke README:

````markdown
## 2. Variabel Golang

Variabel di Golang digunakan untuk menyimpan data yang dapat digunakan dan dimanipulasi dalam program. Berikut adalah beberapa poin penting mengenai variabel di Golang:

### 2.1. Deklarasi Variabel

- Variabel dapat dideklarasikan menggunakan kata kunci `var` diikuti dengan nama variabel dan tipe data.

  ```go
  var name string
  var age int
  ```
````

- Variabel juga dapat dideklarasikan dan diinisialisasi dalam satu baris:

  ```go
  var name string = "Reza"
  var age int = 20
  ```

- Dengan menggunakan inferensi tipe, kamu bisa menghilangkan tipe data dan membiarkan Go menentukan tipe secara otomatis:

  ```go
  name := "Reza"
  age := 20
  ```

### 2.2 Tipe Data

Golang mendukung beberapa tipe data, antara lain:

- **Tipe Data Dasar**:

  - `int`: Bilangan bulat
  - `float64`: Bilangan desimal
  - `string`: Teks
  - `bool`: Nilai benar/salah

- **Tipe Data Komposit**:
  - **Array**: Kumpulan elemen dengan tipe yang sama.
  - **Slice**: Versi dinamis dari array.
  - **Map**: Kumpulan pasangan kunci-nilai.
  - **Struct**: Tipe data yang dapat menyimpan beberapa nilai dengan tipe yang berbeda.

### 2.3 Aksesibilitas Variabel

- Variabel yang diawali dengan huruf besar dapat diakses dari package lain (public).
- Variabel yang diawali dengan huruf kecil hanya dapat diakses dalam package yang sama (private).

### 2.4 Mengubah Nilai Variabel

Setelah variabel dideklarasikan, kamu dapat mengubah nilainya:

```go
age = 21
```

### 2.5. Contoh Penggunaan Variabel

Berikut adalah contoh penggunaan variabel dalam program sederhana:

```go
package main

import "fmt"

func main() {
    var name string = "Reza"
    var age int = 20

    fmt.Println("Nama:", name)
    fmt.Println("Usia:", age)

    // Mengubah nilai variabel
    age = 21
    fmt.Println("Usia setelah setahun:", age)
}
```

### 2.6. Konstanta

Selain variabel, Golang juga mendukung konstanta yang dideklarasikan dengan kata kunci `const`. Konstanta tidak dapat diubah setelah dideklarasikan.

```go
const Pi = 3.14
```





## 3. Membuka dan Membaca File
- **readFile**:
  - Jika ukuran file kecil, gunakan ReadFile
  - Jika ukuran file besar, cetak baris per baris. gunakan `scanner.Scan()`.

## 4. Penulisan Package
- Penulisan nama package dengan huruf besar dan kecil sangat berpengaruh terhadap level akses:
  - **Huruf Besar**: Dapat dipanggil di package lain (public).
  - **Huruf Kecil**: Hanya dapat dipanggil di package yang sama (private).
```

### Catatan

- Pastikan untuk menambahkan penjelasan lebih lanjut di setiap bagian jika diperlukan, agar pembaca dapat memahami konteks dan penggunaan masing-masing fitur dengan lebih baik.
- Gunakan format markdown yang sesuai untuk menjaga keterbacaan dan struktur yang baik di GitHub.

Jika ada yang ingin ditambahkan atau diubah, silakan beri tahu!
