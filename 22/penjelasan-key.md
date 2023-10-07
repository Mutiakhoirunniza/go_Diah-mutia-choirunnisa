### SSH (Secure Shell) adalah sebuah protokol jaringan yang digunakan untuk mengakses dan mengelola perangkat jarak  jauh secara aman. 

## Key (Kunci SSH)
    Key SSH adalah metode autentikasi yang lebih aman daripada menggunakan kata sandi, dan melibatkan sepasang kunci yaitu   : kunci publik dan kunci pribadi.
## Kunci publik   
    Kunci publik  : adalah kunci yang ditempatkan di server yang Anda ingin akses (VM GCP atau EC2). 
    Kunci ini bersifat publik, yang berarti tidak rahasia.
## Kunci pribadi 
   Kunci pribadi   : adalah kunci yang simpan di komputer lokal.


### Proses autentikasi dengan kunci SSH bekerja
    Ketika mencoba mengakses VM dengan SSH, server VM akan meminta untuk mengirimkan kunci publik.
    Jika kunci publik yang dikirimkan sesuai dengan kunci publik yang ada di server, maka akan diizinkan untuk masuk tanpa memerlukan kata sandi.

### Password (Kata Sandi)
    Penggunaan kata sandi adalah metode autentikasi yang lebih sederhana. dan dapat mengaktifkan otentikasi berbasis password untuk SSH dengan langkah-langkah berikut:
            1. Atur password untuk pengguna SSH yang ada di dalam VM Anda.
            2. Ubah konfigurasi SSH di VM untuk mengaktifkan otentikasi berbasis password.
            3.Restart layanan SSH di VM untuk menerapkan perubahan.
Kelemahan penggunaan kata sandi adalah bahwa mereka dapat lebih mudah diretas jika tidak cukup kuat atau jika terjadi serangan kata sandi yang kuat.
Jadi, jika ingin terhubung ke VM di GCP atau EC2 dengan cara yang lebih aman, disarankan untuk menggunakan autentikasi berbasis kunci SSH. Namun, penggunaan kata sandi mungkin masih diperlukan dalam beberapa situasi, seperti ketika belum menyiapkan kunci SSH atau dalam situasi darurat jika kehilangan akses ke kunci pribadi.