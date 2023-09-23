ORM (Object-Relational Mapping) adalah teknik dalam pengembangan perangkat lunak yang memungkinkan pengembang untuk menghubungkan antara objek dalam kode program dengan tabel dalam database relasional. Dalam konteks aplikasi yang menggunakan database relasional

>> ORM Advantages 
1. Lebih sedikit query berulang
2. Secara otomatis mengambil data ke objek siap pakai
3. Cara sederhana jika Anda ingin menyaring data
4. Sebelum menyimpannya di database
5. Beberapa memiliki fitur permintaan cache

Database migration adalah teknik yang digunakan dalam pengembangan perangkat lunak untuk mengelola perubahan struktur atau skema database secara terkontrol dan terdokumentasi. 

// manfaat DB migration
1. Kemudahan Update Database
2. Kemudahan Rollback Database
3. Lacak perubahan pada Struktur database
4. Riwayat struktur basis data ditulis pada sebuah kode
5. Selalu kompatibel dengan perubahan versi aplikasi

>>> Berikut adalah fitur utama dan konsep yang terkait dengan GORM:
1.  Model
2.  CRUD Operations
3. Auto-Migration
4. Query Builder
5. Transaction
6. Preload
7.  Hooks
8. Validasi 
9. Soft Deletes
10.  Concurrency Control

>> GORM adalah library ORM (Object-Relational Mapping) untuk bahasa pemrograman Go (Golang) yang digunakan untuk berinteraksi dengan basis data relasional. GORM menyediakan berbagai fitur untuk memodelkan hubungan antara tabel dalam basis data, termasuk relasi antara tabel.  Beberapa jenis hubungan yang dapat Anda definisikan dalam GORM antara lain: 
*belongs To
* Has one 
* Has many
* many to many 

MVC adalah singkatan dari Model-View-Controller, sebuah pola desain (design pattern) yang digunakan dalam pengembangan perangkat lunak untuk memisahkan komponen-komponen utama dalam aplikasi. 

>> Model: Model bertanggung jawab untuk mengakses dan memanipulasi data, serta memberikan antarmuka untuk berinteraksi dengan basis data jika diperlukan. Model juga dapat mengirim pembaruan ke komponen lain dalam aplikasi jika data berubah. 

>> View: View bertanggung jawab untuk tampilan atau antarmuka pengguna dalam aplikasi. View tidak melakukan pemrosesan data atau logika bisnis. 

>> Controller: Ini adalah bagian yang mengendalikan aliran informasi antara Model dan View. Ketika pengguna berinteraksi dengan aplikasi, Controller mengambil masukan tersebut, memprosesnya jika diperlukan, berinteraksi dengan Model untuk mendapatkan atau memperbarui data, dan kemudian memilih tampilan yang sesuai untuk menampilkan hasilnya kepada pengguna.
