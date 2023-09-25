Middleware adalah entitas yang terhubung ke pemrosesan permintaan/respons server. Middleware, izinkan kami menerapkan berbagai fungsi seputar Permintaan HTTP masuk ke server Anda dan respons keluar.

implementasi Middleware  bagian klien 

loginMiddleware     : untuk mencatat lock apa saja yang keluar masuk ke server
sesionMiddleware    : untuk mencatat apakah klien masih aktif 
AuthMiddleware      : untuk authentication apakah sesi klien berhak mengakses alamat url
customMiddleware    : untuk membuat sebuah custom sebelum melakukan komunikasi data antara sesi klien dan  server

>> tujuan Auth Middleware  
1. keamanan data
2. dapat melihat sebagai admin atau user biasa 
 >> authentication Middleware 
 
 || Information Encoded Format :
    'Authorization  : Basic ' + base64encode('username:password')
    Header Generate : 
    Authorization   : Basic dXN1cm5hbWU6cGFzc3dvcmQ=


 . Basic authentication  = sebuah proses authentication rest API dengan mengirimkan  data username dan password melalui data Headers. 
 >> Aturan utama untuk menggunakan Basic Authentication 
 Format Header: Data username dan password dikirim dalam header HTTP dengan menggunakan format Authorization: Basic base64(username:password)

 . JSON  Web Token = sebuah standar terbuka yang mendefinisikan cara yang aman untuk mentransmisikan informasi antara dua pihak dalam bentuk objek JSON. JWT digunakan untuk mengirim data yang bersifat terverifikasi dan terenkripsi antara server dan klien, JWT sering digunakan dalam aplikasi web modern untuk otentikasi dan otorisasi, serta pertukaran informasi yang aman.
  
  >> JWT memiliki tiga bagian utama yang dienkripsi menjadi satu token(keynya barer)
  1. Header  = algorithma dan token 
  2. Payload  = Data 
  3. Verify Signature 

. Cara kerja umum JWT :

1. Sebuah klien mengirim permintaan autentikasi ke server.
2. Setelah server memverifikasi kredensial pengguna, ia menghasilkan JWT yang berisi klaim-klaim yang relevan.
3. Server mengirim JWT ke klien sebagai respons.
4. Klien menyimpan JWT tersebut dan menggunakannya dalam setiap permintaan berikutnya ke server.
5. Server memeriksa integritas JWT yang diterima dari klien, serta mengizinkan atau menolak permintaan berdasarkan informasi dalam JWT.


|| JWT Echo JWT middleware is located at https://github.com/labstack/echo-jwt

>> Example Third Party Middleware
1. Negroni
2. Echo
3. Interpose
4. Alice

>> Middleware in Echo
1. Basic Auth = Usage 
e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))

2. Body Dump
3. Body Limit
4. CORS
5. CSRF
6. Casbin Auth
7. Gzip
8. JWT = Usage 
e.Use(echojwt.JWT([]byte("secret")))
9. Key Auth
10. Logger = Usage 
e.Use(middleware.Logger())
11. Method Override
12. Proxy
13. Recover
14. Redirect
15. Request ID
16. Rewrite
17. Secure
18. Session
19. Static
20. Trailing Slash


>> setup middleware echo 
 ECHO #Pre() = jenis middleware yang dieksekusi sebelum penanganan permintaan utama (handler).
 ECHO #Use() =  jenis middleware yang dieksekusi sebelum dan setelah penanganan permintaan utama (handler).
 
>> Logger Middleware 
Middleware Logger digunakan untuk mencatat detail permintaan HTTP, seperti metode HTTP, path, kode status, waktu respons, dan informasi lainnya.

