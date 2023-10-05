# Docker 
Docker adalah platform perangkat lunak yang digunakan untuk mengembangkan, menguji, dan menjalankan aplikasi dalam lingkungan yang diisolasi yang disebut "container." 

Konsep dasar di balik Docker adalah kontainerisasi, yang mengizinkan mengemas aplikasi beserta semua dependensinya ke dalam kontainer yang independen, yang dapat dijalankan di berbagai lingkungan seperti mesin pengembangan, server produksi, atau bahkan di lingkungan cloud.

# konsep utama yang terkait dengan Docker   :

1. Kontainerisasi   : Docker memungkinkan Anda untuk mengemas aplikasi Anda beserta semua dependensinya ke dalam kontainer yang terisolasi. Ini memungkinkan aplikasi untuk berjalan dengan konsisten di berbagai lingkungan, menghindari konflik dan masalah yang sering terjadi dengan pemasangan manual.
2. Ringan   : Kontainer Docker jauh lebih ringan dibandingkan mesin virtual (VM). Mereka berbagi kernel sistem operasi tuan rumah, yang membuatnya hemat sumber daya dan memungkinkan Anda menjalankan banyak kontainer pada satu host.
3. Portabilitas : Docker membuat aplikasi dapat berjalan di mana saja yang mendukung Docker, baik itu di laptop pengembang, server lokal, atau di lingkungan produksi. Ini menghilangkan masalah yang sering muncul saat menggeser aplikasi antara lingkungan yang berbeda.
4. Skalabilitas : Anda dapat dengan mudah menambahkan atau mengurangi instance kontainer untuk menangani lonjakan lalu lintas atau beban kerja yang berubah. Docker juga mendukung orkestrasi kontainer, seperti Kubernetes, untuk mengelola kontainer dalam skala besar.
5. Mudah Digunakan  : Docker menyediakan perintah sederhana dan antarmuka pengguna yang mudah digunakan untuk mengelola kontainer. Ini membuatnya sangat berguna bagi pengembang, administrator sistem, dan tim operasi.
6. Ekosistem Besar  : Docker memiliki ekosistem yang luas dengan ribuan gambar kontainer yang tersedia di Docker Hub. Ini memungkinkan Anda untuk dengan mudah menemukan dan menggunakan aplikasi dan layanan yang telah di-dockerize.
7. Open Source  : Docker adalah perangkat lunak open source, yang berarti Anda dapat menggunakannya secara gratis dan mengkustomisasinya sesuai kebutuhan Anda.
8. Kemudahan Integrasi  : Docker dapat diintegrasikan dengan berbagai alat dan layanan seperti Jenkins, Git, dan banyak lagi, membuatnya sangat cocok untuk alur kerja pengembangan berkelanjutan (CI/CD).

# Infastructure and Deployment 
1. Docker 
2. system and software Deployment 
3. Cl/ CD 
4. kubernetes 

# countainer vs virtual Machines 
1. Container :
    > Abstraksi pada lapisan aplikasi
    > Kontainer memakan lebih sedikit ruang dibandingkan VM
    > Menangani lebih banyak aplikasi dan memerlukan lebih sedikit VM dan Sistem Operasi
2. Virtual Machines :
    > Abstraksi perangkat keras fisik
    > Setiap VM menyertakan salinan lengkap Sistem operasi
    > Juga lambat untuk boot

docker infrastructure 
1. Client  = > docker build 
             > docker pull
             > docker run 

2. Docker Host = > docker daemon > image > countainer 
3. registry  

# Syntax Docker
FROM -> Mendapatkan gambar dari registri buruh pelabuhan
RUN -> Jalankan perintah bash saat membuat Kontainer
ENV -> Tetapkan variabel di dalam Kontainer
ADD -> Salin file dengan beberapa proses lainnya
COPY -> Salin file
WORKDIR -> Atur direktori file kerja
ENTRYPOINT -> Jalankan perintah setelah selesai membangun Kontainer
CMD -> Jalankan perintah tetapi dapat ditimpa


# pembuatan container 
{   FROM golang:1.21.0-alpine
    WORKDIR /app
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download
    COPY . .
    RUN go build -o main.app .
    EXPOSE 8000
    CMD ["/app/main.app"] } 


# system  & Software Deployment 
 Deployment  adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/ produk yang telah di kerjakan oleh para pengembang seringkali untuk mengubah dari status Sementara ke permanent.

 # strategi Deployment  
1. big bang Deployment strategi atau sering disebut replace Deployment strategi 
2. rollout Deployment strategi 
3. blue/Green Deployment strategi 
4. canary Deployment strategi   


# Docker Volume
Docker Volume bisa dianggap sebagai storage atau tempat penyimpanan data di container. Tentunya saat kita membuat container kita tidak ingin ketika container kita mati data yang ada pada container ikut terhapus juga. Untuk itu kita dapat memanfaatkan Volume pada docker. 

#  Basic Command (docker  volume) 
 > create volume 
    $ docker create volume <name_volume> 
 > Remove Volume
    $ docker volume rm <name_volume> 
 > Add container to volume 
    $ docker container create --name <name_container> \
    --mount "type=volume, source=<name_volume>, dst=<folder_destionation>"

 # Basic Command (Create Network) 

>> Create Network
> list network
$ docker network ls
# create network
$ docker network create <name_network>
# delete network
$ docker network rm <name_network>


>> Regist container to network
> regist container to network
$ docker container create --name <container_name> \
--network <name_network>
# delete container from network
$ docker network disconect <name_network> <container_name>




 