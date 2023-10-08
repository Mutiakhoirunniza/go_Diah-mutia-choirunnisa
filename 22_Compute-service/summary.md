### Introduction Deployment
## system & Software Deployment 
deployment adala kegiatan yang bertujuan untuk menyebarkan aplikasi atau produk yang telah di kerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanent.
## strategi  Deployment 
    1. Big-Bang deployment strategy atau sering di sebut Replace deployment strategi 

kelebihannya : 
            . Mudah di implementasikan, cara klasik, tinggal replace 
            . perubahan kepada sistem langsung 100% secara instan 

kekurangan : 
            . terlalu beresiko, rata rata downtime cukup lama 

    2. Rollout Deployment Strategi 

Kelebihan : 
        . Lebih aman dan less downtime dari versi sebelumnya
kekurangan :
         . Akan ada 2 versi aplikasi berjalan secara barengan sampai semua server terdeploy, dan bisa membuat bingung
        . Karena sifatnya perlahan satu persatu, untuk deployment dan rollback lebih lama dari yang Bigbang, karena prosesnya perlahan lahan sampai semua server terkena efeknya
        . Tidak ada kontrol request, Server yang baru ter-deploy dengan aplikasi versi bary, langsung mendapatkan request yang sama banyaknya dengan server lain.

    3. Blue/Green deployment strategi 

Kelebihan :
       . Perubahan sangat cepat, sekali switch service langsung berubah 100%
       . Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment
Kekurangan :

. Resource yang dibutuhkan lebih banyak, karena untuk setiap deployment kita harus menyediakan service yang berupa environtmenta dengan yang sedang berjalan di production
. Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak.

    4. cannary Deployment strategi 

Kelebihan :
        . Cukup aman
        . Mudah untuk rollback jika terjadi error/bug, tanpa berimbas kesemua user
Kekurangan :

        . Untuk mencapai 100% cukup lama dibanding dengan Blue/Green deployment. Dengan Blue/Green deployment, aplikasi langsung 100% terdeploy keseluruh user.

### simple deployment process
1. Development Env. (localhost dev config)
2. Test Env. (Test server Test config)
3. UAT Env. (User Acceptance UAT Config)
4. production Env.(Production server production config)
### catatan 
masuk database 
docker run --name crud_go_test -e MYSQL_ROOT_PASSWORD=123 -d mysql:latest
docker exec -it crud_go_test bash
## isi sblm masuk 
mysql -u root -p
## buatt database : 
create database crud_go_test;

exit 
exit 

## masuk docker run : 
docker run -p 8000:8000 compute-latest:latest

## build docker : 
docker build -t compute-latest .

## docker tag : 
docker tag compute-latest:latest mutiakhoirunniza/mut

## docker push : 
docker push mutiakhoirunniza/mut

## catatan ssh 
ssh -i ~/.ssh/id_ecdsa mutia@


## Docker Build : 
docker build -t <nama-image>:<tag> .

## docker run : 
docker run -p 8080:8080 --name <container-name> <image-name>:<tag>
## * with env
docker run -d -p 80:80 -e DBUSER=root -e DBPASS=123 -e DBHOST=127.0.0.1 -e DBPORT=3306 -e DBNAME=docker --name <container-name> <image-name>:<tag>


## remove countainer : 
docker rm <container-name>
docker rm <container-id>


### platform cloud saya
- Google Cloud Platform adalah salah satu penyedia layanan komputasi awan (cloud computing) terkemuka yang disediakan oleh perusahaan teknologi besar, Google.  
- GCP menyediakan berbagai layanan komputasi awan, termasuk komputasi, penyimpanan data, analisis data, kecerdasan buatan (AI), dan banyak layanan lainnya yang memungkinkan perusahaan dan pengembang untuk membangun, mengelola, dan menjalankan aplikasi mereka di lingkungan cloud Google.

- Google Compute Engine: pengguna untuk menyewa mesin virtual (VM) untuk menjalankan aplikasi dan beban kerja mereka di lingkungan cloud.

## Google Cloud SQL adalah layanan manajemen basis data yang disediakan oleh Google Cloud Platform (GCP). Layanan ini dirancang untuk memudahkan pengguna untuk membuat, mengelola, dan mengelola basis data relasional di lingkungan cloud Google. 

- MySQL: MySQL adalah sistem manajemen basis data relasional open-source yang populer dan digunakan secara luas.
    ` Otomatisasi   : Google Cloud SQL menyediakan otomatisasi untuk tugas-tugas seperti penyediaan, perawatan, dan penjadwalan cadangan data.

    ` Keamanan  : Layanan ini mengintegrasikan lapisan keamanan yang kuat untuk melindungi data basis data.

    ` Kinerja   : dapat mengukur dan mengoptimalkan kinerja basis data menggunakan alat dan fitur yang disediakan.

    ` Ketersediaan  : Google Cloud SQL menawarkan opsi tingkat ketersediaan yang berbeda untuk menjaga basis data agar berjalan dengan lancar.

    ` Penyimpanan   : dapat dengan mudah mengatur kapasitas penyimpanan untuk basis data sesuai kebutuhan.

    ` Skalabilitas  : Layanan ini memungkinkan untuk dengan cepat menambahkan kapasitas dan mengubah ukuran basis data ketika diperlukan.