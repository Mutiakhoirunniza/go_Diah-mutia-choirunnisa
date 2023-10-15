### Penggunaan AI di Golang

## Beberapa alasan memilih go adalah :
    Dikembangkan oleh Google untuk pengembangan aplikasi server-side dan berbasis cloud.
    Go adalah compiled language sehingga lebih cepat dibanding interpreted language (contoh : Python).
    Mendukung concurrency melalui fitur goroutines sehingga dapat menjalankan banyak task secara simultan tanpa membenani memori.
    Memiliki standard libraries yang sangat kaya, didukung oleh komunitas yang kuat.
    Termasuk libraries untuk mendukung pengembangan AI.

###  Library Go untuk machine learning

## GoLearn
.   Bersifat open source.
.    Batteries included machine learning library. 
.    Simplicity & customizability.
contoh penerapan [link](https://github.com/sjwhitworth/golearn#getting-started)

## Gorgonia
.    Memudahkan menulis dan mengevaluasi persamaan matematika yang melibatkan array multidimensi.
.    Low-level library, seperti Theano, namun lebih tinggi dibanding Tensorflow
contoh penerapanny : [link](https://github.com/gorgonia/gorgonia#usage)

## Gonum
    Sebuah set packages yang didesain untuk menulis numerical & algoritma sains menjadi produktif, memiliki performa tinggi, dan scalable.
    Mirip seperti numpy dan spicy, library yang dibuat menggunakan Python
Referensi : [link](https://www.gonum.org/post/intro_to_gonum/)

## Gomid 
contoh penggunaanya : [Link] (https://github.com/surenderthakran/gomind#usage)

### Definisi 
## A.I as a services (AIaas)
    Sebuah tools AI yang dapat segera digunakan (pre-built or off-the-shelf AI tools).
    Berguna bagi bisnis yang ingin menerapkan AI tanpa merekrut ahlinya dan tanpa mengeluarkan biaya yang relatif banyak.
## Beberapa perusahaan yang menyediakan AIaas, contohnya :
1. Amazon Web Services (AWS)
2. Microsft Azure
3. Google Cloud Platform (GCP)
4. IBM
5. OpenAI
## Tipe-tipe AIaas

# Bots/Chatbots
    Bots/Chatbots create AI-Based customer services
    .    Amazon lex (AWS), Azure Bot Service (Microsft Azure), DialogFlow(GCP) 

# APIs
    APIs mengintegrasikan AI milik vendor dengan aplikasi sendiri 
    .    Amazon Rekognition, Amazon Comprehend
    .    Azure Cognitive Service, Azure Speech Services
    .    Google Cloud Vision, Cloud Natural language 

# Machine learning
    Machine learning build dan deploy  ML model easily 
    1.  Amazon SageMaker
    2.  Azure Machine learning
    3. Google Cloud AI

### Keuntungan menggunakan AI as a services adalah sebagai berikut :
    1. Pengurangan Cost
    2. Low-mode
    3. Proses deployment cepat
    4. Flexibility
    5. Usability
    6. Scalability
    7. Customization
## kelamahan menggunakan AI as a services sebagai berikut :
    1. Cost yang berlebih dalam jangka panjang
    2. Privasi dapat
    3. Keamanan
    4. Transparasi
    5. Vendor lock-in

### OpenAI API
    Target :
        1. Mendesain prompt yang tepat (prompt engineering)
        2. Membuat aplikasi berbasis AI menggunakan API OpenAI
[Dokumentasi OpenAI API](https://platform.openai.com/docs/introduction)
[OpenAI API Reference](https://platform.openai.com/docs/api-reference/introduction)
Opsional : [install library Go OpenAI](https://github.com/sashabaranov/go-openai)

### OpenAI API Pricing
    1. Start for free , mendapatkan free credit senilai $5 yang dapat digunakan selama 3 bulan pertama semenjak register.
    2. Pay as you go, setiap model memiliki tarif yang berbeda berbeda
## prompt engineering 
    PRINSIP     : 
        1. Gunakan  Prompt yang jelas dan spesifik
            Gunakan delimiters Contoh : ''', <>, <tag> </tag> untuk menandakan bagian mana  yang ingin menjadi input.
            Tuliskan Struktur output yang diinginkan.
            minta model  untuk memeriksa apakah input sudah sesuai. 
            few-Short prompoting.
        2. berikan waktu berfikir untuk menghindari solusi yang salah 
            menulis langkah- langkah yang dibutuhkan untuk menyelesaikan tugas dan output yang diinginkan. 
            menginstrusikan model untuk menuliskan solusinya sendiri sebelum masuk ke simpulan.


### Terimakasih 
telah selesai sudah Lms one alta 100% mastering Golang : [Link](https://one.alterra.academy/profile/)
notion-notion [link](https://cobalt-bike-c9e.notion.site/Reguler-7f796c6eb585414186cd01ff308e099a)

