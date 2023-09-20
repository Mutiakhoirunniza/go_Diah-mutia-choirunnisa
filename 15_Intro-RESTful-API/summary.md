API [application programming interface], yaitu sekumpulan aturan dan protokol yang memungkinkan berbagai perangkat lunak atau kompenen perangkat lunak berinteraksi satu sama lain.

API dapat digunakan untuk berbagai keperluan termasuk
1. integration
2. pengembangan Aplikasi 
3. aotomatisasi
4. pemograman backend
5. pemograman perangkat lunak 

cara kerja API 
{ CLIENT >>>>> request >>>> SERVER
SERVER >>>>> response >>>> CLIENT }

* REST[representational state transfer], yaitu merupakan suatu arsitektur berbasis standar yang digunakan dalam pengembangan aplikasi web dan layanan web 

* prinsip prinsip dasar REST 
1. RESOURCE 
2. REPRESENTASI 
3. STATELESS
4. INTERAKSI BERDASARKAN PESAN 
5. SISTEM TERDISTRIBUSI
6. UNIFORM INTERFACE 
7. LAYERED SYSTEM

** request dan response format 
1. JSON 
2. XML 
3. SOAP 

** HTTP request method 
 > GET, POST, PUT, DELETE, HEAD, OPTION, PATCH 

 . HTTP response code 
 200 = OK
 201 = CREATED 
 400 = BAD REQUEST 
 401 = UNAUTHORIZED 
 404 = NOT FOUND 
 405 = METHOD NOT ALLOWED 
 500 = INTERNAL SERVER ERROR 

 POSTMANT yaitu, pengembangan perangkat lunak untuk menguji, mengelola dan mengotomatisasi permintaan HTTP

 . package API di pemograman golang 
 NET/HTTP & ENCODING JSON 

 >  NET/HTTP = Membuat server API dan juga mengkonsum API 

 > fungsi fungsi untuk operasi JSON 
 - decode JSON to Object struct 
 - decode JSON to  map [string] interface{} & interface{}
 - decode  array JSON to array object 
 - encode object to JSON string 
 

