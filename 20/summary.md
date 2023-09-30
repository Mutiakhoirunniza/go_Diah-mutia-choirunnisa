# Clean Architecture
>> Clean Architecture adalah suatu pendekatan dalam pengembangan perangkat lunak yang dirancang untuk menghasilkan kode yang mudah dipahami, mudah diuji, dan mudah diubah. 

>> Clean Architecture biasa punya takeline "the separation of concerns using layers" 

>> mission statement
Architecture yang bagus yaitu Architecture yang bisa terpisahkan secara layer dan membuat sebuah modular yang bisa di scalabel(upgarade), di maintainable(perbaiki) dan dapat di test aplikasi tersebut.

# prinsip-prinsip 
1. Pemisahan Konsep     : Clean Architecture memisahkan konsep inti dari aplikasi dari detail implementasi teknis.
2. Lapisan              : Arsitektur ini terdiri dari beberapa lapisan yang disusun dalam hierarki.
            
            Lapisan-lapisan tersebut biasanya mencakup:

    1. Entities
    2. Use Cases (Interactors)
    3. Interface Adapters
    4. Frameworks and Drivers

 # Clean and Hexagonal

 >> Clean and Hexagonal  pola arsitektur yang digunakan dalam pengembangan perangkat lunak untuk menciptakan aplikasi yang mudah dipelihara, dapat diuji, dan dapat diskalakan. Di Golang, prinsip-prinsip ini dapat diterapkan untuk membangun kode yang kuat dan terorganisir.
 > Prinsip utama dari Hexagonal Architecture adalah bahwa aplikasi inti (core application) harus tetap tidak bergantung pada teknologi apa pun dan harus fokus pada logika bisnis. 
> Port dan adapter memungkinkan fleksibilitas dalam mengganti atau mengupgrade teknologi atau sumber daya eksternal tanpa memengaruhi aplikasi inti.
 # context 
 > context adalah suatu package yang membawa deadline, concellation signal atau value lain pada request
 > paket context digunakan untuk mengelola siklus hidup operasi dan menyebar deadlines, pembatalan, dan nilai-nilai yang berhubungan dengan permintaan melintasi batas API. Ini sangat berguna dalam pemrograman mikro dan pemrograman konkuren.
 # implementation 
 { var ctx = context.Background()
    ctx = context.WithValue(ctx, "Key", "Value")}
 1. membuat root context dengan fungsi background()
 fungsi Background akan mengembalikan root context dimana kita bisa memakainya untuk operasi lain.
 2. contextWithValue akan sangat sering kita lihat pada middleware, kita dapat mengirim data middleware (contoh : user_id dari token) ke handler menggunakan context

 # before Migrate 
 ada 3 opsi apabila kita mau melakukan migrasi arsitektur kode dari mvc ke clean code 
 1. pertahankan desain sekarang dengan memisahkan dependensi 
 2. pertahankan desainnya tetapi kita pindahkan kodenya kedalam suatu layer 
 3. ubah desainnya dan pisahkan dependensi 

 # struktur kode yang akan kita pakai 
 <Controller [berisi kode yang berhubungan langsung ke user (interface layer)] -->
 <repository [berisi kode yang berhubungan langsung ke database (interface layer)]
 < Usecase [berisi] bisnis logik yang di pakai >


 # write Unit Test 
    > Pengujian unit di Golang penting untuk memastikan keandalan dan kemudahan pemeliharaan kode Anda. Kerangka kerja pengujian bawaan Go membuatnya mudah untuk menulis dan menjalankan pengujian unit.
    
    /Langkah-langkah untuk menulis pengujian unit di Golang

1. Buat Berkas Pengujian    : Buat berkas pengujian terpisah untuk setiap paket yang ingin diuji. Berkas pengujian harus memiliki sufiks_test.go.
2. Impor Paket Pengujian    : Impor paket testing dalam berkas pengujian Anda.
3. Tulis Fungsi Pengujian   : Tulis fungsi yang dimulai dengan kata "Test" diikuti dengan nama fungsi dan menerima parameter *testing.T. Fungsi-fungsi ini akan berisi kasus uji Anda.
4. Gunakan Fungsi Pengujian : Dalam fungsi pengujian Anda, gunakan fungsi pengujian seperti t.Errorf atau t.Fatalf untuk melaporkan kesalahan atau kegagalan.
5. Jalankan Pengujian       : Gunakan perintah go test untuk menjalankan pengujian Anda. Golang akan secara otomatis menemukan dan menjalankan fungsi pengujian Anda.

 