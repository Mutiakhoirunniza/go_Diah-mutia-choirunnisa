String, Advance Function , Pointer , Method, Struct and Interface 

String: tipe data yang mewakili kumpulan karakter. String diwakili oleh serangkaian byte yang mewakili karakter Unicode. String dapat dideklarasikan dengan menggunakan tanda kutip ganda ("") atau tanda kutip terbalik (`) untuk string multiline.

Advance Function: fungsi dapat memiliki karakteristik yang lebih maju, seperti kemampuan untuk mengembalikan beberapa nilai (multiple return values), fungsi sebagai parameter, dan menyimpan fungsi dalam variabel.

Variadic Function: fungsi yang dapat menerima jumlah argumen yang bervariasi. fungsi variadic dengan menambahkan tanda ellipsis (...) sebelum tipe parameter terakhir. Ini memungkinkan untuk memanggil fungsi dengan sejumlah argumen yang dapat berubah-ubah.

Anonymous Function:  function (fungsi anonim) adalah fungsi tanpa nama yang dapat didefinisikan dan digunakan secara langsung di dalam kode. Anonymous function sering digunakan untuk membuat closure atau sebagai parameter dalam pemanggilan fungsi lain.

Closure: konsep di mana sebuah fungsi dapat mengakses dan mengingat variabel di luar cakupan (scope) fungsinya sendiri.  fungsi closure memungkinkan untuk membuat fungsi yang dapat mengakses variabel dari cakupan luar, bahkan setelah cakupan luar tersebut sudah selesai dieksekusi.

Defer Function: Fungsi defer digunakan untuk menjadwalkan eksekusi suatu fungsi hingga selesai eksekusi blok yang mengandung pernyataan defer. Fungsi ini berguna untuk membersihkan sumber daya atau melakukan tindakan tertentu segera sebelum keluar dari fungsi atau blok.

Pointer: Pointer adalah tipe data khusus menyimpan alamat memori dari suatu nilai. 

Package: package adalah cara untuk mengorganisir kode menjadi unit-unit terpisah. Setiap file Go harus dimulai dengan deklarasi package, dan kode dari package lain dapat diimpor dan digunakan.

Error Handling: Error handling adalah praktik dalam pemrograman untuk mengatasi situasi ketika sesuatu tidak berjalan sesuai rencana.  error handling umumnya dilakukan dengan mengembalikan nilai error dari fungsi dan menggunakan mekanisme if statement atau package seperti errors untuk mengelola dan mengidentifikasi error.

sebuah interface adalah kontrak yang menguraikan metode-metode yang harus ada dalam objek yang mengimplementasikannya. Sebuah objek yang mengikuti kontrak ini dianggap "mematuhi" interface dan diharapkan memiliki metode-metode yang ditentukan dalam interface tersebut.


Deklarasi Interface:
Dalam blok deklarasi ini,  metode-metode yang perlu ada dalam objek yang mengimplementasikan interface ini, bersama dengan tipe pengembalian (return type) dari masing-masing metode.

Interface:
Kata kunci interface diikuti oleh nama interface yang ingin di buat. kemudian menuliskan definisi metode-metode yang diharapkan ada dalam objek yang mengimplementasikan interface ini.

Zero Value Interface:
membuat suatu tipe baru. Ini adalah tipe abstrak yang mewakili objek yang mematuhi kontrak interface tersebut. Zero value dari suatu interface adalah nil atau nilai nol. Saat Anda mendeklarasikan sebuah interface, jika tidak memberikan referensi konkret ke sebuah objek yang mengimplementasikan interface tersebut, maka nilainya akan menjadi nil.

Nil:
nil adalah nilai nol dalam Go yang mewakili ketiadaan nilai. Ketika mendeklarasikan sebuah interface tetapi tidak menginisialisasi dengan objek yang mengimplementasikan interface tersebut, maka nilainya akan menjadi nil.