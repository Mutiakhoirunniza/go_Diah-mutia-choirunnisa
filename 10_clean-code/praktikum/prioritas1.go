package main

type user struct {
	// dalam bahasa go penamaan type data diawali dengan huruf besar agar di akses secara publik. " ID, Username, Password"
	id int
	// type data yang tidak sesuai,seharusnya type data menggunakan type data string, karena tidak semua username dan password menggunakan angka
	username int
	password int
}

type userservice struct {
	// Penamaan t sangat membingungkan
	t []user
}

// seharusnya penerima " u userservice " menggunakan *pointer agar lebih efesien dan tidak memakan banyak memori,karena sering di panggil  ( u *userservice )
func (u userservice) getallusers() []user {

	return u.t // Penamaan Variabel u.t sangat membingungkan

}

// seharusnya penerima " u userservice " menggunakan *pointer agar lebih efesien dan tidak memakan banyak memori,karena sering di panggil  ( u *userservice )
func (u userservice) getuserbyid(id int) user {

	for _, r := range u.t { // Penamaan Variabel u.t sangat membingungkan

		// penamaan variabel r sangat membingungkan
		if id == r.id {

			return r

		}

	}

	return user{}

}

// dan tidak tahu apa yang ingin di cetak atau di tampilkan
