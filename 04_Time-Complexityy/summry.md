complexity ialah bagaimana kita menghitung sebuah algoritma atau sebuah sistem itu akan berjalan dalam hal waktu.
cara menghitung waktu tersebut yaitu, dengan cara menen9tuka beberapa banyak operasi dominant yang dilakukan contohnya operasi yang primitif ( perkalian, pertambahan serta slice )

O(n) {Dominant} -----> operasi yang dominant yaitu operasi yang mengikuti banyaknya  n, jika n adalah 10 maka operasi tersebut akan dilakukan sebanyak iterasi 10 kali 
O(1) {Constant time} ----> operasi yang menjalankan hanyak 1 kali dan tidak berpengaruh dengan inputan atau n
O(n) {Linear time} ----> operasi yang dilakukan sebanyak n, mempunyai slice dan index
slice yaitu []
index yaitu i

O(n+m) {Linear time} ----> operasi yang dilakukan atau di kerjakan secara 2x, operasi pertama dan terakhir. yang pertama O(n) lalu O(m) setelah itu di jumlahkan. 
O(log n) {Logarithmic time} ----> operasi yang dilakukan dengan memecahkan perkelompokkan tersebut. atau /2
O(n ^ 2) {Quadratic time} ---- > melakukan perulanga pertama yaitu dimulai 0 sampai n dan dilakukan iterai sebanyak n tersebut. 


contoh N = 8 

O(n) = ia di lakukan sebanyak 8 x operasi 
O(1) = operasinya hanyak dilakukan sebanyak 1 kali 
O (log n) Operasinya dilakukan /2.

space complexity( yaitu tempat penyimpanan sementara dalam operasi)