"echo" adalah sebuah library yang digunakan untuk membangun aplikasi web dengan mudah dan efisien. Echo adalah salah satu framework web yang populer dalam ekosistem Go yang membantu pengembang dalam membuat API web dan aplikasi web dengan cepat.

* beberapa fitur dan karakteristik utama dari library Echo di Golang:
1. Routing: Echo menyediakan sistem routing yang kuat dan mudah digunakan, yang memungkinkan Anda untuk menentukan rute dan mengaitkannya dengan penanganan (handler) yang sesuai untuk permintaan HTTP tertentu.
2. Middleware: Anda dapat menggunakan middleware untuk menambahkan fungsionalitas tambahan ke permintaan HTTP sebelum atau setelah penanganan utama dilakukan. Ini sangat membantu dalam hal otentikasi, autorisasi, log, dll.
3. Penanganan (Handlers): Echo memudahkan Anda dalam menulis penanganan untuk permintaan HTTP. Anda dapat dengan mudah mengikat fungsi atau metode ke rute tertentu.
4. Validasi dan Binding: Echo menyediakan fasilitas untuk mengikat data dari permintaan HTTP dan melakukan validasi data ini, yang sangat berguna saat mengelola input pengguna.
5. Mengelola Error: Anda dapat mengelola dan mengirimkan respons berbasis JSON untuk kesalahan dengan mudah menggunakan Echo.
6. Performa: Echo dikenal karena kinerja yang baik dan ringan, yang memungkinkan Anda untuk menangani lalu lintas web dengan cepat dan efisien.

* keunggulan utama dari Echo:
1. Kinerja yang Cepat
2. Ringan dan Efisien
3. Sintaksis yang Mudah Dimengerti
4. Routing yang Kuat
5. Middleware
6. Dokumentasi yang Baik


>>> Menginstall Echo Framework
go get github.com/labstack/echo/v4

>> Golang HTTP Routing Libraries - Benchmark
* Beego -> 1296.0ns
* Echo -> 75.7ns
* Gin -> 62,4ns
* Goji -> 646.0ns
* GorilaMux -> 2854.0ns
* HttpRouter -> 107.0ns

