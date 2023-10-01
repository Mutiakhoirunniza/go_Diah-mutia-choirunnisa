>> Definisi software testing adalah sebuah proses untuk menganalisa sebuah software atau sistem untuk mendeteksi perbedaan antara kondisi fitur sistem saat ini dan fitur yang di inginkan stakeholder (user atau produck owner )
>> Tujuan software testing adalah untuk mengetahui apakah sistem yang kita buat sudah berjalan baik dan tidak muncul bug atau kerusakan sistem 

# kenapa kita harus melakukan testing yaitu, 
1. preventing Regression
(memastikan bahwa perubahan baru yang diterapkan dalam perangkat lunak tidak merusak fungsionalitas yang sudah ada dan tidak memunculkan bug baru. )

2. Confidence Level in refactoring 
(proses untuk mengubah struktur atau desain kode tanpa mengubah fungsionalitas eksternalnya. )

3. improve code design  
(memastikan bahwa perangkat lunak dapat berfungsi dengan baik, mudah dimengerti, dikelola, dan dikembangkan lebih lanjut.)

4. code documentation 
(proses atau praktik dalam pengembangan perangkat lunak yang melibatkan penulisan dokumentasi atau catatan yang menjelaskan kode komputer yang telah ditulis)

5. scaling the team
((peningkatan skala tim) dalam konteks pengembangan perangkat lunak merujuk pada tindakan untuk mengubah ukuran atau kapasitas tim pengembangan untuk mengakomodasi tuntutan yang lebih besar dalam proyek perangkat lunak.)


# Level Testing 
 1. UI TEST  
 2. INTEGRATION TEST 
 3. UNIT TEST 


# framework go 

(framework testtify)
framework unit testing yaitu sebuah kerangka kerja atau library yang digunakan dalam unit testing 
link  https://en.wikipedia.org/wiki/List_of_unit_testing_frameworks#Go

# structure 
structure dalam unit testing adalah strategi yang kita implementasikan dalam menaruh ile-file pengujian unit (unit testing) ke dalam proyek perangkat lunak yang lebih besar.

2 pattern dalam implementasi unit testing 
1. Centralized Unit Testing :  pengujian unit ditempatkan dalam satu lokasi yang mudah diakses  
kelebihannya : direktori dari apk kita lebih CLEAN 
[
/project
   /src          # Kode produksi
   /tests        # Semua file pengujian unit terpusat di sini
]

2. Test Together    : setiap modul atau komponen kode produksi memiliki direktori atau file pengujian unit yang berdampingan. 
kelebihannya : kita lebih mudah mencari file test dari function- function dari komponen 

[
    /project
   /src
      /module1
         module1.py            # Kode produksi untuk modul 1
         module1_test.py       # Pengujian unit untuk modul 1
      /module2
         module2.py            # Kode produksi untuk modul 2
         module2_test.py       # Pengujian unit untuk modul 2

]
 # RUNNER 
 >> runner adalah Sebuah aplikasi yang didesain untuk menjalankan test tersebut 
 tool atau fitur yang terdapat di runner 
 1. Watch mode adalah jika kita melakukan perubahan fitur tersebut akan melakukan test secara otomatis 
 2. choose  a runner  that  can  run fastest (mencari runner yang secara cepat.)

 # mocking 
mocking adalah sebuah Kebutuhan Anda untuk membuat objek tiruan (dan) objek palsu yang mensimulasikan perilaku objek nyata

# coverage 
coverage  adalah Pembuat kode perlu memastikan apakah mereka telah membuat pengujian yang cukup. Alat cakupan menganalisis kode aplikasi saat pengujian sedang berjalan.
 
>> coverage  report 
1. function 
2. Statement 
3. branch 
4. Lines 

>> Report Format
1. CLI 
2. XML 
3. HTML 
4. LCOV 

# simple steps to  test 
1. buat file pengujian baru
    ( Kriteria 
     1. Nama berakhiran library_test.go
     2. menaruh filenya di same folder, same package 
        same folder, different package 
    )

2. Tulis fungsi pengujian dengan Nama : Tes & Kata dengan huruf kapital

# running test 
go test ./lib/calculate -cover 

# running implementasi report 
go test ./lib/calculate -coverprofil=cover.out && go tool cover -html=cover.out
 
# test scenario VS test case 
test scenario membicarakan apa yang mau kita test  jika test case bagaimana cara kita mengetestnya 
# testing  scenario 
testing  scenario  adalah gambaran umum terhadap apa yang kita mau test atau bisa kita sebut high level test case
# test  case 
test  case  adalah kumpulan langkah langkah uji positive dan negative dalam sebuah test scenario.
test case ini meliputi : 
1. pre-conditions 
2. steps 
3. expected result 
4. status 
5. actual result 
 