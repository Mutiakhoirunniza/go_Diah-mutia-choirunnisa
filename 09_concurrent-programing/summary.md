>> concurrent programming >>

senquential program yaitu melakukan tugasnya secara ber urutan . 
paralleel program yaitu dapat melakukan tugas  secara bersamaan dengan syarat mempunyai multi-core diatas 1 core 
sedangkan concurrent program yaitu menjalankan tugas secara bersamaan dan independent dan juga tidak membutuhkan multi -core. akan tetapi nantinya akan terdapat pemecahan masalah dalam menjalankan tugas.
   jika kita mempunyai "concurrent dan multi - core"  nantinya tugas akan berjalan secara concurrency dengan parallelism atau menjalankan secara bersamaan dengan core yang sudah ada

>> goroutine adalah sebuah proses dimana proses concurrent untuk menjalankan dalam golang 
gorutine adalah sebuah fungsi yang dimana kita tentukan akan berjalan concurrently dibandingkan fungsi yang lain dan goroutine ini lebih ringan.
>> channel yaitu dapat di gunakan sebagai komunikasi 
>> select yaitu untuk mengontrol jalannya concurrent
perbedaan unbuffered channel dengan buffered, jika unbuffered channel yaitu penggunaan yang perlu goroutines tetapi jika buffered tidak perlu menggunakan goroutines, buffered dapat menggunakan data langsung bisa di buffered ke memory.
race condition yaitu salah satu kelemahan concurrent programming ketika kita memakai concurrent programming kita bisa mengakses file yang sama.
race condition adalah dimana goroutine mengakses variabel memory yang sama dimana salah satunya adalah writing.
race condition terjadi karena adanya unsynchronize ketika mengakses alat memori hyang di cari.
mutex adalah suatu alat yang digunakan untuk memastikan bahwa hanya 1 goroutine yang dapat mengakses sebuah bagian kode tertentu pada suatu waktu.
sehingga menghindari akses secara bersamaan yang dapat mengakibatkan masalah seperti data race
