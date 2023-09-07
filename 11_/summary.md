Data base adalah sekumpulan data yang terorganisir 
schema database adalah sebuah struktur atau rancangan yang mendefinisikan bagaimana data akan disimpan, diatur dan dihubungkan dalam database.
database relationship;
1. ONE TO ONE , suatu jenis hubungan antara 2 tabel atau entitas dalam database degan contoh : tabel pegawai dan tabel kontak darurat. setiap pegawai mungkin memiliki satu dan hanya satu informasi kontak darurat dan informasi kontak drurat tidak terhubung dengan pegawai lainnya.
2. ONE TO MANY, jenis hubungan antara 2 tabel atau entitas dalam sebuah database dimana satu entitas dalam tabel pertama memiliki banyak entitas yag terkait dalam tabel ke dua, tetapi entitas dalam tabel ke dua hanya terhubung dengan satu entitas dalam tabel pertama, contohnya tabel dosen dan tabel mahasiswa.
3. MANY TO MANY, jenis hubungan antara dua tabel dalam data base dimana banyak entitas dalam tabel pertama dapat terhubung dengan banyak entitasbdalam tabel kedua. contoh : tabel mahasiswa dan tabel mata kuliah. " mahasiswa dapat mendaftar matakuliah dan banyak mata kuliah dapat diikuti oleh banyak mahasiswa " 
ERD : entity relationship diagram
RDBMS : sebuah aplikasi untuk memanajementkan sistem 
SQL : bahasa pemograman khusu yang digunakan untuk memanajemenkan data dalam RDBMS.
non SQL : jenis data base yang tidak struktural atau tidak mengikuti model data relasional yang digunakan oleh RDBMS.

jenis jenis SQL 
1. DDL (data definition laguage) suatu perintah yang digunakan untuk membuat, mengubah dan menghapus struktural 
>> CREATE DATABASE :    membuat database 
>> SHOW DATABASW :  menampilkan database 
>> USE DATABASE :  mengakses database 
>> DROP DATABASE : Menghapus tabel 
>> DESCRIBE : melihat struktur tabel 

modifikasi tabel ada dua yaitu " alter tabel dan add column"
 DML : suatu perintah untuk memanipulasi data dalam tabel dari suatu database 
 statement operation: 
 1. insert into 
 2. select * from ( * menampilkan semua )
 jika ingin mengganti nama tael bisa menggunakan 'AS' 

 join standar SQL 
 >> inner join akan mengembalikan baris baris dari dua tabel atau lebih yang memenuhi syarat
 >> left join yaitu akan mengembalikan seluruh baris dari tabel di sebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join 
>> right join yaitu mengembalikan semua baris dari tabel sebelah kanan yang di kenai kondisi ON dengan data dari tabel sebelah kiri yang memenuhi kondisi join. 


fungsi agregasi yaitu fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.
