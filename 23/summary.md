### Pengertian Continuous
Continuous Integration (CI), Continuous Delivery (CD), dan Continuous Deployment (CD) adalah praktik-praktik yang bertujuan untuk mempercepat pengembangan dan pengiriman perangkat lunak, meningkatkan kualitas, dan mengurangi risiko dalam siklus pengembangan perangkat lunak. 

## Pengertian Continuous Integration 
Continuous Integration adalah praktik pengembangan perangkat lunak di mana anggota tim pengembang secara teratur menggabungkan kode mereka ke dalam repositori bersama (biasanya Git) setidaknya sekali sehari atau bahkan lebih sering. 

# Tujuan CI
Tujuannya adalah untuk secara otomatis memeriksa apakah perubahan kode yang baru telah memecahkan atau merusak fungsionalitas yang ada. Proses CI mencakup pengujian otomatis dan linting kode untuk memastikan bahwa kode yang baru tidak mengakibatkan kerusakan pada aplikasi yang ada.

## Pengertian Continuous Delivery
Continuous Delivery adalah kelanjutan dari CI, di mana perubahan kode yang telah melewati proses CI dipersiapkan untuk rilis dan dapat diunggah ke produksi kapan saja.

## Pengertian Continuous Deployment
Continuous Deployment adalah tingkat yang lebih tinggi dari CD di mana setiap perubahan kode yang lulus pengujian otomatis segera diterapkan ke lingkungan produksi tanpa campur tangan manusia.

## Pengertian Continuous Monitoring: 
Continuous Monitoring Adalah Lingkungan produksi terus dimonitor untuk kinerja, keamanan, dan keandalan. 
Feedback Loop yaitu, Jika terjadi masalah, informasi kembali ke tim pengembangan untuk perbaikan lebih lanjut.

## Pengertian Continuous Feedback and Improvement
1. Feedback Gathering adalah Pemangku kepentingan dan pengguna memberikan umpan balik yang digunakan untuk perbaikan berkelanjutan. 

2. Back to Development yaitu, Berdasarkan umpan balik, tim pengembangan membuat perubahan dan memulai siklus CI/CD lagi untuk memperbarui perangkat lunak. 

## Siklus Continous Integration
    > Check-in Change           : Tim pengembang mengirimkan perubahan kode ke repositori bersama.
    > Fetch Change              : Sistem CI mengambil perubahan terbaru dari repositori.
    > Build                     : Kode yang diubah dikompilasi menjadi aplikasi yang dapat dijalankan.
    > Test                      : Aplikasi yang dikompilasi diuji dengan berbagai jenis pengujian, termasuk     pengujian unit dan pengujian fungsional.
    > Fail or Success           : Hasil pengujian ditentukan, dan apakah perubahan kode berhasil atau gagal     diidentifikasi.
    > Notify Success or Failure : Tim atau pemangku kepentingan diberi tahu tentang hasil pengujian, apakah     perubahan berhasil atau gagal.

## Siklus Continuous Deployment
    > Unit Test                     : Pengujian unit untuk memeriksa bagian-bagian kecil kode. (Auto)
    > Platform Test                 : Pengujian aplikasi pada platform yang sesuai. (Auto)
    > Deliver to Staging            : Pengiriman ke lingkungan staging untuk pengujian lebih lanjut. (Auto)
    > Application Acceptance Test   : Pengujian oleh pengguna atau pemangku kepentingan. (Auto)
    > Deploy to Production          : Implementasi otomatis ke produksi. (Auto)
    > Post Deploy Tests             : Pengujian pasca-implementasi. (Auto)

## Siklus Continuous Delivery
    > Unit Test                     : Pengujian unit untuk memeriksa bagian-bagian kecil kode. (Auto)
    > Platform Test                 : Pengujian aplikasi pada platform yang sesuai. (Auto)
    > Deliver to Staging            : Pengiriman ke lingkungan staging untuk pengujian lebih lanjut. (Auto)
    > Application Acceptance Test   : Pengujian oleh pengguna atau pemangku kepentingan. (Auto)
    > Deploy to Production          : Implementasi ke produksi. (Manual)
    > Post Deploy Tests             : Pengujian pasca-implementasi. (Auto)


### Tools CI/CD
1. Code & commit
    Git
    IBM Relational
    Github
    Stash
    Visual Studio
2. Build & Config
    Maven
    Docker
    Gradle
    CHEF
3. Scan & Test
    Gerrit
    Sonar
    Soasta
    Fitnesse
    JUnit
4. Release
    uDeploy
    Serena
    Release
5. Deploy
    Docker
    VMware
    Google App Engine
    Microsoft .net
    Openstack
    Kubernetes

    
## Kesimpulan

    > CI/CD adalah metodologi pengembangan perangkat lunak yang memfokuskan pada otomatisasi, pengujian, dan    pengiriman perubahan perangkat lunak secara terus-menerus. Dalam CI, perubahan kode secara rutin    diintegrasikan ke dalam repositori bersama, diikuti oleh proses otomatisasi yang mencakup kompilasi,    pengujian, dan pelaporan. CD memperluas konsep ini dengan memungkinkan pengiriman perangkat lunak ke    lingkungan produksi secara otomatis setelah melalui tahap pengujian, meminimalkan risiko kesalahan manusia.

    > Tujuan utama dari CI, CD, dan CD adalah meningkatkan efisiensi, kualitas, dan kecepatan pengembangan  perangkat lunak, serta mengurangi risiko dalam siklus pengembangan. Tingkat yang diterapkan dalam sebuah    proyek dapat bervariasi tergantung pada kebutuhan dan risiko yang terkait.

    > Manfaat utama dari CI/CD adalah peningkatan kecepatan pengembangan, peningkatan kualitas perangkat lunak,  dan kemampuan untuk merespons perubahan pasar lebih cepat. Dengan memungkinkan pengiriman perangkat lunak  yang lebih sering dan stabil, CI/CD membantu organisasi untuk menjadi lebih responsif, inovatif, dan efisien dalam pengembangan dan pengiriman perangkat lunak mereka.

### Kubernetes
Kubernetes adalah platform sumber terbuka (open-source) untuk mengelola aplikasi kontainer. Ini adalah salah satu proyek yang paling populer dalam lingkup komputasi kontainer dan orkestrasi. Kubernetes, sering disebut K8s (karena ada delapan huruf "ubernete" di tengah kata)

### Hal Yang Tidak Dapat Dilakukan Kubernetes
    > Tidak Membatasi Jenis Aplikasi yang Didukung
    > Tidak Mengimplementasikan Kode Sumber dan Tidak Membangun Aplikasi
    > Tidak Menyediakan Layanan Tingkat Aplikasi, Seperti Middleware
    > Tidak Mengatur Solusi Logging, Monitoring, atau Alerting

## Kubernetes Workloads
    Namespace       : Cara untuk mengorganisasi dan mengisolasi objek-objek Kubernetes dalam cluster.
    Pod             : Unit terkecil dalam Kubernetes yang berisi satu atau beberapa kontainer.
    Labels          : Metadata yang digunakan untuk mengkategorikan dan mengidentifikasi objek Kubernetes.
    Deployment      : Deployment digunakan untuk mengelola aplikasi dengan replika pod, mendukung rolling   updates, dan penskalaan otomatis.
    Rolling Update  : Metode pembaruan aplikasi tanpa waktu henti.
    Secret          : Objek yang digunakan untuk menyimpan data rahasia seperti kata sandi.
    Service         : Abstraksi lapisan jaringan untuk menghubungkan pod dalam cluster.
    Ingress         : Kontroler yang mengatur lalu lintas eksternal ke layanan dalam cluster.

## Domain Name System
    Domain Name System (DNS) adalah sistem yang digunakan untuk menerjemahkan alamat IP numerik ke nama domain  yang lebih mudah diingat.


