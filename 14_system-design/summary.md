diagram design, yaitu sebuah simbol yang mempresentasikan informasi apa yang ingin disampaikan. 
dalam membuat diagram design kita menggunakan tools untuk mempermudah dalam pembuatan diagram diagram tersebut.

jenis jenis system design 
1. flowchart , sebuah representasikan dari sebuah diagraman yang mengvisualisasikan dari sebuah alur.
2. use case diagram, yaitu menggambarkan fungsionalitas sistem dari perspektif pengguna atau aktor yang berinteraksi dengan sistem tersebut 
3. ERD, sebuah entity untuk menggambarkan struktur data dan hubungan antar entitas dalam sistem atau organisasi.
4. HLA (high level architecture) yaitu konsep dasar yang digunakan dalam pengembangan sistem, perangkat lunak atau infrastruktur teknologi informasi untuk menggambarkan struktur keseluruhan dari suatu sistem atau aplikasi secara abstrak 

* horizontal scaling yaitu bertujuan untuk meningkatkan daya tampung atau kinerja dengan menambahkan lebih banyak sumber daya komputasi dalam bentuk mesin fisik atau node komputasi tambahan 
* vertical scaling untuk meningkatkan daya tampung atau kinerja dengan meningkatkan sumber daya pada satu mesin fisik atau node komputasi tunggal 

>> karakteristik system distribusi 
1. scalability 
2. reliability 
3. availability 
4. efficiency 
5. serviceability or manageability 

* job/work Queue; mengatur da mengelola tugas tugas atau pekerjaan yang perlu dieksekusi dalam satu aplikasi atau sistem.
* load blancing; praktik distribusi lalu lintas jaringan atau beban kerja komputasi secara merata diantara beberapa sumber daya atau server untuk meningkatkan kinerja, efisiensi dan ketersediaan sistem.
* monolithic ; seluruh aplikasi atau sistem dikembangkan sebagai satu entitas tunggal atau modul  besar yang berisi semua fungsi dan komponen yang d perlukan
* microservices; sebuah aplikasi besar dibagi menjadi sejumlah kecil layanan mandiri yang beroperasi secara independent

* SQL yaitu sebuah relational database yang berhubungan dengan sistem manajement basis data (DBMS)
* non SQL yaitu database yang tidak strutur dan bersifat dinamis

prinsip relational DB 
1. atomicity = transaksi terjadi semua atau tidak sama sekali 
2. consistency = data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu 
3. isolation = pada saat terjadi request yang bersamaan (consurrent) memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial
4. durability = jaminan bahwa transaksi yang telah tersimpan tetap tersimpan 


>> schema -less 

> traditional RDBMS                                      
tidak bisa menambahkan data yang tidak sesuai skema
memiliki  tipe data yang strict
tidak dapat menambahkan beberapa item data pada sebuah field

 > no SQL DBMS
tidak memiliki skema ketika menambahkan data 
apk menangani proses validasi tipe data 
mendukung proses aggregasi dokumen pada item 

>> caching ; penyimpanan data sementara yang sering diakses dilokasi yang memungkinkan akses lebih cepat, mengurangi kebutuhan untuk mengambil data dari sumber aslinya 
>> ram yaitu penyimpanan data sementara dan sedangkan database adalah penyimpanan data yang permanent

. database replication yaitu proses membuat dan menjaga salinan data dari sebuah basis data kedalam satu atau lebih lokasi lainnya.
> redundancy yaitu suatu pendekatan dalam data atau sumber daya yang disalin atau disimpan secara berlebihan untuk mengantisipasi kegagalan atau kerusakan 
> replication, proses pembuatan da menjaga salinan data atau sumber daya yang sama dibebrapa lokasi 
> database indexing yaitu proses pengoptimalan dalam manajemen basis data dimana struktur data tambahan disiapkan untuk mempercepat operasi pencarian atau pengurutan dan akses data dalam sebuah tabel atau basis data. 