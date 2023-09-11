join adalah sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel
>>> join standar sql ansi
1. iner join yaitu akan mengembalikanbaris baris dari dua tabel atau lebih yang memenuhi syarat 
     " SELECT t.message FROM users.u
     inner join tweets t ON u.id = t.user_id

2. left join yaitu mengembalikan seluruh baris dari tabel disebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join
    " SELECT U.ussername, t.message from users u left join tweetst ON u.id = t.user_id

3. union yaitu jumlah field yang dikeluarkan atau dipanggil harus sama 
    'SELECT ussername, fullname from users where id = 1
    union SELECT username, fullname from users where id = 2

>> aggregate yaitu fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai rinkasan tunggal 

>> fungsi agregasi SQL 
1. MIN = menampilkan sebuah nilai terkecil 
2. max = menampilkan sebuah nilai terbesar 
3. sum = menjumlahkan sebuah nilai 
4. avg = mencari nilai rata rata 
5. count = menjumlahkan banyaknya data 
6. having = menyelesaikan data 

subquery atau nested query yaitu query didalam query sql lain 
subquery digunakan untuk mengembalikkan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil
subquery digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil.
subquery dapat digunakan dengan insert, select, update, dan delete 
function adalah sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya

 